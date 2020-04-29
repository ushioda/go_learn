package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("The remainder of %d divided by 4 is %d \n", i, i%4)
		} else if i == 10 {
			fmt.Println("This is the first loop!")
		} else {
			fmt.Println("I'm an odd number, but not one!")
		}
	}
}
