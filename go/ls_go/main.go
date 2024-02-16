package main

import (
	"fmt"
	"os"
)

func main() {
	// Get current path
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	dirs, err := os.ReadDir(cwd)
	_ = err

	for _, file := range dirs {
		if file.IsDir() {
			fmt.Println(file)
		} else {
			fmt.Println(file)
		}
	}
}
