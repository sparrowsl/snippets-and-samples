package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	byt := flag.Bool("b", false, "Count bytes")

	flag.Parse() // parse all flags provided by user

	fmt.Println(count(os.Stdin, *lines, *byt))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

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
