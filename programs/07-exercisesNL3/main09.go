package main

import (
	"fmt"
)

var a string = "favSport"

func main() {
	switch a {
	case "favMovie":
		fmt.Println("Wall-E")
	case "favSport":
		fmt.Println("Hockey")
	}
}
