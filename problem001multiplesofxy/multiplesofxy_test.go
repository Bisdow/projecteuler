package problem001multiplesofxy

import "testing"

func TestSumOfMultiples(t *testing.T) {
	var result int
	var expected int

	result = SumOfMultiples(10, 2, 3)
	expected = 32 // 2 + 3 + 4 + 6 + 8 + 9 + 10 => 32
	checkInt(t, "SumOfMultiples(10,3,5)", result, expected)

	result = SumOfMultiples(10, 3, 5)
	expected = 23 // 3 + 5 + 6 + 9 => 23
	checkInt(t, "SumOfMultiples(10,3,5)", result, expected)

}

var result int

func benchmarkSumOfMultiples(limit int, firstDivider int, secondDivider int, b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		value = SumOfMultiples(limit, firstDivider, secondDivider)
	}
	result = value
}

func BenchmarkSumOfMultiples100(b *testing.B) {
	const limit int = 100
	const firstDivider int = 3
	const secondDivider int = 5
	benchmarkSumOfMultiples(limit, firstDivider, secondDivider, b)
}

func BenchmarkSumOfMultiples1000(b *testing.B) {
	const limit int = 1000
	const firstDivider int = 3
	const secondDivider int = 5
	benchmarkSumOfMultiples(limit, firstDivider, secondDivider, b)
}

func BenchmarkSumOfMultiples10000(b *testing.B) {
	const limit int = 10000
	const firstDivider int = 3
	const secondDivider int = 5
	benchmarkSumOfMultiples(limit, firstDivider, secondDivider, b)
}

func checkInt(t *testing.T, call string, result int, expected int) {
	if result != expected {
		t.Error(call, " - Expected", expected, " - Result", result)
	}
}
