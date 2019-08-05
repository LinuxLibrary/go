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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, v := range v.favIceCreamFlavors {
			fmt.Println(i, v)
		}
	}
}
