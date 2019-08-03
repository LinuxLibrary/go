package main

import (
	"fmt"
)

func main() {
	switch "Bond" {
	case "MoneyPenny":
		fmt.Println("Miss Penny")
	case "Bond":
		fmt.Println("James Bond")
	case "Q":
		fmt.Println("This is Q")
	default:
		fmt.Println("This is default!")
	}
}
