package main

import "fmt"

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
	luxury bool
}

func main() {
	a := truck{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		fourWheel: true,
	}

	b := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "gray",
		},
		luxury: false,
	}

	fmt.Println("Truck specifications:")
	fmt.Printf("\tDoors: %d\n", a.doors)
	fmt.Printf("\tColor: %s\n", a.color)
	fmt.Printf("\tFour Wheel: %t\n", a.fourWheel)

	fmt.Println("Sedan specifications:")
	fmt.Printf("\tDoors: %d\n", b.doors)
	fmt.Printf("\tColor: %s\n", b.color)
	fmt.Printf("\tLuxury: %t\n", b.luxury)
}
