package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is your current age? ")
	scanner.Scan()
	age, _ := strconv.Atoi(scanner.Text())

	fmt.Print("At what age would you like to retire? ")
	scanner.Scan()
	retireAge, _ := strconv.Atoi(scanner.Text())

	// quit program if user is already retire
	if age >= retireAge {
		fmt.Println("You are already retired!!")
		return
	}

	currentYear, _ := strconv.Atoi(time.Now().Format("2006"))

	fmt.Printf("You have %v years left until you can retire.\n", retireAge-age)
	fmt.Printf("It's %v, so you can retire in %v.\n", currentYear, currentYear+(retireAge-age))
}
