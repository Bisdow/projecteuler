package problem026reciprocalcycles

import "testing"

func TestSpecific(t *testing.T) {
	input := 16
	expected := 0
	checkfindLengthOfRecurringCycle(input, expected, t)

}
func TestFindLengthOfRecurringCycle(t *testing.T) {
	var input int
	var expected int

	input = 2
	expected = 0
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 3
	expected = 1
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 4
	expected = 0
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 5
	expected = 0
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 6
	expected = 1
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 7
	expected = 6
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 8
	expected = 0
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 9
	expected = 1
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 10
	expected = 0
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 11
	expected = 2
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 12
	expected = 1
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 13
	expected = 6
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 14
	expected = 6
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 15
	expected = 1
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 16
	expected = 0
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 17
	expected = 16
	checkfindLengthOfRecurringCycle(input, expected, t)
	input = 18
	expected = 1
	checkfindLengthOfRecurringCycle(input, expected, t)
}

func checkfindLengthOfRecurringCycle(input int, expected int, t *testing.T) {
	result := findLengthOfRecurringCycle(input)
	if result != expected {
		t.Errorf("Input: %v, Result: %v, Expected: %v", input, result, expected)
	}
}
