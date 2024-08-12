package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.NewSource(time.Now().Unix())
	scanner := bufio.NewScanner(os.Stdin)

	players := []string{}

	for {
		fmt.Print("Enter a name: ")
		scanner.Scan()
		name := strings.TrimSpace(scanner.Text())

		if name == "" {
			break
		}

		players = append(players, name)
	}

	fmt.Printf("The winner is... %s\n", shuffle(players))
}

func shuffle(array []string) string {
	num := rand.Intn(len(array))

	return array[num]
}
