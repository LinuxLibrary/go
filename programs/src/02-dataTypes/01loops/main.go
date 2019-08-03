package main

import "fmt"

func main() {
	fmt.Println("Hello Everyone",
		"Hakuna Matata!")
	foo()
}

func foo() {
	fmt.Println("I'm in foo!")

	fmt.Println("Lets try Iterating / Looping")
	fmt.Println("Printing even numbers...")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
