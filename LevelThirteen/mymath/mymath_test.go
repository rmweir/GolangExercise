package mymath

import (
	"testing"
	"fmt"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{
			data:   []int{4, 8, 19, 21, 3},
			answer: float64(10.3333),
		},
		test{
			data:   []int{1, 2, 5, 4, 3},
			answer: float64(3),
		},
	}

	for _, x := range tests {
		if CenteredAvg(x.data) - x.answer < -1 || CenteredAvg(x.data) - x.answer > 1 {
			t.Error("Expected", x.answer, "got", CenteredAvg(x.data))
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 5, 4, 3}))
	// Output:
	// 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{10, 12,4, 3, 22, 1})
	}
}
