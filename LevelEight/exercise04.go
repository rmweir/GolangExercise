package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 1}
	xs := []string{"random", "rainbow", "delights", "int", "torpedo", "summers"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
