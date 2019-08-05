package main

import (
	"fmt"
)

func main() {
	fmt.Println(`If we want to use structs which will be used only once in the entire programs,
there is no need of creating those structs. Those can be declared anonymously without giving a name`)
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)
}
