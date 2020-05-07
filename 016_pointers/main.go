package main

import "fmt"

func main() {
	y := 0
	fmt.Println(&y)
	fmt.Println(y)
	foo(&y)
	fmt.Println(&y)
	fmt.Println(y)
}

func foo(x *int) {
	fmt.Println(x)
	fmt.Println(*x)
	*x = 43
	fmt.Println(x)
	fmt.Println(*x)
}
