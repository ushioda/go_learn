package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I am", sa.first, sa.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Kim",
			last:  "Philby",
		},
		ltk: false,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
