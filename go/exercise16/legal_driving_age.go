package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is your age? ")
	scanner.Scan()
	age, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("enter a valid number")
		os.Exit(0)
	}

	if age < 16 {
		print("You are not old enough to legally drive")
	} else {
		print("You are old enough to legally drive")
	}
}
