package main

import (
	"fmt"
)

func main() {
	fmt.Println("DECLARING A SLICE")

	x := []int{1, 2, 3, 4, 5, 34, 54, 65, 46}
	fmt.Println(x)

	fmt.Println("\nAPPENDING ELEMENTS INTO A SLICE")
	x = append(x, 44, 55, 66, 77)
	fmt.Println(x)

	fmt.Println("\nAPPENDING ELEMENTS OF A SLICE INTO ANOTHER SLICE")
	y := []int{123, 234, 345, 456}
	x = append(x, y...)
	fmt.Println(y)
	fmt.Println(x)

	fmt.Println("\nDELETING ELEMENTS FROM A SLICE")
	x = append(x[:5], x[9:]...)
	fmt.Println(x)
}
