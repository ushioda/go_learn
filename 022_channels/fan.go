package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go sen(even, odd)
	go rec(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
	fmt.Println("exit main")
}

func sen(e, o chan<- int) {
	for i := 0; i < 20; i++ {
		switch {
		case i%2 == 0:
			e <- i
		case i%2 != 0:
			o <- i
		}
	}
	close(e)
	close(o)
}

func rec(e, o <-chan int, fan chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fan <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range o {
			fan <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fan)
}
