package main

import "fmt"

func main() {
	g := 1 == 2
	h := 1 <= 2
	i := 1 >= 2
	j := 1 != 2
	k := 1 < 2
	l := 1 > 2

	fmt.Printf("%t\n%t\n%t\n%t\n%t\n%t", g, h, i, j, k, l)
}
