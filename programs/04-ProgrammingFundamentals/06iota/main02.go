package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f
)

func main() {
	fmt.Printf("a: %v, Type: %T\n", a, a)
	fmt.Printf("b: %v, Type: %T\n", b, b)
	fmt.Printf("c: %v, Type: %T\n", c, c)
	fmt.Printf("d: %v, Type: %T\n", d, d)
	fmt.Printf("e: %v, Type: %T\n", e, e)
	fmt.Printf("f: %v, Type: %T\n", f, f)
}
