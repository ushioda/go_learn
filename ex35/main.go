package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++{
		fmt.Printf("The remainder of %d divided by 4 is %d \n", i, i%4)
	}
}

