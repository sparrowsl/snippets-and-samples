package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("How many people? ")
	scanner.Scan()
	people, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Please enter a valid number!!")
		return
	}

	fmt.Print("How many pizzas do you have? ")
	scanner.Scan()
	pizzas, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Please enter a valid number!!")
		return
	}

	leftovers := people % pizzas
	fmt.Printf("%v people with %v pizzas\nEach person gets %v pieces of pizza.\n", people, pizzas, (people/pizzas)/2)
	fmt.Printf("There are %v leftovers pieces.\n", leftovers)

}
