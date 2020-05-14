package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
}
