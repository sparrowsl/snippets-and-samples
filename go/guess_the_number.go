package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var difficulty = map[int]int{
	1: 10,
	2: 100,
	3: 1000,
}

func main() {
	rand.NewSource(time.Now().Unix())
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Let's play Guess the Number.")
	fmt.Print("Pick a difficulty level (1, 2, or 3): ")
	scanner.Scan()

	level, err := strconv.Atoi(scanner.Text())
	if err != nil || level > 3 || level < 1 {
		fmt.Printf("%d is not a valid difficulty level!\n", level)
		os.Exit(0)
	}

	secret := rand.Intn(difficulty[level])
	guessCount := 0

	for {
		fmt.Print("I have my number. What is your guess? ")
		scanner.Scan()
		guess, err := strconv.Atoi(scanner.Text())

		guessCount += 1
		if err != nil {
			fmt.Println("Invalid guess!!")
			continue
		}

		if guess > secret {
			fmt.Println("Too high. Guess again.")
		} else if guess < secret {

			fmt.Println("Too low. Guess again.")
		} else {
			fmt.Printf("You got it in %d guesses.\n", guessCount)
			break
		}
	}
}
