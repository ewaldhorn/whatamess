package main

import (
	"math/rand"
	"testing"
)

func TestSumFunction(t *testing.T) {
	for i := 0; i < 10; i++ {
		num1 := rand.Int()
		num2 := rand.Int()
		total := num1 + num2

		result := CalculateSum(num1, num2)

		if total != result {
			t.Errorf("Error in test for %d and %d. Expected %d, got %d.", num1, num2, total, result)
		}
	}
}

func FuzzTestSumFunction(f *testing.F) {
	f.Fuzz(func(t *testing.T, num1 int, num2 int) {
		total := num1 + num2
		result := CalculateSum(num1, num2)
		if total != result {
			t.Errorf("Error in test for %d and %d. Expected %d, got %d.", num1, num2, total, result)
		}
	})
}
