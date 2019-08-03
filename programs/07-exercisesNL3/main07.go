package main

import (
	"fmt"
)

func main() {
	if 2 == 3 {
		fmt.Println("Not true")
	} else if 2 == 4 {
		fmt.Println("Also not true")
	} else {
		fmt.Println("Got you finally!")
	}
}
