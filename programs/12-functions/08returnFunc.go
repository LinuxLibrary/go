package main

import (
	"fmt"
)

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println("The value of i is: ", i)

	fmt.Println("The value of x() is: ", x())
}

func foo() string {
	s := "Hello, World!"
	return s
}

func bar() func() int {
	return func() int {
		return 200
	}
}
