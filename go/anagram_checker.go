package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter two strings and I'll tell you if they are anagrams:")

	fmt.Print("Enter the first string: ")
	scanner.Scan()
	firstString := scanner.Text()

	fmt.Print("Enter the second string: ")
	scanner.Scan()
	secondString := scanner.Text()

	if isAnagram(firstString, secondString) {
		fmt.Printf("%s and %s are anagrams.\n", firstString, secondString)
	} else {
		fmt.Printf("%s and %s are not anagrams.\n", firstString, secondString)
	}
}

func isAnagram(string1 string, string2 string) bool {
	// check if the strings have same length
	if len(string1) != len(string2) {
		return false
	}

	for _, val := range string1 {
		if !strings.Contains(string2, string(val)) {
			return false
		}
	}

	return true
}
