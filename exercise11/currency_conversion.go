package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("How many euros are you exchanging? ")
	scanner.Scan()
	euros, _ := strconv.Atoi(scanner.Text())

	fmt.Printf("What is the exchange rate? ")
	scanner.Scan()
	exchangeRate, _ := strconv.ParseFloat(scanner.Text(), 64)

	toDollars := (float64(euros) * exchangeRate) / exchangeRate

	fmt.Printf("%v euros at an exchange rate of %v is\n%.2f U.S. dollars.\n", euros, exchangeRate, toDollars)
}
