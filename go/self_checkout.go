package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	items := make([]map[string]float64, 3)
	var price float64
	var quantity int

	for i := 0; i < 3; i++ {
		fmt.Printf("Enter the price of item %v: ", i+1)
		scanner.Scan()
		price, _ = strconv.ParseFloat(scanner.Text(), 64)

		fmt.Printf("Enter the quantity of item %v: ", i+1)
		scanner.Scan()
		quantity, _ = strconv.Atoi(scanner.Text())

		items[i] = map[string]float64{"price": price, "quantity": float64(quantity)}
	}

	var subtotal float64
	taxRate := 5.5

	for _, v := range items {
		subtotal += v["price"] * v["quantity"]
	}

	tax := (taxRate / 100) * subtotal
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Printf("Tax: $%.2f\n", tax)
	fmt.Printf("Total: $%.2f\n", subtotal+tax)
}
