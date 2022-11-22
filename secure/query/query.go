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
	Phone string
	Email string
	Desc  string
	File  string
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

func NewUser(dat *sql.DB, uname string, pass string) string {
	//SQL injection
	// ', ''); DROP TABLE users; --
	totalRows := 0
	// exec := "INSERT INTO users (Username, Password) VALUES ('" + uname + "', '" + pass + "');"
	// fmt.Println(exec)
	// result, err := dat.Exec(exec)

	result, err := dat.Exec("INSERT INTO users (Username, Password) VALUES ($1,$2);", uname, pass)
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

// // Performs a query on the database and returns a list of products.
// func QueryProducts(dat *sql.DB, vendor string) []string {
// 	rows, err := dat.Query("SELECT DISTINCT product FROM vulnerabilities WHERE vendorProject = '" + vendor + "' ORDER BY product ASC;")
// 	if err != nil {
// 		log.Fatalf("could not execute query: %v", err)
// 	}

// 	products := []string{}
// 	for rows.Next() {
// 		var product string

// 		if err := rows.Scan(&product); err != nil {
// 			log.Fatalf("could not scan row: %v", err)
// 		}
// 		products = append(products, product)
// 	}
// 	return products
// }

// // Performs a query on the database and returns a list of records given vendor, project
// func QueryResults(dat *sql.DB, vendor string, product string) []ProdW {
// 	rows, err := dat.Query("SELECT * FROM vulnerabilities WHERE vendorProject = '" + vendor + "' AND product = '" + product + "';")
// 	if err != nil {
// 		log.Fatalf("could not execute query: %v", err)
// 	}

// 	result := []ProdW{}
// 	for rows.Next() {
// 		res := ProdW{}

// 		if err := rows.Scan(&res.Cvekey, &res.Cveid, &res.VendorProject, &res.Product, &res.VulnerabilityName, &res.DateAdded, &res.ShortDescription, &res.RequiredAction, &res.DueDate, &res.Notes); err != nil {
// 			log.Fatalf("could not scan row: %v", err)
// 		}
// 		result = append(result, res)
// 	}

// 	return result
// }

// func QueryScores(dat *sql.DB, prodList []ProdW) ScoreResult {
// 	var cveList []string
// 	for _, product := range prodList {
// 		cveList = append(cveList, product.Cveid)
// 	}
// 	query := "SELECT v3score, v2score FROM base WHERE "

// 	for index, cve := range cveList {
// 		if index == 0 {
// 			queryStr := "\"cveid\" = '" + cve + "'"
// 			query = query + queryStr
// 		} else {
// 			queryStr := " OR \"cveid\" = '" + cve + "'"
// 			query = query + queryStr
// 		}
// 	}
// 	query = query + ";"

// 	//println(query)

// 	rows, err := dat.Query(query)
// 	if err != nil {
// 		log.Fatalf("could not execute query: %v", err)
// 	}

// 	// This row will return the v3score & the v2score
// 	result := ScoreResult{}
// 	V3result := Score{Critical: 0, High: 0, Medium: 0, Low: 0, NA: 0}
// 	V2result := Score{Critical: 0, High: 0, Medium: 0, Low: 0, NA: 0}

// 	for rows.Next() {
// 		res := ScoreEntry{}
// 		// scan here
// 		if err := rows.Scan(&res.V3score, &res.V2score); err != nil {
// 			log.Fatalf("could not scan row: %v", err)
// 		}

// 		// filter to add count to appropriate list
// 		if strings.Contains(res.V3score, "CRITICAL") {
// 			V3result.Critical += 1
// 		} else if strings.Contains(res.V3score, "HIGH") {
// 			V3result.High += 1
// 		} else if strings.Contains(res.V3score, "MEDIUM") {
// 			V3result.Medium += 1
// 		} else if strings.Contains(res.V3score, "LOW") {
// 			V3result.Low += 1
// 		} else {
// 			V3result.NA += 1
// 		}

// 		if strings.Contains(res.V2score, "CRITICAL") {
// 			V2result.Critical += 1
// 		} else if strings.Contains(res.V2score, "HIGH") {
// 			V2result.High += 1
// 		} else if strings.Contains(res.V2score, "MEDIUM") {
// 			V2result.Medium += 1
// 		} else if strings.Contains(res.V2score, "LOW") {
// 			V2result.Low += 1
// 		} else {
// 			V2result.NA += 1
// 		}
// 	}
// 	result.V3list = V3result
// 	result.V2list = V2result
// 	return result
// }

// func QueryWeaknesses(dat *sql.DB, prodList []ProdW) []Weakness {
// 	var cveList []string
// 	for _, product := range prodList {
// 		cveList = append(cveList, product.Cveid)
// 	}
// 	query := "SELECT source, cwename, COUNT(*) as count FROM cwe WHERE cveid IN ("

// 	for index, cve := range cveList {
// 		if index == 0 {
// 			query += ("'" + cve + "'")
// 		} else {
// 			query += (",'" + cve + "'")
// 		}
// 	}

// 	query = query + ") GROUP BY source, cwename;"

// 	//println(query)

// 	rows, err := dat.Query(query)
// 	if err != nil {
// 		log.Fatalf("could not execute query: %v", err)
// 	}

// 	result := []Weakness{}

// 	for rows.Next() {
// 		res := Weakness{}
// 		// scan here
// 		if err := rows.Scan(&res.Name, &res.CWEID, &res.Count); err != nil {
// 			log.Fatalf("could not scan row: %v", err)
// 		}
// 		result = append(result, res)
// 		//println("name: ", res.CWEID, "\tCount: ", res.Count, "\n")
// 	}
// 	return result
// }
