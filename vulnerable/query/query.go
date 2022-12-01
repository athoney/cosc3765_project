package query

import (
	"database/sql"
	"fmt"
)

type User struct {
	Username string
	Password string
}

type Contact struct {
	Name  string
	Email string
	Desc  string
}

// Performs a query on the database and returns a list of products.
func QueryUsers(dat *sql.DB) []User {
	rows, err := dat.Query("SELECT * FROM users;")
	if err != nil {
		fmt.Printf("could not execute query: %v", err)
		return []User{}
	}

	result := []User{}
	for rows.Next() {
		res := User{}

		if err := rows.Scan(&res.Username, &res.Password); err != nil {
			fmt.Printf("could not scan row: %v", err)
		}
		result = append(result, res)
	}

	return result
}

// Performs a query on the database and returns a list of requests.
func QueryRequests(dat *sql.DB) []Contact {
	rows, err := dat.Query("SELECT * FROM contacts;")
	if err != nil {
		fmt.Printf("could not execute query: %v", err)
		return []Contact{}
	}

	result := []Contact{}
	for rows.Next() {
		res := Contact{}

		if err := rows.Scan(&res.Name, &res.Email, &res.Desc); err != nil {
			fmt.Printf("could not scan row: %v", err)
		}
		result = append(result, res)
	}

	return result
}

func NewUser(dat *sql.DB, uname string, pass string) string {
	//SQL injection
	// ', ''); DROP TABLE users; --
	totalRows := 0
	exec := "INSERT INTO users (Username, Password) VALUES ('" + uname + "', '" + pass + "');"
	fmt.Println(exec)
	result, err := dat.Exec(exec)

	// result, err := dat.Exec("INSERT INTO users (Username, Password) VALUES ($1,$2);", uname, pass)
	if err != nil {
		fmt.Printf("row not affected: %v", err)
		return "FAIL"
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("could not get affected rows: %v", err)
		return "FAIL"
	}
	totalRows += int(rowsAffected)

	fmt.Println("Total rows inserted for user table: ", totalRows)
	return "OK"
}

func NewRequest(dat *sql.DB, name string, email string, desc string) string {
	//SQL injection
	// ', '', ''); DROP TABLE users; --
	totalRows := 0
	exec := "INSERT INTO contacts (Name, Email, Description) VALUES('" + name + "', '" + email + "', '" + desc + "');"
	fmt.Println(exec)
	result, err := dat.Exec(exec)

	if err != nil {
		fmt.Printf("row not affected: %v", err)
		return "FAIL"
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("could not get affected rows: %v", err)
		return "FAIL"
	}
	totalRows += int(rowsAffected)

	fmt.Println("Total rows inserted for user table: ", totalRows)
	return "OK"
}
