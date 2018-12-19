package main

import "fmt"

func main() {
	a := "kitty"

	if a == "dog" {
		fmt.Println("This will not print")
	} else if a == "kitty" {
		fmt.Println("Kittiesssssssss")
	} else {
		fmt.Println("This will not print")
	}
}
