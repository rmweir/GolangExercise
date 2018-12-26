package main

import "fmt"

func execute(f func() string) {
	fmt.Println(f())
}

func callback() string {
	return "Echo"
}

func main() {
	execute(callback)
}
