package nonabundantsums

import "testing"

func TestCheckIsAbundant(t *testing.T) {
	for i := 1; i < 12; i++ {
		checkCheckIsAbundant(i, false, t)
	}
	checkCheckIsAbundant(12, true, t)
	for i := 13; i < 18; i++ {
		checkCheckIsAbundant(i, false, t)
	}
	checkCheckIsAbundant(18, true, t)
	checkCheckIsAbundant(20, true, t)
}

func checkCheckIsAbundant(input int, expected bool, t *testing.T) {
	result := checkIsAbundant(input)
	if result != expected {
		t.Errorf("Input: %v, Result: %v, Expected: %v", input, result, expected)
	}
}

func TestCheckIfIsSumOfAbundants(t *testing.T) {
	abundantNumbers := []int{12, 18, 20, 24, 30, 36, 40, 42, 48}
	checkCheckIfIsSumOfAbundants(1, abundantNumbers, false, t)
	checkCheckIfIsSumOfAbundants(12, abundantNumbers, false, t)
	checkCheckIfIsSumOfAbundants(24, abundantNumbers, true, t)
	checkCheckIfIsSumOfAbundants(30, abundantNumbers, true, t)
	checkCheckIfIsSumOfAbundants(31, abundantNumbers, false, t)
}

func checkCheckIfIsSumOfAbundants(value int, abundantNumber []int, expected bool, t *testing.T) {
	result := checkIfIsSumOfAbundants(value, abundantNumber)
	if result != expected {
		t.Errorf("Value: %v, AbundantNumbers: %v, Result: %v, Expected: %v", value, abundantNumber, result, expected)
	}
}
