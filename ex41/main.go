package main

import "fmt"

func main() {
	a := []int{23, 34, 23, 12, 55, 565, 324, 876, 45, 23}
	length := 5
	for i := 0; i < len(a)-length+1; i++ {
		fmt.Println(a[i : i+length])
	}
}
