package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(("What is your name? "))
	scanner.Scan()
	sayHello(scanner.Text())
}

func sayHello(name string) {
	fmt.Printf("Hello, %v, nice to meet you!\n", name)
}
