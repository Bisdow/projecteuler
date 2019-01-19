package problem017numberlettercounts

import "testing"

func TestGetWordLengthAllNumbers(t *testing.T) {
	initHashMap()
	var input, expected, result int

	input = 5
	expected = 19
	result = getWordLengthAllNumbers(input)
	if expected != result {
		t.Errorf("Input %d - Result %d - Expected %d", input, result, expected)
	}

	input = 9
	expected = 36
	result = getWordLengthAllNumbers(input)
	if expected != result {
		t.Errorf("Input %d - Result %d - Expected %d", input, result, expected)
	}

	input = 19
	// one + two + three + four + five + six + seven + eight
	// nine + ten
	// => 3+3+5+4+4+3+5+5+4+3 = 39
	expected = 106
	result = getWordLengthAllNumbers(input)
	if expected != result {
		t.Errorf("Input %d - Result %d - Expected %d", input, result, expected)
	}

	input = 20
	// 1 -> 10 = 39
	// + eleven + twelve + thirteen + fourteen + fifteen +
	// sixteen + seventeen + eighteen + nineteen + twenty
	// 6+6+8+8+7+7+9+8+8+6 = 73
	// 39 + 73 = 112
	expected = 112
	result = getWordLengthAllNumbers(input)
	if expected != result {
		t.Errorf("Input %d - Result %d - Expected %d", input, result, expected)
	}

	input = 25
	// 1 -> 20 = 112
	// twenty-one + twenty-two + twenty-three + twenty-four + twenty-five
	// 9 + 9 + 11 + 10 + 10 = 49
	// 112 + 49 = 161
	expected = 161
	result = getWordLengthAllNumbers(input)
	if expected != result {
		t.Errorf("Input %d - Result %d - Expected %d", input, result, expected)
	}

	input = 99
	expected = 854
	result = getWordLengthAllNumbers(input)
	if expected != result {
		t.Errorf("Input %d - Result %d - Expected %d", input, result, expected)
	}
}

func TestGetWordLenghtOfNumber(t *testing.T) {
	var input, expected, result int
	initHashMap()

	input = 5
	expected = 4 // five
	result = getWordLengthOfNumber(input)
	if expected != result {
		t.Errorf("Input %d - Result %d - Expected %d", input, result, expected)
	}

	input = 23
	expected = 11 // twenty three
	result = getWordLengthOfNumber(input)
	if expected != result {
		t.Errorf("Input %d - Result %d - Expected %d", input, result, expected)
	}

	input = 115 // one hundred and fifteen
	expected = 20
	result = getWordLengthOfNumber(input)
	if expected != result {
		t.Errorf("Input %d - Result %d - Expected %d", input, result, expected)
	}

	input = 342
	expected = 23 // three hundred and forty-two
	result = getWordLengthOfNumber(input)
	if expected != result {
		t.Errorf("Input %d - Result %d - Expected %d", input, result, expected)
	}

	input = 400
	expected = 11 // four hundred
	result = getWordLengthOfNumber(input)
	if expected != result {
		t.Errorf("Input %d - Result %d - Expected %d", input, result, expected)
	}
}

var result int

func benchmarkGetWordLengthAllNumbers(target int, b *testing.B) {
	var val int
	for i := 0; i < b.N; i++ {
		val = getWordLengthAllNumbers(target)
	}
	result = val
}

func BenchmarkGetWordLengthAllNumbers10(b *testing.B)    { benchmarkGetWordLengthAllNumbers(10, b) }
func BenchmarkGetWordLengthAllNumbers100(b *testing.B)   { benchmarkGetWordLengthAllNumbers(100, b) }
func BenchmarkGetWordLengthAllNumbers1000(b *testing.B)  { benchmarkGetWordLengthAllNumbers(1000, b) }
func BenchmarkGetWordLengthAllNumbers10000(b *testing.B) { benchmarkGetWordLengthAllNumbers(10000, b) }
