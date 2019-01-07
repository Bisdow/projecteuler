package evenfibonaccinumbers

import "testing"

func TestEvenValuedFibonacciNumbers(t *testing.T) {
	var input int
	var expected int

	//	1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
	//	0, 2, 2, 2, 10, 10, 10, 44, 44, 44, ...
	input = 1
	expected = 0
	checkEvenValuedFibonacciNumbers(input, expected, t)

	input = 2
	expected = 2
	checkEvenValuedFibonacciNumbers(input, expected, t)

	input = 8
	expected = 10
	checkEvenValuedFibonacciNumbers(input, expected, t)

	input = 13
	expected = 10
	checkEvenValuedFibonacciNumbers(input, expected, t)

	input = 34
	expected = 44
	checkEvenValuedFibonacciNumbers(input, expected, t)
}

var result int

func benchmarkEvenValuedFibonacciNumbers(limit int, b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		value = EvenValuedFibonacciNumbers(limit)
	}
	result = value
}

func BenchmarkEvenValuedFibonacciNumbers100(b *testing.B) {
	benchmarkEvenValuedFibonacciNumbers(100, b)
}
func BenchmarkEvenValuedFibonacciNumbers1000(b *testing.B) {
	benchmarkEvenValuedFibonacciNumbers(1000, b)
}
func BenchmarkEvenValuedFibonacciNumbers10000(b *testing.B) {
	benchmarkEvenValuedFibonacciNumbers(10000, b)
}
func BenchmarkEvenValuedFibonacciNumbers100000(b *testing.B) {
	benchmarkEvenValuedFibonacciNumbers(100000, b)
}
func checkEvenValuedFibonacciNumbers(input int, expected int, t *testing.T) {
	var result = EvenValuedFibonacciNumbers(input)
	if result != expected {
		t.Errorf("Input %v - Result: %v - Expected: %v", input, result, expected)
	}
}
