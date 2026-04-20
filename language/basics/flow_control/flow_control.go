package flowcontrol

import "fmt"

func Init() {
	// basic if else-if flow
	var x = 5
	if x > 10 {
		fmt.Println("x is greater than 10.")
	} else if x < 10 {
		fmt.Println("x is less than 10")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	}

	// shorter if flow: injects statement value into condition
	if y := 5; y < 10 {
		fmt.Println("y is less than 10.")
	}

	/* NOTE: flow statement variable is not scoped to out of
	   condition scope */
	// fmt.Println("Does y still exist?:", y) // <- compile error

	// simple for loop
	for i := 0; i < 4; i++ {
		fmt.Printf("Printing %d...\n", i)
	}
}
