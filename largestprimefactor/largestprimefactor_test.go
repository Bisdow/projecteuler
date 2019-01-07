package largestprimefactor

import (
	"testing"

	"github.com/Bisdow/projecteuler/tools/primenumbers"
)

func TestFindLargestPrimeFactor(t *testing.T) {
	var input int
	var expected int

	input = 6
	expected = 3
	checkFindLargestPrimeFactor(input, expected, t)

	input = 10
	expected = 5
	checkFindLargestPrimeFactor(input, expected, t)

	input = 12
	expected = 3
	checkFindLargestPrimeFactor(input, expected, t)

	input = 14
	expected = 7
	checkFindLargestPrimeFactor(input, expected, t)

	input = 21
	expected = 7
	checkFindLargestPrimeFactor(input, expected, t)
}

var result int

func benchmarkFindLargestPrimeFactor(input int, b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		value = FindLargestPrimeFactor(input)
	}
	result = value
}

func BenchmarkFindLargestPrimeFactor10(b *testing.B) {
	benchmarkFindLargestPrimeFactor(10, b)
}
func BenchmarkFindLargestPrimeFactor100(b *testing.B) {
	benchmarkFindLargestPrimeFactor(100, b)
}
func BenchmarkFindLargestPrimeFactor1000(b *testing.B) {
	benchmarkFindLargestPrimeFactor(1000, b)
}
func BenchmarkFindLargestPrimeFactor10000(b *testing.B) {
	benchmarkFindLargestPrimeFactor(10000, b)
}

func checkFindLargestPrimeFactor(input int, expected int, t *testing.T) {
	primenumbers.ResetPrimenumbers()
	var result = FindLargestPrimeFactor(input)
	if result != expected {
		t.Errorf("Input: %v - Result: %v - Expected: %v", input, result, expected)
	}
}
