package main

import (
	"fmt"
)

func main() {
	x := []string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Connecticut`, `Delware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minesotta`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dacota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvenia`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennesse`, `Texas`, `Utah`, `Vermont`, `Verginia`, `Washington`, `West Verginia`, `Wiskonsin`, `Wyoming`}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
