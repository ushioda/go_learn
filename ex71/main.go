package main

import "fmt"

var x int

func main() {
	x = 13
	fmt.Println(&x)
}