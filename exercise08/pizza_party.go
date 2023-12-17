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
	people, _ := strconv.Atoi(scanner.Text())

	fmt.Print("How many pizzas do you have? ")
	scanner.Scan()
	pizzas, _ := strconv.Atoi(scanner.Text())

	leftovers := people % pizzas
	fmt.Printf("%v people with %v pizzas\nEach person gets %v pieces of pizza.\n", people, pizzas, (people/pizzas)/2)
	fmt.Printf("There are %v leftovers pieces.\n", leftovers)

}
