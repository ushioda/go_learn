package main

import (
	"fmt"
	"go_learn/028_test_1/myMath"
)

func main() {
	a := 11
	b := 22
	c := myMAth.Sum(a, b)
	fmt.Printf("Sum of %v and %v is %v \n", a, b, c)
}
