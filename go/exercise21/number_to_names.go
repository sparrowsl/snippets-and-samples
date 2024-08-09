package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// maps can't be const at global scope
var months = map[int64]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Please enter the number of the month: ")
	scanner.Scan()

	number, _ := strconv.Atoi(scanner.Text())
	value, ok := months[int64(number)]
	if ok {
		fmt.Printf("The name of the month is %s.\n", value)
	}

}
