package summationofprimes

import (
	"fmt"

	"github.com/Bisdow/projecteuler/tools/primenumbers"
)

// ExecProjectEulerProblem - Execute Problem 10 Solution
func ExecProjectEulerProblem() {
	const limit int = 2000000
	fmt.Println("Ergebnis: ", SummationOfPrimes(limit))
}

// SummationOfPrimes -- Sum of all primes that are lower than input target
func SummationOfPrimes(targetLimit int) int {
	var primArr = primenumbers.GetPrimeArray()
	for lastElement(primArr) < targetLimit {
		primArr = primenumbers.NextPrime()
	}
	var result int
	for i := 0; i < len(primArr)-1 && primArr[i] < targetLimit; i++ {
		result += primArr[i]
	}
	return result
}

func lastElement(arr []int) int {
	return arr[len(arr)-1]
}
