package main

import (
	"fmt"
)

func main() {
	// x := type{values}
	// composite literal
	// A SLICE allows you to group togather the values of the same type
	x := []int{1, 2, 3, 4, 5, 34, 54, 65, 46}
	fmt.Println(x)
	fmt.Println(len(x))
}
