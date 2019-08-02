package main

import (
	"fmt"
)

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("Decimal:\t%v\n", n)
	fmt.Printf("Binary:\t\t%b\n", n)
	fmt.Printf("HexaDecimal:\t%#X\n", n)
}
