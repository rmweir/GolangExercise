package main

import "fmt"

func main() {
	year := 1990
	for {
		if year > 2018 {
			break
		} else {
			fmt.Println(year)
			year++
		}

	}
}
