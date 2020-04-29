package main

import "fmt"

func main() {
	mn := "Yu"
	switch {
	case mn == "Yu":
		fmt.Println("My name is Yu")
	case mn == "Jay":
		fmt.Println("Jay")
	case mn == "Hokiat":
		fmt.Println("My name is Hokiat")
	default:
		fmt.Println("no name found")
	}
}
