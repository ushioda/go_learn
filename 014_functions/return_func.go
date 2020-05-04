package main

import "fmt"

func ref() func() int {
	return func() int {
		return 451
	}
}

func main() {
	x := ref()
	fmt.Printf("This is type %T \n", x)
	i := ref()
	fmt.Println(i())
}
