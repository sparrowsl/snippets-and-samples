package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the length in feet? ")
	scanner.Scan()
	length, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter the width in feet? ")
	scanner.Scan()
	width, _ := strconv.Atoi(scanner.Text())

	fmt.Printf("You will need to purchase %v gallons of\npaint to cover %v square feet.\n", (length*width)/350, length*width)

}
