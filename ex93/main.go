package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	counter := 0
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			time.Sleep(time.Millisecond)
			counter = v
			fmt.Println("Counter value:", counter)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)

}
