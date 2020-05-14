package main

import (
	"fmt"
)

func main() {

	const numWorkers = 10
	const workLoad = 10
	ch := make(chan int)

	for n := 0; n < numWorkers; n++ {
		go func(vn int) {
			for i := 0; i < workLoad; i++ {
				ch <- vn*10 + i
			}
		}(n)
	}

	for v := 0; v < numWorkers*workLoad; v++ {
		fmt.Println(<-ch)
	}
	fmt.Println("exit")
}
