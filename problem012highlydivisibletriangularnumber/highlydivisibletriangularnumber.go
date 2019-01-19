package problem012highlydivisibletriangularnumber

import (
	"fmt"
	"github.com/Bisdow/projecteuler/tools/mathtools"
)

// ExecProjectEulerProblem - Executes Problem 12
func ExecProjectEulerProblem() {
	const divisorTarget int = 500
	fmt.Println(Highlydivisibletriangularnumber(divisorTarget))
}

// Highlydivisibletriangularnumber -
func Highlydivisibletriangularnumber(divisorTarget int) int {
	var divisorAmount = 0
	var triangleElements = 1
	var triangleSum = 1

	for divisorAmount <= divisorTarget {
		triangleElements++
		triangleSum += triangleElements
		divisorAmount = amountOfDivisorsOf(triangleSum)
	}
	return triangleSum
}

func amountOfDivisorsOf(value int) int {
	divisors := mathtools.FindDivisorsOf(value)
	return len(divisors)
}
