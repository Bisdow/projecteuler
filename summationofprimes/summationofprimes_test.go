package summationofprimes

import (
	"testing"

	"github.com/Bisdow/projecteuler/tools/primenumbers"
)

func TestSummationOfPrimes(t *testing.T) {
	var input int
	var expected int

	input = 3
	expected = 2 // 2
	checkSummationOfPrimes(input, expected, t)

	input = 3
	expected = 2 // 2
	checkSummationOfPrimes(input, expected, t)

	input = 6
	expected = 10 // 2 + 3 + 5 = 10
	checkSummationOfPrimes(input, expected, t)

	input = 11
	expected = 17 // 2 + 3 + 5 + 7= 17
	checkSummationOfPrimes(input, expected, t)

	input = 2000000
	expected = 142913828922
	checkSummationOfPrimes(input, expected, t)
}

var result int

func benchmarkSummationOfPrimes(input int, b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		primenumbers.ResetPrimenumbers()
		value = SummationOfPrimes(input)
	}
	result = value
}

func BenchmarkSummationOfPrimes10(b *testing.B) {
	benchmarkSummationOfPrimes(10, b)
}
func BenchmarkSummationOfPrimes100(b *testing.B) {
	benchmarkSummationOfPrimes(100, b)
}

func BenchmarkSummationOfPrimes1000(b *testing.B) {
	benchmarkSummationOfPrimes(1000, b)
}

func BenchmarkSummationOfPrimes10000(b *testing.B) {
	benchmarkSummationOfPrimes(10000, b)
}

func checkSummationOfPrimes(input int, expected int, t *testing.T) {
	primenumbers.ResetPrimenumbers()
	result := SummationOfPrimes(input)
	if result != expected {
		t.Errorf("Input: %v, Result: %v, Expected: %v", input, result, expected)
	}
}
