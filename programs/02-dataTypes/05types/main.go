package main

import (
	"fmt"
)

var y = 42
var z string = "Hello, Bingo"
var a string = `Dev said, "This is string literral(which means print everything in the backticks)"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
