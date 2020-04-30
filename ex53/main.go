package main

import "fmt"

func main() {

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}

	t1 := truck{
		vehicle:   vehicle{2, "yellow"},
		fourWheel: true,
	}
	s1 := sedan{
		vehicle: vehicle{4, "white"},
		luxury:  false,
	}

	fmt.Println(t1)
	fmt.Println(t1.doors)
	fmt.Println(t1.color)
	fmt.Println(t1.fourWheel)
	fmt.Println(s1)
	fmt.Println(s1.doors)
	fmt.Println(s1.color)
	fmt.Println(s1.luxury)
}