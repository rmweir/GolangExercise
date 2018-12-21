package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {
	b := person{
		first:   "Bob",
		last:    "Builder",
		flavors: []string{"hard work", "chocolate"},
	}

	c := person{
		first:   "Grace",
		last:    "Hopper",
		flavors: []string{"mint", "pistachio"},
	}

	people := map[string]person{}
	people["Builder"] = b
	people["Hopper"] = c

	for _, val := range people {
		fmt.Printf("%s %s's favorite ice cream flavors:\n", val.first, val.last)
		for _, v := range val.flavors {
			fmt.Printf("\t%s\n", v)
		}
	}
}
