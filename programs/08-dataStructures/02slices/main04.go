package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 34, 54, 65, 46}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:8])
	fmt.Println(x[2:7])
	fmt.Println(x[3:6])
	fmt.Println(x[4:5])
}
