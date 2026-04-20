// every go program starts with a package
// no package(root dir) - "main"
package main

import (
	flowcontrol "basics/flow_control"
	"basics/functions"
	importexport "basics/import_export"
	"basics/variables"
	"fmt"
	"log"
)

// main() function - where program runs
func main() {
	fmt.Println("===== basics main.go =====")
	var numOfOptions = 4
	for {
		fmt.Println(`Choose function to run: 
1. runImportExport()
2. runFunctions()
3. runVariables()
4. runFlowControl()
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
			runImportExport,
			runFunctions,
			runVariables,
			runFlowControl,
		}

		// runs respective chosen function
		runnableFunctions[inputOption - 1]()

		// adds separators and new lines after each run
		fmt.Println("=================================")
	}

	fmt.Println("Program terminated.")
}

// option 1
func runImportExport() {
	fmt.Println("Current Exported version:", importexport.ExportedVersion)

	// fmt.Println("Local version:", importexport.localVersion)
	// ^ this will give compile time error

	// exported variable can be modified freely
	importexport.ExportedVersion = "NewE0.1.0"
	fmt.Println("After modifiying ExportedVersion in main.go ...")
	importexport.PrintVersions()

	// local version can still be changed by getters and setters
	fmt.Println("Current local version:", importexport.GetLocalVersion()) // <- usage of public getter
	importexport.SetLocalVersion("Newl0.1.0") // <- usage of public setter
	fmt.Println("After modifying localVersion in main.go ...")
	importexport.PrintVersions()
}

// option 2
func runFunctions() {
	fmt.Println("Usage of functions")
	fmt.Println("1 + 2 =", functions.Add(1, 2))
	fmt.Println("1 + 2 =", functions.FunctionOfAdding(1, 2)) // <- variable function
	fmt.Println("1 - 2 =", functions.Subtract(1, 2))
	fmt.Println("0.5 - 0.2 =", functions.SubtractFloat(0.5, 0.2))
	fmt.Println("==========================")
	fmt.Println("Simulation of y = (x + 2)²")
	y := functions.Add2(3, functions.Square)
	fmt.Println("x = 3; y =", y)
}

// option 3
func runVariables() {
	variables.MainRun()
}

// option 4
func runFlowControl() {
	flowcontrol.Init()
}
