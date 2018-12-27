package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "Bob",
		Age:   19,
	}
	u2 := user{
		First: "Linda",
		Age:   27,
	}
	users := []user{u1, u2}
	fmt.Println(users)

	usersjson, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(usersjson))
}
