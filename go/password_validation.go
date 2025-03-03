package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the password? ")
	scanner.Scan()
	password := scanner.Text()

	if password == "abc$123" {
		fmt.Println("Welcome!")
	} else {
		fmt.Println("I don't know you.")
	}
}
