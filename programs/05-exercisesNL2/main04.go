package main

import (
	"fmt"
)

func main() {
	fmt.Println("Declare number print it, shift left by one digit declare it to another and print that.")
	a := 33
	fmt.Printf("%d\t%b\t%#X\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#X\n", b, b, b)
}
