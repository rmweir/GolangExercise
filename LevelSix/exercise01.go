package main

import "fmt"

func main() {
	fmt.Println(foo())
	k, v := bar()
	fmt.Println(k, v)
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 3, "Hello"
}
