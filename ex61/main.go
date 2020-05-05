package main

import "fmt"

func foo() (int, string) {
	return 0, "James"
}

func bar() int {
	x := 0
	return x
}

func main() {
	fmt.Println(bar())
	fmt.Println(foo())
}
