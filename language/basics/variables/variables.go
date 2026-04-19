package variables

import "fmt"

/* this is package-scoped variable
   (global var of this package) */
var PackageScopeVar = "apple"

// declaration
var newWord string
// declaration + assignment (allowed only in executable block)
// newWord2 := "Go" // <- gives compile time error


func MainRun() {
	subRun1()
	fmt.Println("Value of PackageScopeVar in MainRun():", PackageScopeVar)
	subRun2()
	fmt.Println("Value of PackageScopeVar in MainRun():", PackageScopeVar)
	subRun3()
}

func subRun1() {
	/* this does not change global value
	   the value lives only inside this func */
	PackageScopeVar := "orange"
	fmt.Println("Value of PackageScopeVar in subRun1():", PackageScopeVar)
}

func subRun2() {
	/* if it is changed with assignment without declaration
	   it affects the global value */
	PackageScopeVar = "banana"
	fmt.Println("Value of PackageScopeVar in subRun2():", PackageScopeVar)
}

func subRun3() {
	var emptyBool bool
	var emptyString string
	var emptyInt int

	/* values of unassigned variables are not nil
	   they are default zero-values of that variable type
	   string = "", int = 0, bool = 0 etc */
	fmt.Println("Empty values:", emptyString, emptyInt, emptyBool)
	
	/* type conversions
	   simple T(v) the value */
	var newInt = 98
	var floatOfNewInt = float64(newInt)
	fmt.Printf("floatOfNewInt: %.2f %T\n", floatOfNewInt, floatOfNewInt)
	
	// const: denotes fixed value
	// NOTE: Go does not give compile time error for unused const
	const fixedValue = 5
}
