package main

import (
	"fmt"
)

func main() {
	foo()

	func() {
		fmt.Println("Anonymous function printing....")
	}()

	func(x int) {
		fmt.Println("Anonymous function with parameter...", x)
	}(42)
}

func foo() {
	fmt.Println("Foo printing...")
}
