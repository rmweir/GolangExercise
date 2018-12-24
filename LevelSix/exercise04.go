package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main() {
	v := person {
		first: "Dr",
		last: "Seuss",
		age: 88,
	}
	v.speak()
}

func (per person) speak() {
	fmt.Println("Name:", per.first, per.last)
	fmt.Println("Age:", per.age)
}
