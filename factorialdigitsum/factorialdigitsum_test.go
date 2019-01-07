package factorialdigitsum

import (
	"testing"
)

func TestFactorialdigitsum(t *testing.T) {
	var input int
	var expected int

	input = 1
	expected = 1
	checkFactorialdigitsum(input, expected, t)

	input = 4 // => 24
	expected = 6
	checkFactorialdigitsum(input, expected, t)

	input = 5 // => 120
	expected = 3
	checkFactorialdigitsum(input, expected, t)

	input = 10 // => 3628800
	expected = 27
	checkFactorialdigitsum(input, expected, t)
}

func checkFactorialdigitsum(input int, expected int, t *testing.T) {
	result := Factorialdigitsum(input)
	if result != expected {
		t.Errorf("Input: %v, Result: %v, Expected: %v", input, result, expected)
	}
}

var result int

func benchmarkFactorialdigitsum(input int, b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		value = Factorialdigitsum(input)
	}
	result = value
}

func BenchmarkFactorialdigitsum10(b *testing.B) {
	benchmarkFactorialdigitsum(10, b)
}
func BenchmarkFactorialdigitsum50(b *testing.B) {
	benchmarkFactorialdigitsum(50, b)
}
func BenchmarkFactorialdigitsum100(b *testing.B) {
	benchmarkFactorialdigitsum(100, b)
}
