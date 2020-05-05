package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorial2(n int) int {
	ans := 1
	for ; n > 0; n-- {
		ans *= n
	}
	return ans
}

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorial2(4))
}
