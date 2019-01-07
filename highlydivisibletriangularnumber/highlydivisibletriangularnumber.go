package highlydivisibletriangularnumber

import (
	"fmt"
	"math"

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

func elementInList(element int, list []int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == element {
			return true
		}
	}
	return false
}

func getMaxDividable(triangleSum int) int {
	return int(math.Sqrt(float64(triangleSum)))
}

func lastElement(arr []int) int {
	return arr[len(arr)-1]
}
