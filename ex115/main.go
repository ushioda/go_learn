package main

import (
	"fmt"
	"go_learn/ex115/myPacks"
)

func main() {
	a := 2.2
	b := 3.3
	c := myPacks.Sum(a, b)
	fmt.Printf("The sum of %v and %v is %v \n", a, b, c)
}
