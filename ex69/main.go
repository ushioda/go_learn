package main

import "fmt"

func double(x float64) float64 {
	return 2 * x
}

func apply(f func(float64) float64, x []float64) float64 {
	total := 0.0
	for _, v := range x {
		total += f(v)
	}
	return total
}

func main() {
	a := []float64{2.1, 3.3}
	fmt.Println(apply(double, a))
}
