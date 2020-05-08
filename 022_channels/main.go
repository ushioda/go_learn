package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	go func() {
		c <- 42
	}()
	//fmt.Println(<-c)
	fmt.Println("--------------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	cr = c // from general to specific works
	fmt.Println(<-cr)
}
