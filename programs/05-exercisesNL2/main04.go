package main

import (
	"fmt"
)

func main() {
	a := 33
	fmt.Printf("%d\t%b\t%#X\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#X\n", b, b, b)
}
