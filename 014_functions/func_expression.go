package main

import "fmt"

func hi() string {
	s := "Hello World"
	return s
}

func main() {
	f := func() {
		fmt.Println("my first function expression")
	}
	f2 := func(x int) {
		fmt.Println("second func expression with", x)
	}
	f()
	f2(42)

	s1 := hi()
	fmt.Println(s1)
}
