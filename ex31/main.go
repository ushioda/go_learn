package main

import (
	"fmt"
)

func main() {
	bd := 1986
	for {
		fmt.Println(bd)
		if bd >= 2020 {
			break
		}
		bd ++
		}
	}
