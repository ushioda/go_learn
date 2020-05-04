package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("my first function expression")
	}
	f2 := func(x int) {
		fmt.Println("second func expression with", x)
	}
	f()
	f2(42)
}
