package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("true")
	}

	if false {
		fmt.Println("false")
	}

	if !true {
		fmt.Println("not true")
	}

	if !false {
		fmt.Println("not false")
	}
}
