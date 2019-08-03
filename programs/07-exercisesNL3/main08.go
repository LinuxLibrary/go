package main

import (
	"fmt"
)

func main() {
	switch {
	case (2 == 4):
		fmt.Println("Not true")
		fallthrough
	case (2 == 2):
		fmt.Println("This is true")
	default:
		fmt.Println("This is default!")
	}
}
