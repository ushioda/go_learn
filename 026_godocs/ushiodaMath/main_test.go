package ushiodaMath_test

import (
	"go_learn/026_godocs/ushiodaMath"
	"testing"
)

func TestSum(t *testing.T) {
	p := 4.5
	q := 3.2
	v := ushiodaMath.Sum(p, q)
	if v != 7.7 {
		t.Error("Expected 1.5, got ", v)
	}
}
