package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the principal: ")
	scanner.Scan()
	principal, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter the rate of interest: ")
	scanner.Scan()
	rate, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Enter the years: ")
	scanner.Scan()
	years, _ := strconv.Atoi(scanner.Text())

	interest := (float64(principal) * (1 + float64(years)*(rate/100)))

	fmt.Printf("After %v years at %.1f%%, the investment will be worth $%v.\n", years, rate, interest)

}
