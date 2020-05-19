// Package myMAth2 performs basic math operations
package myMath2

// Sum adds an unlimited number of values of int
func Sum(xi ...int) int {
	s := 0
	for _, x := range xi {
		s += x
	}
	return s
}
