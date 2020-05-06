package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	p1 := "password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(p1), bcrypt.MinCost)
	if err != nil {
		println(err)
	}
	fmt.Println(p1)
	fmt.Println(bs)

	p2 := "password123"
	err = bcrypt.CompareHashAndPassword(bs, []byte(p2))
	if err != nil {
		fmt.Println("Pass word does not match")
	}
	fmt.Println(err)
}
