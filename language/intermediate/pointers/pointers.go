package pointers

import "fmt"

func Init() {
	// * = pointed value of
	// & = address of

	var x int = 50
	var addressOfX = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", addressOfX)
	fmt.Println("Address of x is pointing to:", *addressOfX)

	// will output x
	// pointed value of address of x = value of x
	fmt.Println("Pointed value of address of x is:", *(&x))

	// Let's mutate values across pointers
	var y int = 67
	var p = &y
	*p = 70
	fmt.Println("After changing, y is:", y)
	
	var z = new("string") // <- returns pointer
	fmt.Println(z)

	// TODO: add applicable uses of pointers in Go
}
