package main

import (
	"fmt"
)

// Variables of global scope can be declared in this way using the var keyword
// Var type can also be specified while declaring
// By default it assigns 0 for int, False for boolean and "" for strings
//	if nothing is specified
var y = 43
var z int
var a string
var b bool

func main() {
	// Short declaration operator
	// Declare a variable and assign value (of a certain type)
	// This variable serves as local variable, which limits to work within the function
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)
}
