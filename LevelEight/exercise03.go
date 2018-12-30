package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
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
	json.NewEncoder(os.Stdout).Encode(&users)
	fmt.Println(users)
}
