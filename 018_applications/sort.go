package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

// ByAge implements sort.Interface for []Person based on
// the Age field.

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByFirst []Person

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func (a ByFirst) Len() int           { return len(a) }
func (a ByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirst) Less(i, j int) bool { return a[i].First < a[j].First }

//func (a ByFirst) Less(i, j int) bool {
//	for k := 0; k < min(len(a[i].First), len(a[j].First)); k++ {
//		if byte(a[i].First[k]) < byte(a[j].First[k]) {
//			return true
//		} else if byte(a[i].First[k]) > byte(a[j].First[k]) {
//			return false
//		}
//	}
//	if len(a[i].First) < len(a[j].First) {
//		return true
//	} else {
//		return false
//	}
//}

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Money", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByFirst(people))
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].First < people[j].First
	})
	fmt.Println(people)
}
