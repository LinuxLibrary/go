package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martini", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being Evil", "Ice Cream", "Sunsets"},
	}

	fmt.Println("Adding record for Arjun")
	m["Arjun"] = []string{"Music", "Movies", "Blogging", "Learning"}
	
	fmt.Println("Removing record for 'Dr. No'")
	delete(m, "no_dr")
	
	for k, v := range m {
		fmt.Println(k)
		for i, s := range v {
			fmt.Printf("\tIndex: %v\tValue: %v\n", i, s)
		}
	}
}
