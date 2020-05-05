package main

import "fmt"

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func main() {
	x := []int{10, 9, 8, 7}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}
