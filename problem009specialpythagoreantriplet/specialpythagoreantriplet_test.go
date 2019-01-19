package problem009specialpythagoreantriplet

import "testing"

func TestSpecialPythagoreanTriplet(t *testing.T) {
	var input int
	var expected int

	input = 12 // 3 + 4 + 5 ==> 3^2 + 4^2 = 5^2 => 3*4*5=60
	expected = 60
	checkSpecialPythagoreanTriplet(input, expected, t)

	input = 1000
	expected = 31875000
	checkSpecialPythagoreanTriplet(input, expected, t)
}

var result int

func benchmarkSpecialPythagoreanTriplet(input int, b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		value = SpecialPythagoreanTriplet(input)
	}
	result = value
}

func BenchmarkSpecialPythagoreanTriplet10(b *testing.B) {
	benchmarkSpecialPythagoreanTriplet(10, b)
}
func BenchmarkSpecialPythagoreanTriplet100(b *testing.B) {
	benchmarkSpecialPythagoreanTriplet(100, b)
}
func BenchmarkSpecialPythagoreanTriplet1000(b *testing.B) {
	benchmarkSpecialPythagoreanTriplet(1000, b)
}

func checkSpecialPythagoreanTriplet(input int, expected int, t *testing.T) {
	result := SpecialPythagoreanTriplet(input)
	if result != expected {
		t.Errorf("Input: %v, Result: %v, Expected: %v", input, result, expected)
	}
}
