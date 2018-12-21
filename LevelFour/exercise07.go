package main

import "fmt"

func main() {
	a := [][]string{{"James", "bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooooo, James"}}
	fmt.Println(a)
	for _, v := range a {
		for _, k := range v {
			fmt.Println(k)
		}
	}
}
