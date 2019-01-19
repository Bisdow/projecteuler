package problem018maximumpathsumi

import (
	"testing"

	"github.com/Bisdow/projecteuler/tools/arraytools"
)

func TestCalcNextLevel(t *testing.T) {
	var input1 []int
	var input2 []int
	var expected []int

	input1 = []int{2}
	input2 = []int{4, 8}
	expected = []int{6, 10}
	checkCalcNextLevel(input1, input2, expected, t)

	input1 = []int{6, 10}
	input2 = []int{1, 5, 8}
	expected = []int{7, 15, 18}
	checkCalcNextLevel(input1, input2, expected, t)

	input1 = []int{7, 16, 18}
	input2 = []int{1, 5, 8, 12}
	expected = []int{8, 21, 26, 30}
	checkCalcNextLevel(input1, input2, expected, t)
}

func TestFindMaximumTotal(t *testing.T) {
	var input [][]int
	var expected int

	input = [][]int{
		{3},
		{7, 4}}
	expected = 10
	checkFindMaximumTotal(input, expected, t)

	input = [][]int{
		{3},
		{7, 4},
		{2, 4, 6},
		{8, 5, 9, 3}}
	expected = 23
	checkFindMaximumTotal(input, expected, t)
}

var resultCalcNextLevel []int

func BenchmarkCalcNextLevel(b *testing.B) {
	testdata1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	testdata2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	var value []int
	for i := 0; i < b.N; i++ {
		value = calcNextLevel(testdata1, testdata2)
	}
	resultCalcNextLevel = value
}

var resultFindMaximumTotal int

func BenchmarkFindMaximumTotal(b *testing.B) {
	var triangle = [][]int{
		{75},
		{95, 64},
		{17, 47, 82},
		{18, 35, 87, 10},
		{20, 4, 82, 47, 65},
		{19, 1, 23, 75, 3, 34},
		{88, 2, 77, 73, 7, 63, 67},
		{99, 65, 4, 28, 6, 16, 70, 92},
		{41, 41, 26, 56, 83, 40, 80, 70, 33},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23}}
	var value int
	for i := 0; i < b.N; i++ {
		value = findMaximumTotal(triangle)
	}
	resultFindMaximumTotal = value
}

func checkCalcNextLevel(input1 []int, input2 []int, expected []int, t *testing.T) {
	result := calcNextLevel(input1, input2)
	if !arraytools.EqualInt(result, expected) {
		t.Errorf("Input1: %v\nInput2: %v\nResult: %v\nExpected: %v\n", input1, input2, result, expected)
	}
}

func checkFindMaximumTotal(input [][]int, expected int, t *testing.T) {
	result := findMaximumTotal(input)
	if result != expected {
		t.Errorf("Input: %v\nResult: %v\nExpected: %v\n", input, result, expected)
	}
}
