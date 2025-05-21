package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	var high, low, num_trials int = 1, 100, 3
// 	groundTruth := generateGroundTruthNum(high, low)

// 	for trial := 0; trial < num_trials; trial++ {
// 		var guess int = guessNum(high, low)

// 		if guess > groundTruth {
// 			fmt.Println(guess, "is too high! Try again (", num_trials-(trial+1), "trials left )!")
// 			fmt.Println()
// 		} else if guess < groundTruth {
// 			fmt.Println(guess, "is too low! Try again (", num_trials-(trial+1), "trials left )!")
// 			fmt.Println()
// 		} else {
// 			fmt.Println(guess, "is correct! Good job 💡!")
// 			return
// 		}
// 	}

// 	fmt.Println("You have exhausted your trials. Come back later 😭!")
// }

func generateGroundTruthNum(low int, high int) int {
	var randNum int = rand.Intn(high) + 1

	if randNum < low {
		randNum = low
	}
	return randNum
}

func guessNum(low int, high int) int {
	var err error
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess a number between", low, "and", high, ":")

	var guess string

	guess, err = reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	guess = strings.Replace(guess, "\r\n", "", 1)

	var guessInt int
	guessInt, err = strconv.Atoi(guess)

	if err != nil {
		log.Fatal(err)
	}

	return int(guessInt)
}
