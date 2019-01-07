package mathtools

import (
	"testing"

	"github.com/Bisdow/projecteuler/tools/arraytools"
)

func TestFindDivisorsOf(t *testing.T) {
	var input int
	var expected []int

	input = 0
	expected = []int{0}
	checkFindDivisorOf(input, expected, t)

	input = 1
	expected = []int{1}
	checkFindDivisorOf(input, expected, t)

	input = 2
	expected = []int{1, 2}
	checkFindDivisorOf(input, expected, t)

	input = 4
	expected = []int{1, 2, 4}
	checkFindDivisorOf(input, expected, t)

	input = 9
	expected = []int{1, 3, 9}
	checkFindDivisorOf(input, expected, t)

	input = 24
	expected = []int{1, 2, 3, 4, 6, 8, 12, 24}
	checkFindDivisorOf(input, expected, t)
}

func checkFindDivisorOf(input int, expected []int, t *testing.T) {
	result := FindDivisorsOf(input)
	if !arraytools.EqualInt(result, expected) {
		t.Errorf("Input: %v, Result: %v, Expected: %v", input, result, expected)
	}
}

func TestFactorial(t *testing.T) {
	checkFactorial(0, 1, t)
	checkFactorial(1, 1, t)
	checkFactorial(2, 2, t)
	checkFactorial(3, 6, t)
	checkFactorial(4, 24, t)
	checkFactorial(5, 120, t)
}

func checkFactorial(input uint64, expected uint64, t *testing.T) {
	result := Factorial(input)
	if result != expected {
		t.Errorf("Input %v, Result %v, Expected %v", input, result, expected)
	}
}

// BENCHMARKS
var resultFindDivisorOf []int

func benchmarkFindDivisorOf(input int, b *testing.B) {
	var value []int
	for i := 0; i < b.N; i++ {
		value = FindDivisorsOf(input)
	}
	resultFindDivisorOf = value
}
func BenchmarkFindDivisorOf100(b *testing.B) {
	benchmarkFindDivisorOf(100, b)
}
func BenchmarkFindDivisorOf1000(b *testing.B) {
	benchmarkFindDivisorOf(1000, b)
}
func BenchmarkFindDivisorOf10000(b *testing.B) {
	benchmarkFindDivisorOf(10000, b)
}

var resultFactorial uint64

func benchmarkFactorial(input uint64, b *testing.B) {
	var value uint64
	for i := 0; i < b.N; i++ {
		value = Factorial(input)
	}
	resultFactorial = value
}

func BenchmarkFactorial10(b *testing.B) {
	benchmarkFactorial(10, b)
}
func BenchmarkFactorial100(b *testing.B) {
	benchmarkFactorial(100, b)
}
func BenchmarkFactorial1000(b *testing.B) {
	benchmarkFactorial(1000, b)
}
func BenchmarkFactorial10000(b *testing.B) {
	benchmarkFactorial(10000, b)
}
