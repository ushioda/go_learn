package main

import "fmt"

func sum2(xi ...int) int {
	total := 0
	for _, val := range xi {
		total += val
	}
	return total
}

func even(f func(vi ...int) int, xi ...int) int {
	var evens []int
	for _, val := range xi {
		if val%2 == 0 {
			evens = append(evens, val)
		}
	}
	return f(evens...)
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum2(x...))
	fmt.Println(even(sum2, x...))
}
