package main

import "fmt"

func main() {

	v1 := struct {
		colors map[string]string
		dm     []int
	}{
		colors: map[string]string{
			"door": "white",
			"body": "black",
		},
		dm: []int{2, 5816},
	}

	for key, value := range v1.colors {
		fmt.Println(key, value)
	}
	for _, v := range v1.dm {
		fmt.Println(v)
	}
}
