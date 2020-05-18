package main

import (
	"fmt"
	"go_learn/026_godocs/ushiodaMath"
)

func main() {
	a := 2.2
	b := 3.3
	c := ushiodaMath.Sum(a, b)
	fmt.Printf("The sum of %v and %v is %v \n", a, b, c)
}
