package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var counter int64

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			time.Sleep(time.Millisecond)
			fmt.Println("Counter value:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)

}
