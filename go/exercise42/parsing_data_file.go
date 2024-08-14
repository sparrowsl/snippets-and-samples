package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	trimmed := strings.TrimSpace(string(file))

	fmt.Printf("%-10s %-10s %-10s\n", "Last", "First", "Salary")
	fmt.Println(strings.Repeat("-", 30))

	for _, line := range strings.Split(trimmed, "\n") {
		parsed := strings.Split(line, ",")
		fmt.Printf("%-10s %-10s %-10s\n", parsed[0], parsed[1], parsed[2])

	}
}
