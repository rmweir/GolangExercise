
package main

import "fmt"

func main() {
	a := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"}}
	a["strange_dr"] = []string{"medicine", "money", "defying physics"}
	for key, val := range a {
		fmt.Printf("%s's favorite things:\n", key)
		for _, fav := range val {
			fmt.Printf("\t%s\n", fav)
		}
		fmt.Println()
	}
}
