package problem021amicablenumbers

import (
	"github.com/Bisdow/projecteuler/tools/mathtools"
)

// ExecProjectEulerProblem - Problem 21
func ExecProjectEulerProblem() {
	const target = 10000
	println(SumOfAmicableNumbersBelow(target))
}

func SumOfAmicableNumbersBelow(target int) int {
	divisorSums := []int{}
	for i := 0; i < target; i++ {
		divisorSums = append(divisorSums, sumOfDivisors(i))
	}
	result := 0
	for i := 2; i < target; i++ {
		value1 := divisorSums[i]
		if value1 >= target {
			continue
		}
		if value1 <= i {
			continue
		}
		value2 := divisorSums[value1]
		if i == value2 {
			result += i
			result += value1
		}
	}
	return result
}

func sumOfDivisors(value int) int {
	divisors := mathtools.FindDivisorsOf(value)
	result := 0
	// -1 to ignore the value itself
	for i := 0; i < len(divisors)-1; i++ {
		result += divisors[i]
	}
	return result
}
