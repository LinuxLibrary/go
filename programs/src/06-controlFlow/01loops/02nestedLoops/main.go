package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Printf("Outer loop i: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tInner loop j: %d\n", j)
		}
	}
}
