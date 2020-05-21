package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	ms := "I am invincible."
	n := Count(ms)
	if n != 3 {
		t.Error("Expected 3, got ", n)
	}
}

func ExampleCount() {
	fmt.Println(Count("I am invincible."))
	//Output:
	//3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("I am invincible.")
	}
}

func TestUseCount(t *testing.T) {
	ms := "I am invincible."
	ans := map[string]int{"I": 1, "am": 1, "invincible.": 1}
	mm := UseCount(ms)
	pass := true
	for i, v := range mm {
		if v != ans[i] {
			pass = false
		}
	}
	if pass == false {
		t.Error("Did not count correctly")
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("I am invincible.")
	}
}
