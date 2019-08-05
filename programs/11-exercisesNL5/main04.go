package main

import (
	"fmt"
)

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"M":          666,
			"Q":          777,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}

	fmt.Println(s.first)
	for k, v := range s.friends {
		fmt.Println(k, v)
	}
	for i, v := range s.favDrinks {
		fmt.Println(i, v)
	}
}
