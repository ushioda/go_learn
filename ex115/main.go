package main

import "fmt"

func main() {
	a := 2.2
	b := 3.3
	c := Sum(a, b)
	fmt.Printf("The sum of %v and %v is %v \n", a, b, c)
}

func Sum(m, n float64) float64 {
	return m + n
}