package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p *person) speak() {
	fmt.Println("I am speaking!")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		Name: "James",
		Age:  32,
	}

	//saySomething(human(p1))
	saySomething(human(&p1))

}
