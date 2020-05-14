package main

import (
	"fmt"
)

func main() {

	q := make(chan bool)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- bool) <-chan int {

	c := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		q <- true
		close(c)
		close(q)
	}()

	return c
}

func receive(c <-chan int, q <-chan bool) {
	for {
		select {
		case msg := <-c:
			fmt.Println("received", msg)
		case <-q:
			return
		}
	}
}
