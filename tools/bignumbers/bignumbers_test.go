package bignumbers

import (
	"testing"

	"github.com/Bisdow/projecteuler/tools/arraytools"
)

func TestParseInt(t *testing.T) {
	var input = 0
	var expected = []int{0}
	if !arraytools.EqualInt(ParseInt(input), expected) {
		t.Error("Input ", input, " Result", ParseInt(input), " Expected ", expected)
	}

	input = 11
	expected = []int{1, 1}
	if !arraytools.EqualInt(ParseInt(input), expected) {
		t.Error("Input ", input, " Result", ParseInt(input), " Expected ", expected)
	}

	input = 9876
	expected = []int{9, 8, 7, 6}
	if !arraytools.EqualInt(ParseInt(input), expected) {
		t.Error("Input ", input, " Result", ParseInt(input), " Expected ", expected)
	}
}

func TestAddInt(t *testing.T) {
	var input1 []int
	var input2 []int
	var expected []int

	input1 = []int{}
	input2 = []int{0}
	expected = []int{0}
	if !arraytools.EqualInt(AddInt(input1, input2), expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", AddInt(input1, input2), " Expected ", expected)
	}

	input1 = []int{0}
	input2 = []int{0}
	expected = []int{0}
	if !arraytools.EqualInt(AddInt(input1, input2), expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", AddInt(input1, input2), " Expected ", expected)
	}

	input1 = []int{2}
	input2 = []int{3}
	expected = []int{5}
	if !arraytools.EqualInt(AddInt(input1, input2), expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", AddInt(input1, input2), " Expected ", expected)
	}

	input1 = []int{7}
	input2 = []int{9}
	expected = []int{1, 6}
	if !arraytools.EqualInt(AddInt(input1, input2), expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", AddInt(input1, input2), " Expected ", expected)
	}

	input1 = []int{9, 9}
	input2 = []int{1, 1}
	expected = []int{1, 1, 0}
	if !arraytools.EqualInt(AddInt(input1, input2), expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", AddInt(input1, input2), " Expected ", expected)
	}

	input1 = []int{9, 9}
	input2 = []int{1}
	expected = []int{1, 0, 0}
	if !arraytools.EqualInt(AddInt(input1, input2), expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", AddInt(input1, input2), " Expected ", expected)
	}

	input1 = []int{1}
	input2 = []int{9, 9}
	expected = []int{1, 0, 0}
	if !arraytools.EqualInt(AddInt(input1, input2), expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", AddInt(input1, input2), " Expected ", expected)
	}
	input1 = []int{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}
	input2 = []int{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}
	expected = []int{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8}
	if !arraytools.EqualInt(AddInt(input1, input2), expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", AddInt(input1, input2), " Expected ", expected)
	}
}
func TestBaseOn10Potent(t *testing.T) {
	var input []int
	var zehnerPotenz int
	var expected []int
	var result []int

	// 0 * 10^0
	input = []int{0}
	zehnerPotenz = 0
	expected = []int{0}
	result = baseOn10Potent(input, zehnerPotenz)
	if !arraytools.EqualInt(result, expected) {
		t.Error("Input ", input, " Zehnerpotenz ", zehnerPotenz, " Result ", result, " Expected ", expected)
	}

	// 12 * 10^0
	input = []int{1, 2}
	zehnerPotenz = 0
	expected = []int{1, 2}
	result = baseOn10Potent(input, zehnerPotenz)
	if !arraytools.EqualInt(result, expected) {
		t.Error("Input ", input, " Zehnerpotenz ", zehnerPotenz, " Result ", result, " Expected ", expected)
	}

	// 12 * 10^1
	input = []int{1, 2}
	zehnerPotenz = 1
	expected = []int{1, 2, 0}
	result = baseOn10Potent(input, zehnerPotenz)
	if !arraytools.EqualInt(result, expected) {
		t.Error("Input ", input, " Zehnerpotenz ", zehnerPotenz, " Result ", result, " Expected ", expected)
	}

	// 12 * 10^2
	input = []int{1, 2}
	zehnerPotenz = 2
	expected = []int{1, 2, 0, 0}
	result = baseOn10Potent(input, zehnerPotenz)
	if !arraytools.EqualInt(result, expected) {
		t.Error("Input ", input, " Zehnerpotenz ", zehnerPotenz, " Result ", result, " Expected ", expected)
	}
	// 123 * 10^3
	input = []int{1, 2, 3}
	zehnerPotenz = 3
	expected = []int{1, 2, 3, 0, 0, 0}
	result = baseOn10Potent(input, zehnerPotenz)
	if !arraytools.EqualInt(result, expected) {
		t.Error("Input ", input, " Zehnerpotenz ", zehnerPotenz, " Result ", result, " Expected ", expected)
	}
}

