package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// John Smith
// Jackie Jackson
// Chris Jones
// Amanda Cullen
// Jeremy Goodwin

var filename = "./employee_list.txt"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	employees := strings.Split(strings.TrimSpace(string(file)), "\n")
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

	// Write to file
	err = os.WriteFile(filename, []byte(strings.Join(employees, "\n")), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func displayEmployees(names []string) {
	fmt.Printf("There are %d employees:\n", len(names))

	for _, name := range names {
		fmt.Printf("  - %s\n", name)
	}
}
