package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hakuna matata", 3, true)
	fmt.Println(n)
	fmt.Println(err)
}
