package arraysslicesstructs

import "fmt"

func Init() {
	initArr()
	initStructs()
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
	// structs serve the purpose of objects (kinda?)
	
	// defines a type
	type Fruit struct {
		Name string
		Color string
		Flavor string
	}

	// struct init
	var apple1 = Fruit{"apple", "red", "sweet"}
	var orange1 = Fruit{"orange", "orange", "sweet"}
	fmt.Println("Apple:", apple1)
	fmt.Println("Orange:", orange1)

	// use property syntax for structs
	fmt.Println("Flavour of apple1:", apple1.Flavor)

	// let's try referencing pointer here too : )

	var apple2 = &apple1

	apple2.Color = "apple 2 color"
	// it changes as in referenced value
	fmt.Println("Apple1 color:", apple1.Color)
	fmt.Println("Apple2 color:", apple2.Color)

	*apple2 = apple1
	apple2.Color = "new new apple 2 color"
	// after de-referencing and copying
	fmt.Println("Apple1 color:", apple1.Color)
	fmt.Println("Apple2 color:", apple2.Color)

	// uses new to allocate new memory
	apple2 = new(Fruit)
	apple2.Color = "third change"
	// after de-referencing and copying
	fmt.Println("Apple1 color:", apple1.Color)
	fmt.Println("Apple2 color:", apple2.Color)
}
