package main

import (
	"fmt"
)

func main() {
	foo()
	bar("Geek!")
	s1 := woo("Moneypenny")
}

// func (r receiver) identifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("Hello, Hakuna Matata")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) {
	return fmt.Sprint("Hello from woo,", st)
}