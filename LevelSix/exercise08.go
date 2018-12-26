package main

import "fmt"

func main() {
	x := hello()
	fmt.Println(x())
}

func hello() func() int {
	return func() int {
		return 32
	}
}
