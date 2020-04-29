package main

import "fmt"

func main() {
	for x := 0; x < 20; x++ {
		if x%2 == 0 {
			fmt.Println(x)
		}
	}
}
