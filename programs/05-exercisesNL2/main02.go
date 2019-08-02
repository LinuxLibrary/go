package main

import (
	"fmt"
)

func main() {
	a := (42 == 42)
	b := (42 >= 42)
	c := (42 <= 42)
	d := (42 != 42)
	e := (42 > 42)
	f := (42 < 42)
	fmt.Printf("Boolean value of a: %v\n", a)
	fmt.Printf("Boolean value of b: %v\n", b)
	fmt.Printf("Boolean value of c: %v\n", c)
	fmt.Printf("Boolean value of d: %v\n", d)
	fmt.Printf("Boolean value of e: %v\n", e)
	fmt.Printf("Boolean value of f: %v\n", f)
}
