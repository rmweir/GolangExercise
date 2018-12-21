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
	fmt.Printf("%s %s's favorite ice cream flavors:\n", b.first, b.last)
	for _, v := range b.flavors {
		fmt.Printf("\t%s\n", v)
	}

	fmt.Printf("%s %s's favorite ice cream flavors:\n", c.first, c.last)
	for _, v := range c.flavors {
		fmt.Printf("\t%s\n", v)
	}

}
