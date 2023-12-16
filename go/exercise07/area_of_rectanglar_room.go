package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the length of room in feet? ")
	scanner.Scan()
	length, _ := strconv.Atoi(scanner.Text())

	fmt.Print("What is the width of room in feet? ")
	scanner.Scan()
	width, _ := strconv.Atoi(scanner.Text())

	squareMeters := (float64(length) * float64(width) * 0.09290304)
	fmt.Printf("You entered dimensions of %v feet by %v feet.\n", length, width)
	fmt.Printf("The area is\n%v square feet\n%v square meters\n", length*width, squareMeters)
}
