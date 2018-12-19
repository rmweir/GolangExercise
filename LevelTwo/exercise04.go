package main

import "fmt"

func main() {
	x := 90
	fmt.Printf("%d\n%b\n%#X", x, x, x)
	y := x << 1
	fmt.Printf("%d\n%b\n%#X", y, y, y)
}
