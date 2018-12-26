package main

import "fmt"

func main() {
	x := func() {
		for i := 0; i < 20; i++ {
			fmt.Println(i)
		}
	}
	x()
}
