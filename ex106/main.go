package main

import "fmt"

func main() {
	ch := make(chan int)
	closed := make(chan bool)
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
		closed <- true
	}()
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case <-closed:
			return
		}
	}
}
