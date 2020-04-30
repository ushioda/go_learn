package main

import "fmt"

func main() {
	type person struct {
		first string
		last  string
		fic   []string
	}

	fic1 := []string{"Vanilla", "Chocolate"}

	p1 := person{
		first: "Jenny",
		last:  "Zhang",
		fic:   fic1,
	}

	for i, v := range p1.fic {
		fmt.Println(i, v)
	}

	fic2 := []string{"Coffee", "Strawberry"}

	p2 := person{
		first: "Michael",
		last:  "Kim",
		fic:   fic2,
	}
	for i, v := range p2.fic {
		fmt.Println(i, v)
	}
}
