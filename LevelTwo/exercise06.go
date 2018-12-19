package main

import "fmt"

const (
	yearOne = 2018 - iota
	yearTwo = 2018 - iota
	yearThree = 2018 - iota
	yearFour = 2018 - iota
)
func main() {
	fmt.Printf("%d\n%d\n%d\n%d", yearOne, yearTwo, yearThree, yearFour)
}
