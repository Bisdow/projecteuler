package problem025xdigitfibonaccinumber

import "testing"

func TestXdigitfibonaccinumber(t *testing.T) {
	check(2, 7, t)
	//check(3, 12, t)
}

func check(limitDigts int, expected uint64, t *testing.T) {
	result := xteFibonacciNumber(limitDigts)
	if result != expected {
		t.Errorf("LimitDigits: %v, Result: %v, Expected: %v", limitDigts, result, expected)
	}
}
