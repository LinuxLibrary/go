package main

import (
	"fmt"
)

func main() {
	fmt.Println("APPENDING ELEMENTS INTO A SLICE")

	x := []int{1, 2, 3, 4, 5, 34, 54, 65, 46}
	fmt.Println(x)

	x = append(x, 44, 55, 66, 77)
	fmt.Println(x)

	y := []int{123, 234, 345, 456}
	fmt.Println(y)

	fmt.Println("x = append(x, y) : This is not possible. We can't directly append an entire slice into another slice")
	fmt.Println("We need to take the elements from the slice 'y' and append to slice x --> 'x = append(x, y...)'")
	x = append(x, y...)
	fmt.Println(y)
	fmt.Println(x)
}
