package main

import "testing"

func Test_getSystemConfig(t *testing.T) {
	for i := 0; i < 1000; i++ {
		result := getSystemConfig()

		if result.step != 1 {
			t.Fatalf("expected step of 1, got %d\n", result.step)
		}
	}
}
