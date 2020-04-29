package main

import "fmt"

func main() {
	fr := []string{"James", "Bond", "Soccer"}
	sr := []string{"Alice", "Smith", "Judo"}
	fmt.Println(fr)
	fmt.Println(sr)
	xp := [][]string{fr, sr}
	fmt.Println(xp)
}
