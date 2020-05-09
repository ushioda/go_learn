package main

import "fmt"

const n = 20

func main() {

	c := make(chan int)

	go foo(c)

	// receive

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("end of main. exiting.")
}

// send

func foo(c chan<- int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)
}
