package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumGoroutine())
	wg.Add(2)
	fmt.Println("This is the main goroutine")

	go two()
	go three()
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()

}

func two() {
	fmt.Println("This is the second function")
	wg.Done()
}

func three() {
	fmt.Println("This is the third function")
	wg.Done()
}
