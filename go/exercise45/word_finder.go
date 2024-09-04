package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const inputFile = "./input.txt"

var outputFile = "./output.txt"

func main() {
	output := scanInput("Enter output file: ")
	if output != "" {
		outputFile = output
	}

	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	strings.Count("s string", "")
	replaced := strings.ReplaceAll(string(file), "utilize", "use")

	os.WriteFile(outputFile, []byte(replaced), os.ModePerm)
}

func scanInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)
	scanner.Scan()

	return strings.TrimSpace(scanner.Text())
}
