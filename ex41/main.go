package main

import "fmt"

func main() {
	a := []int{23, 34, 23, 12, 55, 565, 324, 876, 45, 23}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T \n", a)
}
