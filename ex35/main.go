package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++{
		if i%2 == 0{
			fmt.Printf("The remainder of %d divided by 4 is %d \n", i, i%4)
		}
	}
}

