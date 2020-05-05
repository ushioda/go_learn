package main

import "fmt"

func g() func() int {
	f := func() int {
		return 42
	}
	return f
}

func main() {
	f := g()
	fmt.Println(f())
}
