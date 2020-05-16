package main

import (
	"fmt"
	"log"
)

type mathError struct {
	lat  string
	long string
	err  error
}

func (m mathError) Error() string {
	return fmt.Sprintf("math error occured: %v %v %v", m.lat, m.long, m.err)
}

func main() {
	_, err := squareRoot(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func squareRoot(f float64) (float64, error) {
	if f < 0 {
		me := fmt.Errorf("sqaure root of negative: %v", f)
		return 0, mathError{"50.2289 N", "99.4656 W", me}
	}
	return 42, nil
}
