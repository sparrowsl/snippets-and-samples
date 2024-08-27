package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

const file string = "./employees.db"

func main() {
	type Employee struct {
		FirstName string
		LastName  string
		Position  string
		Date      string
	}
	employees := []Employee{}

	db, err := sql.Open("sqlite", file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * from employees")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var emp Employee

		rows.Scan(&emp.FirstName, &emp.LastName, &emp.Position, &emp.Date)
		employees = append(employees, emp)
	}

	fmt.Printf("%-19s | %-17s | %s\n", "Name", "Position", "Separation Date")
	fmt.Println("-----------------------------------------------------")
	for _, emp := range employees {
		fmt.Printf("%-19s | %-17s | %s\n",
			emp.FirstName+" "+emp.LastName,
			emp.Position,
			emp.Date,
		)
	}
}
