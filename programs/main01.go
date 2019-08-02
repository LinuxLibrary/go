package main

import (
	"fmt"
)

var a int = 47

func main() {
	fmt.Println("Print a number in Decimal, Binary and HexaDecimal")
	fmt.Printf("Decimal value of a: %v\n", a)
	fmt.Printf("Binary value of a: %b\n", a)
	fmt.Printf("HexaDecimal value of a: %#X\n", a)
}
