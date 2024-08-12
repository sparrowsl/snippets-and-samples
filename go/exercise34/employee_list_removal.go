package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var employees = []string{
	"John Smith",
	"Jackie Jackson",
	"Chris Jones",
	"Amanda Cullen",
	"Jeremy Goodwin",
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	displayEmployees(employees)

	fmt.Print("\nEnter an employee name to remove: ")
	scanner.Scan()
	name := scanner.Text()

	if !slices.Contains(employees, name) {
		fmt.Printf("the name %q does not exists\n", name)
		os.Exit(0)
	}

	employees = slices.DeleteFunc(employees, func(emp string) bool {
		return emp == name
	})

	displayEmployees(employees)
}

func displayEmployees(names []string) {
	fmt.Printf("There are %d employees:\n", len(names))

	for _, name := range names {
		fmt.Printf("  - %s\n", name)
	}
}
