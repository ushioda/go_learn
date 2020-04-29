package main

import "fmt"

func main() {
	for x := 33; x < 123; x++ {
		fmt.Printf("ASCII value %d corresponds to %#U\n", x, x)
		}
	}
