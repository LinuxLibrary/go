package main

import (
	"fmt"
)

func main() {
	fmt.Println("n, e := fmt.Println('Unused variable declaration is not possible in Go')")
	fmt.Println("If you are not using the variable 'e' then you can't declare it!")
	fmt.Println("!!! In such case you need to use '_' for unused variables !!!")
	n, _ := fmt.Println("Hakuna matata", 4, true)
	fmt.Println(n)
}
