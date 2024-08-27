package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
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
		salary, _ := strconv.ParseFloat(parsed[2], 64)

		fmt.Printf("%-10s %-10s $%-10s\n", parsed[0], parsed[1], humanize.Commaf(salary))
	}
}
