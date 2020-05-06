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
	fmt.Println(people)

	bs, err := json2.Marshal(people)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Marshalled data is", string(bs))
	}

	s := string(bs)
	bs2 := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs2)

	var people2 []person

	err = json2.Unmarshal(bs2, &people2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Unmarshalled data is", people2)
	for i, v := range people2 {
		fmt.Println("Person", i, ":", v)
	}
}
