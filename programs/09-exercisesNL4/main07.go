package main

import (
	"fmt"
)

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Hello, James"}

	a := [][]string{x, y}
	fmt.Println(a)

	for i, ps := range a {
		fmt.Println("Record: ", i)
		for j, cs := range ps {
			fmt.Printf("\tIndex Position: %v\tValue: %v\n", j, cs)
		}
	}
}
