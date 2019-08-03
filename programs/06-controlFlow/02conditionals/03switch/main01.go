package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("This won't print")
	case (2 == 4):
		fmt.Println("This should not print")
	case (2 == 2):
		fmt.Println("This prints")
		fallthrough
	case (3 == 3):
		fmt.Println("This is true, Does it print?")
		fallthrough
	case (7 == 9):
		fmt.Println("Not true")
		fallthrough
	case (11 == 12):
		fmt.Println("Not true 2")
	default:
		fmt.Println("This is default")
	}
}
