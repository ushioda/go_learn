package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 26; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t %#U \n", i + 65)
		}
	}
}
