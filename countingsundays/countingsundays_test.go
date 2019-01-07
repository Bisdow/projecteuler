package countingsundays

import (
	"fmt"
	"testing"
)

func TestAddDays(t *testing.T) {
	var input Date
	var expected Date

	input = Date{1, 1, 2010}
	expected = Date{8, 1, 2010}
	checkAddDays(input, 7, expected, t)

	input = Date{25, 1, 2001}
	expected = Date{1, 2, 2001}
	checkAddDays(input, 7, expected, t)

	input = Date{6, 11, 2010}
	expected = Date{14, 11, 2010}
	checkAddDays(input, 7, expected, t)
}

func TestIsLeapYear(t *testing.T) {
	var input int
	var expected bool

	input = 2000
	expected = true
	checkIsLeapYear(input, expected, t)

	input = 2001
	expected = false
	checkIsLeapYear(input, expected, t)

	input = 2002
	expected = false
	checkIsLeapYear(input, expected, t)

	input = 2004
	expected = true
	checkIsLeapYear(input, expected, t)

	input = 2100
	expected = false
	checkIsLeapYear(input, expected, t)
}

func checkIsLeapYear(input int, expected bool, t *testing.T) {
	result := isLeapYear(input)
	if result != expected {
		fmt.Println("Input: ", input, "Result: ", result, "Expected: ", expected)
	}
}

func checkAddDays(input1 Date, input2 int, expected Date, t *testing.T) {
	result := addDays(input1, input2)
	if result.date != expected.date && result.month != expected.month && result.year != expected.year {
		t.Error("Input 1: ", input1, "Add Days:", input2, " Result: ", result, " Expected: ", expected)
	}
}
