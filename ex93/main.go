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

	var mu sync.Mutex

	counter := 0
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			time.Sleep(time.Millisecond)
			counter = v
			fmt.Println("Counter value:", counter)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)

}
