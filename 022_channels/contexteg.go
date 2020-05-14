package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:\t\t", ctx.Err())
	fmt.Println("num go routines 1:\t\t", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:\t\t", ctx.Err())
	fmt.Println("num go routines 2:\t\t", runtime.NumGoroutine())

	fmt.Println("cancelling context")
	cancel()
	fmt.Println("context cancelled")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:\t\t", ctx.Err())
	fmt.Println("num go routines 3:\t\t", runtime.NumGoroutine())
}
