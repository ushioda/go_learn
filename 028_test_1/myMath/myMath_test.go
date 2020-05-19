package myMAth

import (
	"testing"
)

func TestSum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{
			data:   []int{21, 21},
			answer: 42,
		},
		{
			data:   []int{3, 4, 5},
			answer: 12,
		},
		{
			data:   []int{1, 1},
			answer: 2,
		},
		{
			data:   []int{-1, 0, 1},
			answer: 0,
		},
	}

	for _, test := range tests {
		v := Sum(test.data...)
		if v != test.answer {
			t.Errorf("Expected %d, got %d \n", test.answer, v)
		}
	}

}
