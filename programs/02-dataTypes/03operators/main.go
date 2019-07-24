package main

import (
	"fmt"
)

func main() {
	fmt.Println("To declare a variable you need to use := operator")
	x := 9
	fmt.Println("The value of x is: ", x)
	fmt.Println("But if you want to reassign a new value, you need to use = operator")
	x = 99
	fmt.Println("The value of x now is: ", x)
}
