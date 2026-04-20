// every go program starts with a package
// no package(root dir) - "main"
package main

import (
	"fmt"
	"intermediate/pointers"
	"log"
)

// main() function - where program runs
func main() {
	fmt.Println("===== intermediate main.go =====")
	var numOfOptions = 1
	for {
		fmt.Println(`Choose function to run: 
1. runPointers()
0. Exit`)

		fmt.Print(">> ")
		// scans user input
		var inputOption int
		_,err := fmt.Scan(&inputOption)
		if err != nil {
			log.Fatal("ERROR at scanning input:", err)
		}

		// option validation and exit check
		if inputOption < 0 || inputOption > numOfOptions {
			fmt.Printf("Please choose option between 0-%d\n", numOfOptions)
			fmt.Println("=================================")
			continue
		}
		if inputOption == 0 {
			break
		}

		var runnableFunctions = []func(){
			runPointers,
		}

		// runs respective chosen function
		runnableFunctions[inputOption - 1]()

		// adds separators and new lines after each run
		fmt.Println("=================================")
	}

	fmt.Println("Program terminated.")
}

// option 1
func runPointers() {
	pointers.Init()
}
