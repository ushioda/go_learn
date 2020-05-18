package main

import (
	"fmt"
	"go_learn/ex121/dog"
)

func main() {
	hy := 5
	dy := dog.Years(hy)
	fmt.Printf("%v human years is worth %v dog years", hy, dy)
}
