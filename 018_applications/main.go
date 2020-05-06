package main

import (
	json2 "encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Kim",
		Last:  "Philby",
		Age:   54,
	}

	people := []person{p1, p2}
	bs, err := json2.Marshal(people)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}
}
