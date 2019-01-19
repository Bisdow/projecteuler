package bignumbers

import (
	"strconv"

	"github.com/Bisdow/projecteuler/tools/arraytools"
)

// ParseInt - Parses an int into an OneDigit Array
func ParseInt(value int) []int {
	str := strconv.Itoa(value)
	var result []int
	for i := 0; i < len(str); i++ {
		value, err := strconv.Atoi(str[i : i+1])
		if err != nil {
			println("FEHLER")
		}
		result = append(result, value)
	}
	return result
}

// AddInt - Addiert zwei Zahlen miteinander
func AddInt(val1 []int, val2 []int) []int {
	maxLen := getLongestLength(val1, val2)
	length1 := len(val1)
	length2 := len(val2)
	reverseVal1 := arraytools.ReverseInt(val1)
	reverseVal2 := arraytools.ReverseInt(val2)

	var reverseResult []int
	var uebertrag = 0
	for i := 0; i < maxLen; i++ {
		var value1 = 0
		if i < length1 {
			value1 = reverseVal1[i]
		}
		var value2 = 0
		if i < length2 {
			value2 = reverseVal2[i]
		}
		value := value1 + value2 + uebertrag
		reverseResult = append(reverseResult, value%10)
		uebertrag = value / 10
	}
	for uebertrag > 0 {
		reverseResult = append(reverseResult, uebertrag%10)
		uebertrag /= 10
	}
	return arraytools.ReverseInt(reverseResult)
}

// MultiplyInt - Multipliziert zwei Zahlen miteinander
func MultiplyInt(base []int, multiplier []int) []int {
	var reverseMultiplier = arraytools.ReverseInt(multiplier)
	var result = []int{0}

	multiLen := len(reverseMultiplier)
	for zehnerPotenz := 0; zehnerPotenz < multiLen; zehnerPotenz++ {
		for i := 0; i < reverseMultiplier[zehnerPotenz]; i++ {
			result = AddInt(result, baseOn10Potent(base, zehnerPotenz))
		}
	}
	return result
}

func CompareInt(x []int, y []int) int {
	if len(x) < len(y) {
		return -1
	}
	if len(x) > len(y) {
		return 1
	}
	for i := len(x) - 1; i >= 0; i-- {
		if x[i] < y[i] {
			return -1
		}
		if x[i] > y[i] {
			return 1
		}
	}
	return 0
}

// FactorialInt - Calculates the factorial of number
func FactorialInt(number []int) []int {
	result := []int{1}

	for i := []int{2}; CompareInt(i, number) <= 0; i = AddInt(i, []int{1}) {
		result = MultiplyInt(result, i)
	}
	return result
}

func baseOn10Potent(base []int, zehnerPotenz int) []int {
	if zehnerPotenz == 0 {
		return base
	}
	var result []int
	for i := 1; i <= 10; i++ {
		result = AddInt(result, baseOn10Potent(base, zehnerPotenz-1))
	}
	return result
}

func getLongestLength(val1 []int, val2 []int) int {
	maxLen := len(val1)
	if maxLen < len(val2) {
		maxLen = len(val2)
	}
	return maxLen
}
