package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("type of x: %T\ntype of y: %T\ntype of z: %T\n", x, y, z)
	// the values printed will be each types "Zero Value"
	fmt.Println(x, y, z)
}
