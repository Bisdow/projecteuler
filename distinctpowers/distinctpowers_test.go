package distinctpowers

import "testing"

func TestDistinctPowers(t *testing.T) {
	check(5, 5, 15, t)
}

func check(a int, b int, expected int, t *testing.T) {
	result := distinctPowers(a, b)
	if result != expected {
		t.Errorf("A: %v, B: %v, Result: %v, Expected: %v", a, b, result, expected)
	}
}
