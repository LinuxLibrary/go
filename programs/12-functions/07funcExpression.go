package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("Printing from funcExpression...")
	}

	f()

	g := func(x int) {
		fmt.Println("funcExpression with parameter", x)
	}

	g(12)
}
