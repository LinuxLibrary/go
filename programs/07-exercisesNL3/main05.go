package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		rem := i % 4
		if rem != 0 {
			fmt.Printf("%d\t%d\n", i, rem)
		}
	}
}
