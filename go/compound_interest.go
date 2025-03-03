package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the principal amount? ")
	scanner.Scan()
	principal, _ := strconv.Atoi(scanner.Text())

	fmt.Print("What is the rate? ")
	scanner.Scan()
	rate, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("What is the number of years? ")
	scanner.Scan()
	years, _ := strconv.Atoi(scanner.Text())

	fmt.Print("What is the number of times the interest is compounded per year? ")
	scanner.Scan()
	compoundYears, _ := strconv.Atoi(scanner.Text())

	investment := float64(principal) * ((1 + ((rate / float64(100)) / float64(compoundYears))) * float64(compoundYears*years))

	fmt.Printf("$%v invested at %.2f for %v years\ncompounded %v times per year is $%.2f.\n", principal, rate, years, compoundYears, investment)

}
