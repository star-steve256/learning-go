package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	fmt.Println("===== Number guessing game =====")
	var target = rand.Intn(99) + 1
	var numOfAttempts = 0

	for {
		var userInput int
		fmt.Printf("#%d Pick a number between 1-99: ", numOfAttempts + 1)
		_, err := fmt.Scan(&userInput)
		if err != nil {
			log.Fatal("ERROR at scanning input:", err)
		}

		if userInput < 1 || userInput > 99 {
			fmt.Println("Please choose number between 1-99.")
			fmt.Println("================================")
			continue
		}

		numOfAttempts++
		if userInput == target{
			break
		} else if userInput > target {
			fmt.Printf("Wrong guess: %d > target\n", userInput)
		} else if userInput < target {
			fmt.Printf("Wrong guess: %d < target\n", userInput)
		}

		fmt.Println("================================")
	}

	fmt.Println("\n🚀 Guess CORRECT: Answer is", target)
	fmt.Println("🔍 Number of attempts:", numOfAttempts)
}
