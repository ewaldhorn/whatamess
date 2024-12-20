package basic

import "testing"

func BenchmarkSliceAppend(b *testing.B) {
	a := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		a = append(a, i)
	}
}

func BenchmarkSliceSet(b *testing.B) {
	a := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		a[i] = i
	}
}
