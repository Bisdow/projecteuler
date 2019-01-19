package problem007xteprime

import (
	"testing"

	"github.com/Bisdow/projecteuler/tools/primenumbers"
)

func TestFindPrimeNumber(t *testing.T) {
	var input int
	var expected int

	input = 1
	expected = 2
	checkFindPrimeNumber(input, expected, t)

	input = 2
	expected = 3
	checkFindPrimeNumber(input, expected, t)

	input = 3
	expected = 5
	checkFindPrimeNumber(input, expected, t)

	input = 4
	expected = 7
	checkFindPrimeNumber(input, expected, t)

	input = 5
	expected = 11
	checkFindPrimeNumber(input, expected, t)

	input = 6
	expected = 13
	checkFindPrimeNumber(input, expected, t)

	input = 7
	expected = 17
	checkFindPrimeNumber(input, expected, t)
}

var result int

func benchmarkFindPrimeNumber(input int, b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		primenumbers.ResetPrimenumbers()
		value = FindPrimeNumber(input)
	}
	result = value
}

func BenchmarkFindPrimeNumber10(b *testing.B) {
	benchmarkFindPrimeNumber(10, b)
}
func BenchmarkFindPrimeNumber100(b *testing.B) {
	benchmarkFindPrimeNumber(100, b)
}
func BenchmarkFindPrimeNumber1000(b *testing.B) {
	benchmarkFindPrimeNumber(1000, b)
}
func BenchmarkFindPrimeNumber10000(b *testing.B) {
	benchmarkFindPrimeNumber(10000, b)
}

func checkFindPrimeNumber(input int, expected int, t *testing.T) {
	result := FindPrimeNumber(input)
	if result != expected {
		t.Errorf("Input: %v, Result: %v, Expected: %v", input, result, expected)
	}
}
