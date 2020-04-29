package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}
