package problem007xteprime

import (
	"fmt"

	"github.com/Bisdow/projecteuler/tools/primenumbers"
)

// ExecProjectEulerProblem - Start calculation for Problem
func ExecProjectEulerProblem() {
	const targetPrime = 10001
	fmt.Println(FindPrimeNumber(targetPrime))
}

// FindPrimeNumber - Find the xte prime number
func FindPrimeNumber(targetPrime int) int {
	var primArr = primenumbers.GetPrimeArray()
	length := len(primArr)
	if targetPrime < length {
		return primArr[targetPrime-1]
	}

	for i := len(primArr); i <= targetPrime; i++ {
		primArr = primenumbers.NextPrime()
	}
	return primArr[targetPrime-1]
}
