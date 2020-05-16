package main

import (
	"fmt"
)

type customErr struct {
	blah string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("From error method: %v", ce.blah)
}

func foo(e error) {
	fmt.Println("Running foo:", e)
}

func main() {

	ce := customErr{
		blah: "I am a value of CE",
	}

	foo(ce)
}
