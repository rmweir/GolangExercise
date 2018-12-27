package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	j := `[{"First":"Bob","Last":"the Builder","Age":32,"Sayings":["Lets build.", "I love tractors!","Hammers are good."]},{"First":"Hermione","Last":"Granger","Age":23,"Sayings":["It's leviosa not leviosar", "My parents are dentists"]},{"First":"Ron","Last":"Weasley","Age":22,"Sayings":["Bloody hell!","Did you see that?","My father works for the ministry"]}]`
	byteform := []byte(j)
	fmt.Println(j)
	people := []person{}
	err := json.Unmarshal(byteform, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
