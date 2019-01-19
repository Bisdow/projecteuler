package problem028numberspiraldiagonals

import "testing"

func TestCalcDiagonals(t *testing.T) {
	check(3, 25, t)
	check(5, 101, t)
}

func check(targetSize int, expected uint64, t *testing.T) {
	result := calcDiagonals(targetSize)
	if result != expected {
		t.Errorf("TargetSize: %v, Result: %v, Expected: %v", targetSize, result, expected)
	}
}
