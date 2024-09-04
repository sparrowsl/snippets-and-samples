package main

import (
	"fmt"
	"os"
	"strings"
)

const inputFile = "./macbeth.txt"

func main() {
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	store := make(map[string]int)
	words := strings.Fields(string(file))

	for _, word := range words {
		trimmed := strings.Trim(word, ",;:_[]?!@\"'.—")
		store[trimmed]++
	}

	// fmt.Println(store)
	displayGraph(store)
}

func displayGraph(data map[string]int) {
	for key, value := range data {
		fmt.Printf("%s: %s\n", key, strings.Repeat("*", value))
	}
}
