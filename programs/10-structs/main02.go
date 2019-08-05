package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	lkt bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		lkt: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
			age:   27,
		},
		lkt: false,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println("Child values of the Value 'sa1' can be directly accessed as 'sa1.first' and no need of mentioning the child while calling")
	fmt.Println("If we have name collision in the same value then we need to mention the complete order starting from parent to grandchild")
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.lkt)
	fmt.Println(sa2.first, sa2.last, sa2.age, sa2.lkt)
}
