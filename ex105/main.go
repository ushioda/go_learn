package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println("check 1", v, ok)
	v, ok = <-c
	fmt.Println("check 2", v, ok)
}
