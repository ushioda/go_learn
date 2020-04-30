package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Ethan",
			last:  "Hawk",
			age:   55,
		},
		ltk: false,
	}
	fmt.Println(sa1.first)
	fmt.Println(sa1.last)
	fmt.Println(sa1.age)
	fmt.Println(sa1.ltk)
	fmt.Println(sa2)
}
