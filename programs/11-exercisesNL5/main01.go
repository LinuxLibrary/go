package main

import (
	"fmt"
)

type person struct {
	first              string
	last               string
	favIceCreamFlavors []string
}

func main() {
	p1 := person{
		first: "Arjun",
		last:  "Shrinivas",
		favIceCreamFlavors: []string{
			"ButterScotch",
			"Venella",
		},
	}

	p2 := person{
		first: "Omkar",
		last:  "Ganesh",
		favIceCreamFlavors: []string{
			"Venella",
			"ButterScotch",
			"Chocolate",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favIceCreamFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favIceCreamFlavors {
		fmt.Println(i, v)
	}
}
