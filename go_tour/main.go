// package - "main" on root directory
package main

import (
	"fmt"
	"go_tour/concurrency"
	"go_tour/functions"
	"math/rand"
)

// program entry point
func main() {
	initMain()
	concurrency.InitConcurrency()
	functions.InitFunctions()
}

func initMain() {
	fmt.Println("Hello World.")

	// generate random number
	randNum := rand.Intn(100) + 1
	fmt.Println("Random number generated:", randNum)

	fmt.Println()
}
