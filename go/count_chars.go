package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the input string? ")
	scanner.Scan()
	value := scanner.Text()

	if len(strings.TrimSpace(value)) == 0 {
		fmt.Println("Please enter something.")
		return
	}

	total := countChars(value)
	fmt.Printf("%v has %v characters.\n", value, total)
}

func countChars(sentence string) int {
	return len(sentence)
}
