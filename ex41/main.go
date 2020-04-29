package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "Alice"
	a[1] = "Bob"
	a[2] = "Chris"
	for i, v := range a{
		fmt.Println(i, v)
	}
	fmt.Printf("%T \n", a)
}
