package main

import (
	"fmt"
	"go_learn/ex132/quote"
	"go_learn/ex132/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}