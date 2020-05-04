package main

import "fmt"

func ref() func() int {
	return func() int {
		return 451
	}
}

func main() {
	fmt.Println(ref()())
}
