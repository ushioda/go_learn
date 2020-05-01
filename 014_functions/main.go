package main

import "fmt"

func main() {
	s := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is", s)
	bar("James")
	s1 := woo("Money")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ` says "Hello"`)
	b := false
	return a, b
}
