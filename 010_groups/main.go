package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Bob":   40,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Bob"])
	fmt.Println(m["Sam"])

	v, ok := m["Sam"]
	fmt.Println(v)
	fmt.Println(ok)

	if _, ok := m["Sam"]; ok {
		println("Sam exists")
	} else {
		println("Sam does not exist")
	}

	for key, value := range m{
		println(key, value)
	}

}
