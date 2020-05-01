package main

import "fmt"

func foo() {
	fmt.Println("Hi from foo")
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

func main() {
	foo()
	bar("James")
	s1 := woo("Money")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}
