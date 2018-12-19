package main

import "fmt"

func main() {
	favSport := "Volleyball"
	switch favSport {
	case "Volleyball":
		fmt.Println("Fav sport is volleyball.")
	case "Basketball":
		fmt.Println("Fav sport is basketball.")
	default:
		fmt.Println("Fav sport is not volleyball or basketball.")
	}
}
