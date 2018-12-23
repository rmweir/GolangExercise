package main

import "fmt"

func main() {
	defer delay()
	fmt.Println("I will print first despite my position")
}

func delay() {
	fmt.Println("I have been deferred")
}
