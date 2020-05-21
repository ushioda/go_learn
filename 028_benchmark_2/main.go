package main

import (
	"fmt"
	"go_learn/028_benchmark_2/mystr"
	"strings"
)

const s = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func main() {

	xs := strings.Split(s, " ")

	for _, w := range xs {
		fmt.Println(w)
	}

	fmt.Printf("\n %s \n", mystr.Cat(xs))
	fmt.Printf("\n %s \n\n", mystr.Join(xs))

}
