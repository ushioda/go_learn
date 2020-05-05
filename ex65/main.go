package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type shape interface {
	area() float64
}

func info(sh shape) float64 {
	return sh.area()
}

func main() {
	s1 := square{length: 5.5}
	c1 := circle{radius: 4.2}
	fmt.Println(info(s1))
	fmt.Println(info(c1))
}
