package db

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

type User struct {
	Username string
	Password string
}

type Contact struct {
	Name  string
	Phone string
	Email string
	Desc  string
	File  string
}

// Initiate db connection and load .csv file.
// Returns db context for queries in API
func Main() *sql.DB {
	curdir, err := os.Getwd()

	fmt.Println("current directory: ", curdir)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("DBPORT"))
	user := os.Getenv("DBUSER")
	pw := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	// Initiate server connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pw, dbname)
	conn, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Read and process entries from the exploited_vulnerabilities .csv file
	users := readCsv("users.csv")
	processUsers(users, conn)
	contacts := readCsv("contacts.csv")
	processContacts(contacts, conn)

	return conn
}

// Reads in from .csv file and returns list of formatted entries
func readCsv(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(file, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	entries, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(file, err)
	}
	return entries
}

func processUsers(ent [][]string, dat *sql.DB) {
	len := len(ent)
	fmt.Printf(strconv.Itoa(len))
	//Drop the table so we can then update it
	dat.Exec("DROP TABLE users;")
	//Create the table
	dat.Exec("CREATE TABLE users (Username VARCHAR(1024) UNIQUE, Password VARCHAR(1024));")
	totalRows := 0
	for i := 1; i < len; i++ {
		result, err := dat.Exec("INSERT INTO users (Username, Password) VALUES ($1,$2);", ent[i][0], ent[i][1])
		if err != nil {
			log.Fatalf("row not affected: %v", err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatalf("could not get affected rows: %v", err)
		}
		totalRows += int(rowsAffected)
	}
	fmt.Println("Total rows inserted for user table: ", totalRows)
}

func processContacts(ent [][]string, dat *sql.DB) {
	len := len(ent)
	fmt.Printf(strconv.Itoa(len))
	//Drop the table so we can then update it
	dat.Exec("DROP TABLE contacts;")
	//Create the table
	dat.Exec("CREATE TABLE contacts (Name VARCHAR(1024), Email VARCHAR(1024), Description VARCHAR(1024)) ;")
	totalRows := 0
	for i := 1; i < len; i++ {
		result, err := dat.Exec("INSERT INTO contacts (Name, Email, Description) VALUES ($1,$2,$3);", ent[i][0], ent[i][1], ent[i][2])
		if err != nil {
			log.Fatalf("row not affected: %v on i:%d", err, i)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatalf("could not get affected rows: %v", err)
		}
		totalRows += int(rowsAffected)
	}
	fmt.Println("Total rows inserted for contacts table: ", totalRows)
}
