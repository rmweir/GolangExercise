package main

import "fmt"

func main() {
	a := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"}}
	for key, val := range a {
		fmt.Println("Key:", key, "\tFav things:", val)
	}
}
