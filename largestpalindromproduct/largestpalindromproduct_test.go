package largestpalindromproduct

import "testing"

func TestFindHighestPalindrom(t *testing.T) {
	var input1, input2 int
	var expected1, expected2 int

	input1 = 10
	input2 = 100
	expected1 = 91
	expected2 = 99
	checkFindHighestPalindrom(input1, input2, expected1, expected2, t)

	input1 = 100
	input2 = 1000
	expected1 = 913
	expected2 = 993
	checkFindHighestPalindrom(input1, input2, expected1, expected2, t)
}

var result1, result2 int

func benchmarkFindHighestPalindrom(lowLimit int, maxLimit int, b *testing.B) {
	var value1, value2 int
	for i := 0; i < b.N; i++ {
		value1, value2 = FindHighestPalindrom(lowLimit, maxLimit)
	}
	result1, result2 = value1, value2
}
func BenchmarkFindHighestPalindrom100(b *testing.B) {
	benchmarkFindHighestPalindrom(1, 100, b)
}
func BenchmarkFindHighestPalindrom1000(b *testing.B) {
	benchmarkFindHighestPalindrom(100, 1000, b)
}
func BenchmarkFindHighestPalindrom10000(b *testing.B) {
	benchmarkFindHighestPalindrom(1, 10000, b)
}

func checkFindHighestPalindrom(lowLimit int, maxLimit int, expected1 int, expected2 int, t *testing.T) {
	resultA, resultB := FindHighestPalindrom(lowLimit, maxLimit)

	if !(resultA == expected1 && resultB == expected2) {
		t.Errorf("LowLimit %v, MaxLimit %v, Expected 1: %v, Result 1 %v, Expected 2: %v, Result 2: %v ", lowLimit, maxLimit, expected1, resultA, expected2, resultB)
	}
}
