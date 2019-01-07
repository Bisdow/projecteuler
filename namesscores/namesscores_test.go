package namesscores

import "testing"

func TestValueOfName(t *testing.T) {
	var input string
	var expected int64

	input = "AAA"
	expected = 3
	checkValueOfName(input, expected, t)

	input = "COLIN"
	expected = 53
	checkValueOfName(input, expected, t)

	input = "JAN"
	expected = 25
	checkValueOfName(input, expected, t)
}

func checkValueOfName(input string, expected int64, t *testing.T) {
	result := valueOfName(input)
	if result != expected {
		t.Errorf("Input %v, Result %v, Expected: %v", input, result, expected)
	}
}
