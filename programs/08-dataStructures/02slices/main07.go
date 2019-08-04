package main

import (
	"fmt"
)

func main() {
	fmt.Println("make() is a builtin function through which we can create a sclice and can manage the underlying array size")
	fmt.Println("x := make([]Type, LengthOfSlice, SizeOfArray)")
	fmt.Println("x := make([]int, 10, 100)\n")

	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Println("This slice can take upto 10 elements with an array of size 12")
	fmt.Println("If we append that slice beyond that size then the size of the underlying array will be doubled\n")

	y := []int{11, 12, 13, 14}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
