package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	byt := flag.Bool("b", false, "Count bytes")
	filename := flag.String("f", "", "Count bytes | lines | words in the file")

	flag.Parse() // parse all flags provided by user

	fmt.Println(count(os.Stdin, *lines, *byt, *filename))
}

func count(r io.Reader, countLines bool, countBytes bool, file string) int {
	var scanner *bufio.Scanner

	if file != "" {
		validFile, err := os.ReadFile(file)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		scanner = bufio.NewScanner(strings.NewReader(string(validFile)))
	} else {
		scanner = bufio.NewScanner(r)
	}

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}
