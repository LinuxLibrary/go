package main

import (
	"fmt"
)

const (
	a int     = 48
	b float64 = 42.78
	c string  = "Hakuna Matata"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
