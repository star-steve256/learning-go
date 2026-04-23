package arraysslicesstructs

import "fmt"

func Init() {
	initArr()
}

func initArr() {
	// declaration using literal
	var fixedSizeArr = [2]int{1, 2}

	// absolute value copy vs pointer reference

	fmt.Println("Fixed size arr:", fixedSizeArr)
	// this actually creates copy of fixedSizeArr
	var referencedArr = fixedSizeArr
	referencedArr[0] = 1000000

	fmt.Println("After referencing and changing value: ")
	fmt.Println("Fixed size arr:", fixedSizeArr)
	fmt.Println("Referenced arr:", referencedArr)

	var secondReferenceArr = &fixedSizeArr
	secondReferenceArr[0] = 20000000
	fmt.Println("After referencing and changing value via pointer: ")
	fmt.Println("Fixed size arr:", fixedSizeArr)
	fmt.Println("2nd Referenced arr:", *secondReferenceArr)

	// SLICES - variable size arrays
	var stringSlice = []string{"hello", "world"}

	// can append like this
	stringSlice = append(stringSlice, "Hi")
	fmt.Println("stringSlice after appending:", stringSlice)

	// make is used to initialize slices, maps, and channels
	// example-
	// s := make([]int, 5) // <- makes int slice of 5 zero values
	/* So what's stopping me from do this? :)
	   s:= []int{0, 0, 0, 0, 0} */
	/* Because make makes it (no pun) more readable and more intentional about
	   initialization */
	stringSliceCopy := make([]string, len(stringSlice))
	copy(stringSliceCopy, stringSlice)
	fmt.Println("stringSliceCopy:", stringSliceCopy)

	// for loops can be used freely as other languages
	var fruitArr = []string{"apple", "banana", "orange"}

	// loop by index
	for i := 0; i < len(fruitArr); i++ {
		fmt.Print(fruitArr[i] + " ")
	}
	fmt.Println()
	// loop by range and value
	for _, fruit := range fruitArr {
		fmt.Print(fruit + " ")
	}
	fmt.Println()
}

func initStructs() {
	// for another day ig? : )
}
