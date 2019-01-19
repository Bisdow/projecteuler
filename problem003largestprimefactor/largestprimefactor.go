package problem003largestprimefactor

import (
	"fmt"

	"github.com/Bisdow/projecteuler/tools/primenumbers"
)

func ExecProjectEulerProblem() {
	const limit = 600851475143
	result := FindLargestPrimeFactor(limit)
	fmt.Printf("Höchster Primefaktor von %v ist %v\n", limit, result)
}

func FindLargestPrimeFactor(value int) int {
	factors := primenumbers.GetFactors(value)
	return lastElement(factors)
}

func lastElement(arr []int) int {
	return arr[len(arr)-1]
}
