package main

import "fmt"

func closure() {
	x := 21
	{
		fmt.Println("Before initializing an inner x:", x)
		x := 30
		fmt.Println("within this scope the x value is:", x)
	}
	fmt.Println("within this outter scope the x calue is:", x)
}

func it() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
func main() {
	closure()
	g := it()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}
