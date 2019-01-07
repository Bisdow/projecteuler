package smallestmultiple

import "testing"

/*
BenchmarkFindSmallestMultiple5-8    	10000000	       190 ns/op	      56 B/op	       3 allocs/op
BenchmarkFindSmallestMultiple10-8   	 1000000	      1051 ns/op	     120 B/op	       4 allocs/op
BenchmarkFindSmallestMultiple15-8   	   20000	     65605 ns/op	     248 B/op	       5 allocs/op
BenchmarkFindSmallestMultiple20-8   	     100	  10498264 ns/op	     248 B/op	       5 allocs/op
*/

func TestFindSmallestMultiple(t *testing.T) {
	var maxDivider int
	var expected int

	maxDivider = 1
	expected = 1
	checkFindSmallestMultiple(maxDivider, expected, t)

	maxDivider = 2
	expected = 2
	checkFindSmallestMultiple(maxDivider, expected, t)

	maxDivider = 3
	expected = 6
	checkFindSmallestMultiple(maxDivider, expected, t)

	maxDivider = 4
	expected = 12
	checkFindSmallestMultiple(maxDivider, expected, t)

	maxDivider = 10
	expected = 2520
	checkFindSmallestMultiple(maxDivider, expected, t)

	maxDivider = 20
	expected = 232792560
	checkFindSmallestMultiple(maxDivider, expected, t)
}

var result int

func benchmarkFindSmallestMultiple(maxDivider int, b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		value = FindSmallestMultiple(maxDivider)
	}
	result = value
}

func BenchmarkFindSmallestMultiple5(b *testing.B) {
	benchmarkFindSmallestMultiple(5, b)
}
func BenchmarkFindSmallestMultiple10(b *testing.B) {
	benchmarkFindSmallestMultiple(10, b)
}

func BenchmarkFindSmallestMultiple15(b *testing.B) {
	benchmarkFindSmallestMultiple(15, b)
}
func BenchmarkFindSmallestMultiple20(b *testing.B) {
	benchmarkFindSmallestMultiple(20, b)
}

func checkFindSmallestMultiple(maxDivider int, expected int, t *testing.T) {
	result := FindSmallestMultiple(maxDivider)
	if result != expected {
		t.Errorf("MaxDivider: %v, Result: %v, Expected: %v", maxDivider, result, expected)
	}
}
