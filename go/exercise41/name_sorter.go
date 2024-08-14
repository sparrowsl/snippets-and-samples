package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const file = "names.txt"

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// trim the spaces at the start and end
	trimmed := strings.TrimSpace(string(data))

	// convert to string array
	names := strings.Split(string(trimmed), "\n")
	slices.Sort(names)

	fmt.Printf("Total of %d names\n", len(names))
	fmt.Println(strings.Repeat("-", 10))
	for _, name := range names {
		fmt.Println(name)
	}
}
