package problem005smallestmultiple

import (
	"fmt"
)

// kleinstmögliche Zahl die durch 1,2,3,4,...20 teilbar ist

func ExecProjectEulerProblem() {
	const maxDivider int = 20
	result := FindSmallestMultiple(maxDivider)
	fmt.Println(result)
}

func FindSmallestMultiple(maxDivider int) int {
	if maxDivider == 1 {
		return 1
	}
	value := 0
	result := 0
	valueStepper := findValueStepper(maxDivider)
	divValues := findDivValues(maxDivider)
	divLen := len(divValues)

	var wasNotDividable bool
	for result == 0 {
		value += valueStepper
		wasNotDividable = false
		for i := 0; i < divLen; i++ {
			if value%divValues[i] != 0 {
				wasNotDividable = true
				break
			}
		}
		if !wasNotDividable {
			result = value
		}
	}
	return result
}

// Zahl ist nur durch 20 und 19 teilbar wenn es eine Vielfache von 20*19 = 380 ist
func findValueStepper(maxDivider int) int {
	return maxDivider * (maxDivider - 1)
}

// Wenn Zahl durch 20 teilbar ist, dann sicher auch durch 2,4,5,10
// Wenn Zahl durch 18 teilbar ist, dann sicher auch durch 3, 6, 9
// Wenn Zahl durch 16 teilbar ist, dann sicher auch durch 8
// Wenn Zahl durch 14 teilbar ist, dann sicher auch durch 7
// => Zahl muss durch die Zahlen von 20 bis 11 teilbar sein
func findDivValues(maxDivider int) []int {
	notNeeded := []int{}

	for i := 2; i < maxDivider; i++ {
		if maxDivider%i != 0 {
			notNeeded = append(notNeeded, i)
		}
	}
	return notNeeded
}
