package main

import "fmt"

func main() {
	fmt.Println(bar([]int{1, 2, 3}))
}

func foo(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func bar(nums []int) int {
	return foo(nums...)
}
