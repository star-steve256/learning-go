package functions

import "math"

// explicit arg types and return type
func Add(x int, y int) int {
	return x + y
}

// if types are same, can be compressed
func Subtract(x, y int) int {
	return x - y
}

// overloading is not allowed in Go
/*
func Subtract(x, y float) {

} */
// change the function name according to variation
func SubtractFloat(x, y float64) float64 {
	return x - y
}

// Go allows multiple return types
func swap(x string, y string) (string, string) {
	return y, x
}
// access by multiple declarations
// example: x, y := swap("world", "hello")

// Named return values: can return without specifying
func swapX(x string, y string) (a string, b string) {
	a = y
	b = x
	// a and b are declared so can return like this
	return
}


// functions are first-class in Go
// can be assigned to variables, pass as parameters etc
var FunctionOfAdding = Add

// callback style, function inside function
// simulation of f(g(x)) - square(add2(x))
func Add2(x int, f func(int) float64) float64 {
	x += 2
	return f(x)
}

func Square(x int) float64 {
	return math.Pow(float64(x), 2)
}
// call by add2(2, functions.Square)
