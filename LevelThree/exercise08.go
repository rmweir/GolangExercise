package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("False.")
	case 1 == 1:
		fmt.Println("1 does equal 1.")
	case 2 == 3:
		fmt.Println("This is not true.")
	}
}
