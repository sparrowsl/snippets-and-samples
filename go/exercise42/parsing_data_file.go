package main

import (
	"encoding/csv"
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

	reader := csv.NewReader(strings.NewReader(string(file)))
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%-10s %-10s %-10s\n", "Last", "First", "Salary")
	fmt.Println(strings.Repeat("-", 30))

	for _, row := range data {
		salary, _ := strconv.ParseFloat(row[2], 64)
		fmt.Printf("%-10s %-10s $%-10s\n", row[0], row[1], humanize.Commaf(salary))
	}
}
