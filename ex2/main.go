package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	s := fmt.Sprintf("%v\t %v\t %v\t",x, y, z)
	fmt.Println(s)
}
