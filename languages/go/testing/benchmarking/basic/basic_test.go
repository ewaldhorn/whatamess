package basic

import "testing"

func Test_ConcatenateBuffer(t *testing.T) {
	result := ConcatenateBuffer("left", "right")

	if result != "leftright" {
		t.Errorf("expected 'leftright', got '%s'", result)
	}
}

func Test_ConcatenateJoin(t *testing.T) {
	result := ConcatenateJoin("left", "right")

	if result != "leftright" {
		t.Errorf("expected 'leftright', got '%s'", result)
	}
}

var result string

func Benchmark_ConcatenateBuffer(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ConcatenateBuffer("left", "right")
	}
	result = s
}

func Benchmark_ConcatenateJoin(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ConcatenateJoin("left", "right")
	}
	result = s
}
