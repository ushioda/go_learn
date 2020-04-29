package main

import "fmt"

func main() {
	favSport := "Tennis"
	switch favSport {
	case "Soccer":
		fmt.Println("He likes soccer")
	case "Baseball":
		fmt.Println("He likes baseball")
	case "Tennis":
		fmt.Println("He likes tennis")
	default:
		fmt.Println("He doesn't like sports")
	}
}
