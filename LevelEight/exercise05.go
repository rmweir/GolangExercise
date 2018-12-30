package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByLast []user

func (a ByLast) Len() int {
	return len(a)
}

func (a ByLast) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByLast) Less(i, j int) bool {
	return a[i].Last < a[j].Last
}
func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation"},
	}
	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is so good to see you.",
			"I would really prefer to be a secret agent myself."},
	}

	users := []user{u1, u2}
	fmt.Println(users)

	sort.Sort(ByAge(users))
	fmt.Println(users)

	sort.Sort(ByLast(users))
	fmt.Println(users)

	for _, v := range users {
		fmt.Println(v.Sayings)
		sort.Strings(v.Sayings)
		fmt.Println(v.Sayings)
		fmt.Println()
	}
}
