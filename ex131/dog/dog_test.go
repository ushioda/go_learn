package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	dy := 4
	hy := Years(dy)
	if hy != 28 {
		t.Error("Expected 28, got ", hy)
	}
}

func ExampleYears() {
	dy := 4
	fmt.Println(Years(dy))
	//	Output:
	//	28
}

func BenchmarkYears(b *testing.B) {
	dy := 4
	for i := 0; i < b.N; i++ {
		Years(dy)
	}
}

func TestYearsTwo(t *testing.T) {
	dy := 4
	hy := YearsTwo(dy)
	if hy != 28 {
		t.Error("Expected 28, got ", hy)
	}
}

func ExampleYearsTwo() {
	dy := 4
	fmt.Println(YearsTwo(dy))
	//Output:
	//28
}

func BenchmarkYearsTwo(b *testing.B) {
	dy := 4
	for i := 0; i < b.N; i++ {
		YearsTwo(dy)
	}
}
