package myPacks_test

import (
	"go_learn/ex115/myPacks"
	"testing"
)

func TestSum(t *testing.T) {
	p := 4.5
	q := 3.2
	v := myPacks.Sum(p, q)
	if v != 7.7 {
		t.Error("Expected 1.5, got ", v)
	}
}
