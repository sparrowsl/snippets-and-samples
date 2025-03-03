package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the first number? ")
	scanner.Scan()
	firstNumber, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("What is the second number? ")
	scanner.Scan()
	secondNumber, _ := strconv.ParseFloat(scanner.Text(), 64)

	sum := add(firstNumber, secondNumber)
	difference := subtract(firstNumber, secondNumber)
	product := multiply(firstNumber, secondNumber)
	quotient := divide(firstNumber, secondNumber)

	fmt.Printf("%v + %v = %v\n", firstNumber, secondNumber, sum)
	fmt.Printf("%v - %v = %v\n", firstNumber, secondNumber, difference)
	fmt.Printf("%v * %v = %v\n", firstNumber, secondNumber, product)
	fmt.Printf("%v / %v = %v\n", firstNumber, secondNumber, quotient)
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
