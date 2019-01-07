package arraytools

import "testing"

func TestEqualInt(t *testing.T) {
	var input1 []int
	var input2 []int
	input1 = []int{1, 2, 3}
	input2 = []int{1, 2, 3}
	if !EqualInt(input1, input2) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " should be equal")
	}

	input1 = []int{1}
	input2 = []int{1}
	if !EqualInt(input1, input2) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " should be equal")
	}

	input1 = []int{}
	input2 = []int{}
	if !EqualInt(input1, input2) {
		t.Error("Input 1 ", input1, " Input 2 ", input2, " should be equal")
	}
}

func TestReverseInt(t *testing.T) {
	var input []int
	var expected []int

	input = []int{}
	expected = []int{}
	if !EqualInt(ReverseInt(input), expected) {
		t.Error("Input: ", input, " Expected: ", expected, " Result: ", ReverseInt(input))
	}

	input = []int{1}
	expected = []int{1}
	if !EqualInt(ReverseInt(input), expected) {
		t.Error("Input: ", input, " Expected: ", expected, " Result: ", ReverseInt(input))
	}

	input = []int{1, 2, 3}
	expected = []int{3, 2, 1}
	if !EqualInt(ReverseInt(input), expected) {
		t.Error("Input: ", input, " Expected: ", expected, " Result: ", ReverseInt(input))
	}
}

func TestStartsWithInt(t *testing.T) {
	var input, startsWith []int
	var result, expected bool

	input = []int{0}
	startsWith = []int{0}
	expected = true
	result = StartsWithInt(input, startsWith)
	if result != expected {
		t.Error("Input: ", input, " StartsWith: ", startsWith, " Expected: ", expected, " Result: ", result)
	}

	input = []int{0}
	startsWith = []int{3}
	expected = false
	result = StartsWithInt(input, startsWith)
	if result != expected {
		t.Error("Input: ", input, " StartsWith: ", startsWith, " Expected: ", expected, " Result: ", result)
	}

	input = []int{}
	startsWith = []int{0}
	expected = false
	result = StartsWithInt(input, startsWith)
	if result != expected {
		t.Error("Input: ", input, " StartsWith: ", startsWith, " Expected: ", expected, " Result: ", result)
	}

	input = []int{0, 1, 2}
	startsWith = []int{0, 1}
	expected = true
	result = StartsWithInt(input, startsWith)
	if result != expected {
		t.Error("Input: ", input, " StartsWith: ", startsWith, " Expected: ", expected, " Result: ", result)
	}

	input = []int{0, 1, 3}
	startsWith = []int{0, 1, 3}
	expected = true
	result = StartsWithInt(input, startsWith)
	if result != expected {
		t.Error("Input: ", input, " StartsWith: ", startsWith, " Expected: ", expected, " Result: ", result)
	}

	input = []int{0, 1}
	startsWith = []int{0, 1, 2}
	expected = false
	result = StartsWithInt(input, startsWith)
	if result != expected {
		t.Error("Input: ", input, " StartsWith: ", startsWith, " Expected: ", expected, " Result: ", result)
	}
}

func TestLastElementInt(t *testing.T) {
	var input []int
	var expected int

	input = []int{1}
	expected = 1
	checkLastElementInt(input, expected, t)

	input = []int{1, 2}
	expected = 2
	checkLastElementInt(input, expected, t)

	input = []int{1, 2, 3, 4, 9}
	expected = 9
	checkLastElementInt(input, expected, t)
}

func checkLastElementInt(input []int, expected int, t *testing.T) {
	result := LastElementInt(input)
	if result != expected {
		t.Errorf("Input: %v, Result: %v, Expected: %v", input, result, expected)
	}
}

func TestElementInListInt(t *testing.T) {
	var input1 int
	var input2 []int
	var expected bool

	input1 = 1
	input2 = []int{1}
	expected = true
	checkElementInListInt(input1, input2, expected, t)

	input1 = 2
	input2 = []int{1}
	expected = false
	checkElementInListInt(input1, input2, expected, t)

	input1 = 1
	input2 = []int{3, 2, 1}
	expected = true
	checkElementInListInt(input1, input2, expected, t)

	input1 = 2
	input2 = []int{3, 2, 1}
	expected = true
	checkElementInListInt(input1, input2, expected, t)

	input1 = 3
	input2 = []int{3, 2, 1}
	expected = true
	checkElementInListInt(input1, input2, expected, t)

	input1 = 4
	input2 = []int{3, 2, 1}
	expected = false
	checkElementInListInt(input1, input2, expected, t)
}

func checkElementInListInt(input1 int, input2 []int, expected bool, t *testing.T) {
	result := ElementInListInt(input1, input2)
	if result != expected {
		t.Errorf("Input1: %v, Input2: %v, Result: %v, Expected: %v", input1, input2, result, expected)
	}
}
