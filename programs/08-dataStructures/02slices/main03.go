package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 34, 54, 65, 46}
	fmt.Println(len(x))
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}
}
