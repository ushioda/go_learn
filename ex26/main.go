package main

import "fmt"

const (
	year1 = 2020 - iota
	year2 = year1 - iota
	year3 = year1 - iota
	year4 = year1 - iota
)

func main() {
	fmt.Printf("%d \t %d \t %d \t %d \t", year1, year2, year3, year4)
}
