package quadraticprimes

import (
	"testing"

	"github.com/Bisdow/projecteuler/tools/primenumbers"
)

func TestFindConsecutivePrimes(t *testing.T) {
	checkConsecutivePrimes(0, 0, -1, t)
	checkConsecutivePrimes(1, 41, 39, t)
	checkConsecutivePrimes(-79, 1601, 79, t)
}
func checkConsecutivePrimes(a int, b int, expected int, t *testing.T) {
	result := findConsecutivePrimes(a, b)
	if result != expected {
		t.Errorf("A:: %v, B: %v, Result: %v, Expected: %v", a, b, result, expected)
	}
}

var result int

func BenchmarkFindCoefficentsProduct(b *testing.B) {
	var value int
	maxA := 1000
	maxB := 1000
	for i := 0; i < b.N; i++ {
		primenumbers.ResetPrimenumbers()
		value = findCoefficentsProduct(maxA, maxB)
	}
	result = value
}
