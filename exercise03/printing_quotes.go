package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What is the quote? ")
	scanner.Scan()
	quote := scanner.Text()

	fmt.Print("Who said it? ")
	scanner.Scan()
	author := scanner.Text()

	result := printQuote(author, quote)
	fmt.Println(result)
}

func printQuote(author, quote string) string {
	return fmt.Sprintf("%v says, \"%v\"", author, quote)
}
