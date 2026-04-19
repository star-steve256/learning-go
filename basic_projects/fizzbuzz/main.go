package main

import "fmt"

func main() {
	fmt.Println("===== Variable FizzBuzz =====")

	var fizzFactor, buzzFactor int
	var start, end int

	fmt.Printf("Enter the factor where Fizz will output: ")
	fmt.Scan(&fizzFactor)
	fmt.Printf("Enter the factor where Buzz will output: ")
	fmt.Scan(&buzzFactor)
	fmt.Printf("Enter starting point: ")
	fmt.Scan(&start)
	fmt.Printf("Enter ending point: ")
	fmt.Scan(&end)

	fmt.Println("====================")

	for i := start; i <= end; i++ {
		if i % fizzFactor == 0 && i % buzzFactor == 0 {
			fmt.Printf("#%d.\tFizzBuzz\n", i)
		} else if i % fizzFactor == 0 {
			fmt.Printf("#%d.\tFizz\n", i)
		} else if i % buzzFactor == 0 {
			fmt.Printf("#%d.\tBuzz\n", i)
		} else {
			fmt.Printf("#%d.\n", i)
		}
	}

	fmt.Println("====================")
}