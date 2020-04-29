package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Bob":   40,
	}

	fmt.Println(m)

	if v, ok := m["Sam"]; ok {
		fmt.Printf("%v exists and delete \n", v)
		delete(m, "Sam")
	} else {
		fmt.Printf("%v did not exist \n", v)
	}
	if v, ok := m["James"]; ok {
		fmt.Printf("%v exists and delete \n", v)
		delete(m, "James")
	} else {
		fmt.Printf("%v did not exist \n", v)
	}

	fmt.Println(m)
}
