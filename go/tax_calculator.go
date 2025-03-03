package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const tax float64 = 5.5

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the order amount? ")
	scanner.Scan()
	amount := scanner.Text()

	fmt.Print("What is the state? ")
	scanner.Scan()
	state := scanner.Text()

	subtotal, _ := strconv.ParseFloat(amount, 64)

	// Check the amount and state condition
	if state == "WI" {
		fmt.Printf("The subtotal is $%.2f\n", subtotal)
		fmt.Printf("The tax is $%.2f\n", (subtotal*tax)/100)
		fmt.Printf("The total is $%.2f\n", subtotal+((subtotal*tax)/100))
		return
	}

	fmt.Printf("The total is $%.2f\n", subtotal)
}
