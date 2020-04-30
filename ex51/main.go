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

	fic2 := []string{"Coffee", "Strawberry"}

	p2 := person{
		first: "Michael",
		last:  "Kim",
		fic:   fic2,
	}

	pm := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for key, value := range pm {
		fmt.Println(value.first)
		fmt.Println(key)
		for i, v := range value.fic {
			fmt.Println(i, v)
		}
	}
}