func TestMultiplyInt(t *testing.T) {
	var input1 []int
	var input2 []int
	var expected []int
	var result []int

	input1 = []int{0}
	input2 = []int{0}
	expected = []int{0}
	result = MultiplyInt(input1, input2)
	if !arraytools.EqualInt(result, expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", result, " Expected ", expected)
	}

	input1 = []int{1}
	input2 = []int{0}
	expected = []int{0}
	result = MultiplyInt(input1, input2)
	if !arraytools.EqualInt(result, expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", result, " Expected ", expected)
	}

	input1 = []int{1}
	input2 = []int{1}
	expected = []int{1}
	result = MultiplyInt(input1, input2)
	if !arraytools.EqualInt(result, expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", result, " Expected ", expected)
	}

	input1 = []int{1, 0}
	input2 = []int{2}
	expected = []int{2, 0}
	result = MultiplyInt(input1, input2)
	if !arraytools.EqualInt(result, expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", result, " Expected ", expected)
	}

	input1 = []int{2, 5}
	input2 = []int{1, 1}
	expected = []int{2, 7, 5}
	result = MultiplyInt(input1, input2)
	if !arraytools.EqualInt(result, expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", result, " Expected ", expected)
	}

	input1 = []int{1, 1, 1}
	input2 = []int{1, 2}
	expected = []int{1, 3, 3, 2}
	result = MultiplyInt(input1, input2)
	if !arraytools.EqualInt(result, expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", result, " Expected ", expected)
	}

	input1 = []int{1, 0, 0, 0, 0, 0, 0, 0}
	input2 = []int{1, 0, 0, 0, 0, 0, 0, 0}
	expected = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	result = MultiplyInt(input1, input2)
	if !arraytools.EqualInt(result, expected) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " Result ", result, " Expected ", expected)
	}
}

func TestFactorialInt(t *testing.T) {
	var input []int
	var expected []int

	input = []int{1}
	expected = []int{1}
	checkFactorialInt(input, expected, t)

	input = []int{2}
	expected = []int{2}
	checkFactorialInt(input, expected, t)

	input = []int{3}
	expected = []int{6}
	checkFactorialInt(input, expected, t)

	input = []int{4}
	expected = []int{2, 4}
	checkFactorialInt(input, expected, t)

	input = []int{5}
	expected = []int{1, 2, 0}
	checkFactorialInt(input, expected, t)
}

func checkFactorialInt(input []int, expected []int, t *testing.T) {
	result := FactorialInt(input)
	if !arraytools.EqualInt(result, expected) {
		t.Errorf("Input: %v, Result: %v, Expected: %v", input, result, expected)
	}
}

func TestCompareInt(t *testing.T) {
	var input1 []int
	var input2 []int
	var expected int

	input1 = []int{5}
	input2 = []int{1, 2}
	expected = -1
	checkCompareInt(input1, input2, expected, t)

	input1 = []int{1, 2}
	input2 = []int{5}
	expected = 1
	checkCompareInt(input1, input2, expected, t)

	input1 = []int{5}
	input2 = []int{5}
	expected = 0
	checkCompareInt(input1, input2, expected, t)

	input1 = []int{1, 2, 3, 4, 5}
	input2 = []int{1, 2, 3, 4, 4}
	expected = 1
	checkCompareInt(input1, input2, expected, t)

	input1 = []int{5, 4, 3, 2, 1}
	input2 = []int{5, 4, 3, 2, 2}
	expected = -1
	checkCompareInt(input1, input2, expected, t)

	input1 = []int{54321}
	input2 = []int{54321}
	expected = 0
	checkCompareInt(input1, input2, expected, t)
}

func checkCompareInt(input1 []int, input2 []int, expected int, t *testing.T) {
	result := CompareInt(input1, input2)
	if result != expected {
		t.Errorf("Input1: %v, Input2: %v, Result: %v, Expected: %v", input1, input2, result, expected)
	}
}
