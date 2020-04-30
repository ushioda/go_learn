package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Ethan",
		last:  "Hawk",
		age:   55,
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
