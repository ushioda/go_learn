package main

import "fmt"

func baseCounter() func() int {
	n := 0
	return func() int {
		n += 1
		return n
	}
}

func main() {
	counter := baseCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
