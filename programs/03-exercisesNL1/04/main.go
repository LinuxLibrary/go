package main

import (
	"fmt"
)

type custType int

var x custType

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}