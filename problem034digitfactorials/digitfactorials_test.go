package problem034digitfactorials

import "testing"

func TestCheckNumber(t *testing.T) {
	factorials0To9 = calcFactorials()
	checkCheckNumber(144, false, t)
	checkCheckNumber(145, true, t)
	checkCheckNumber(40585, true, t)
}

func checkCheckNumber(input int, expected bool, t *testing.T) {
	result := checkNumber(input)
	if result != expected {
		t.Errorf("Input %v, Result %v, Expected %v", input, result, expected)
	}
}
