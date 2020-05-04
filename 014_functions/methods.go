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
	fmt.Println("I am", sa.first, sa.last, "- secretAgent speaks")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- person speaks")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case secretAgent:
		fmt.Println("I am a human, by", h.(secretAgent).first)
	case person:
		fmt.Println("I am a human, by", h.(person).first)
	}
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
	p1 := person{
		first: "Jimmy",
		last:  "John",
	}
	sa1.speak()
	sa2.speak()
	bar(sa1)
	bar(sa2)
	bar(p1)
}
