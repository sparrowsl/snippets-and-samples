package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a noun: ")
	scanner.Scan()
	noun := scanner.Text()

	fmt.Print("Enter a verb: ")
	scanner.Scan()
	verb := scanner.Text()

	fmt.Print("Enter an adjective: ")
	scanner.Scan()
	adjective := scanner.Text()

	fmt.Print("Enter an adverb: ")
	scanner.Scan()
	adverb := scanner.Text()

	result := createMadLibs(noun, verb, adjective, adverb)
	fmt.Println(result)
}

func createMadLibs(noun, verb, adjective, adverb string) string {
	return fmt.Sprintf("Do you %v your %v %v %v? That's hilarious", verb, adjective, noun, adverb)
}
