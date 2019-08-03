package main

import (
	"fmt"
)

func main() {
	n := "Bond"
	switch n {
	case "MoneyPenny", "Bond", "Dr No":
		fmt.Println("Miss Penny or Bond or Dr No")
	case "M":
		fmt.Println("This is M")
	case "Q":
		fmt.Println("This is Q")
	default:
		fmt.Println("This is default!")
	}
}
