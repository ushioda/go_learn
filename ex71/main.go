package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (pa *person) changeMe(p person) {
	(*pa).first = p.first
	(*pa).age = p.age
}

func changeMe2(pa *person) {
	(*pa).first = "Michael"
	(*pa).age = 48
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p1)
	p2 := person{
		first: "Kim",
		last:  "Philby",
		age:   51,
	}
	p1.changeMe(p2)
	fmt.Println(p1)
	changeMe2(&p1)
	fmt.Println(p1)
}
