package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	type Employee struct {
		FirstName string
		LastName  string
		Position  string
		Date      string
	}
	employees := []Employee{}

	db, err := sql.Open("sqlite", "./records.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a search string: ")
	scanner.Scan()

	search := scanner.Text()

	rows, err := db.Query("SELECT * FROM employees WHERE first_name LIKE ? OR last_name LIKE ?",
		fmt.Sprintf("%%%s%%", search),
		fmt.Sprintf("%%%s%%", search),
	)
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
