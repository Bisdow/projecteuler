package problem023nonabundantsums

import (
	"github.com/Bisdow/projecteuler/tools/mathtools"
)

// ExecProjectEulerProblem - Problem 23
func ExecProjectEulerProblem() {
	const limit = 28123
	println(sumOfIntegersNotSumOfTwoAbundants(limit))
}

func sumOfIntegersNotSumOfTwoAbundants(limit int) int {
	abundantNumbers := []int{}
	sum := 0

	for i := 1; i <= limit; i++ {
		if checkIsAbundant(i) {
			abundantNumbers = append(abundantNumbers, i)

		}
		if !checkIfIsSumOfAbundants(i, abundantNumbers) {
			sum += i
		}
	}
	return sum
}

func checkIsAbundant(value int) bool {
	divisors := mathtools.FindDivisorsOf(value)
	sum := 0
	for i := 0; i < len(divisors)-1; i++ {
		sum += divisors[i]
	}
	return sum > value
}

func checkIfIsSumOfAbundants(value int, abundantNumbers []int) bool {
	for i := 0; i < len(abundantNumbers) && abundantNumbers[i] < value; i++ {
		for j := i; j < len(abundantNumbers) && abundantNumbers[j] < value; j++ {
			if abundantNumbers[i]+abundantNumbers[j] == value {
				return true
			}
		}
	}
	return false
}
