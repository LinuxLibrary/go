package main

import (
	"fmt"
)

func main() {
	x := 42
	if x == 40 {
		fmt.Println("x is 40")
	} else if x == 41 {
		fmt.Println("x is 41")
	} else if x == 42 {
		fmt.Println("x is 42")
	} else {
		fmt.Println("x is not between 40 and 42")
	}
}
