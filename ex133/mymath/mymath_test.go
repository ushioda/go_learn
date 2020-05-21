package mymath

import (
	"fmt"
	"testing"
)

func ExampleCenteredAvg() {
	a := []int{1, 4, 6, 8, 100}
	fmt.Println(CenteredAvg(a))
	//Output:
	//6
}

func BenchmarkCenteredAvg(b *testing.B) {
	a := []int{1, 4, 6, 8, 100}
	for i := 0; i < b.N; i++ {
		CenteredAvg(a)
	}
}

func TestCenteredAvg(t *testing.T) {

	type test struct {
		data []int
		ans  float64
	}

	tests := []test{
		{[]int{1, 4, 6, 8, 100}, 6},
		{[]int{0, 8, 10, 1000}, 9},
		{[]int{9000, 4, 10, 8, 6, 12}, 9},
	}

	for _, test := range tests {
		v := CenteredAvg(test.data)
		if v != test.ans {
			t.Errorf("Expected %v, got %v", test.ans, v)
		}
	}
}
