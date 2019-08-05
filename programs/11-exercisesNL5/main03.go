package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxary bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Blue",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Black",
		},
		luxary: true,
	}

	fmt.Println(t1.doors, t1.color, t1.fourWheel)
	fmt.Println(s1.doors, s1.color, s1.luxary)
}
