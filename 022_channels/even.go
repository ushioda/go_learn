package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(odd, even, quit)

	// receive
	receive(odd, even, quit)

	fmt.Println("end of main")
}

func send(o, e, q chan<- int) {
	for i := 0; i < 20; i++ {
		switch {
		case i%2 == 0:
			e <- i
		case i%2 != 0:
			o <- i
		}
	}
	//close(e)
	//close(o)
	q <- 0
}

func receive(o, e, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from even channel", v)
		case v := <-o:
			fmt.Println("from odd channel", v)
		case v := <-q:
			fmt.Println("from quit channel", v)
			return
		}
	}
}
