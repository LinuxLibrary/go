package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])

	fmt.Println("If we are trying to print a value of a key that doesn't exists, by default we get 0 as the value.")
	fmt.Println(m["HakunaMatata"])

	fmt.Println("In this case we can check the existance of the key through ', ok'")
	v, ok := m["HakunaMatata"]
	fmt.Println(v)
	fmt.Println(ok)

	fmt.Println("We can check this using the If condition as well")
	if v, ok := m["HakunaMatata"]; ok {
		fmt.Println(v)
	}

	fmt.Println("Adding a new element to the map")
	m["Arjun"] = 32

	fmt.Println("Loop over the map and get the contents using the range")
	for k, v := range m {
		fmt.Printf("NAME: %s\n\tAGE:%d\n", k, v)
	}

	fmt.Println("Delete an element from a map")
	fmt.Println("Delete using the key: 'delete(<map name>, \"key\")'")
	fmt.Println(m)
	delete(m, "Miss Moneypenny")
	fmt.Println(m)

	for k, v := range m {
		if _, ok := m[k]; ok {
			fmt.Printf("Deleting %s\t%d\n", k, v)
			delete(m, k)
		}
	}

	fmt.Println(m)
}
