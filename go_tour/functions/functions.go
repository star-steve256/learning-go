package functions

import (
	"fmt"
	"math"
)

func InitFunctions() {
	fmt.Println("===== FUNCTIONS =====")
	// functions are first-class in Golang

	// simulates math composite functions
	// (x + 2)^2
	fmt.Println("(3 + 2)^2 =", gMorph(fMorph(3)))

	// functions can be passed as arguments
	// short password
	validatePasswordAndPass("hello", printPasswordSuccess)
	// long password
	validatePasswordAndPass("helloworld", printPasswordSuccess)

	fmt.Println("3 + 2 =", add(3, 2))

	fmt.Println()
}

func fMorph(x int) (res int) {
	return x + 2
}

func gMorph(x int) (res int) {
	res = int(math.Pow(float64(x) , 2))
	return
}

// middleware simulation
func validatePasswordAndPass(password string, next func(password string)) {
	if len(password) < 8 {
		fmt.Println("Password too short. Please try again.")
	} else {
		next(password)
	}
}

func printPasswordSuccess(password string) {
	fmt.Println("Password updated successfully.")
	fmt.Println("Password:", password)
}

// return without specifying return values
func add(x int, y int) (res int) {
	res = x + y
	return
}
