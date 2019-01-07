package sumsquaredifference

import "testing"

func TestCalculateDelta(t *testing.T) {
	var input float64
	var expected float64

	input = 1
	expected = 0
	checkCalculateDelta(input, expected, t)

	input = 2
	expected = 4 // 1 +2 = 3^2 = 9   --- 1^2 + 2^2 = 5  => 4
	checkCalculateDelta(input, expected, t)

	input = 3
	expected = 22 // 1 + 2 + 3 = 6^2 = 36   --- 1^2 + 2^2 + 3^2 = 14  => 22
	checkCalculateDelta(input, expected, t)

	input = 10
	expected = 2640
	checkCalculateDelta(input, expected, t)
}

var result float64

func benchmarkCalculateDelta(input float64, b *testing.B) {
	var value float64
	for i := 0; i < b.N; i++ {
		value = CalculateDelta(input)
	}
	result = value
}

func BenchmarkCalculateDelta5(b *testing.B) {
	benchmarkCalculateDelta(5, b)
}
func BenchmarkCalculateDelta10(b *testing.B) {
	benchmarkCalculateDelta(10, b)
}
func BenchmarkCalculateDelta30(b *testing.B) {
	benchmarkCalculateDelta(30, b)
}
func BenchmarkCalculateDelta50(b *testing.B) {
	benchmarkCalculateDelta(50, b)
}
func BenchmarkCalculateDelta100(b *testing.B) {
	benchmarkCalculateDelta(100, b)
}
func BenchmarkCalculateDelta200(b *testing.B) {
	benchmarkCalculateDelta(200, b)
}

func checkCalculateDelta(input float64, expected float64, t *testing.T) {
	result := CalculateDelta(input)
	if result != expected {
		t.Errorf("Input: %v, Result: %v, Expected: %v", input, result, expected)
	}
}
