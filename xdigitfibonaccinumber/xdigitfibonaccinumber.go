package xdigitfibonaccinumber

import (
	"math/big"
)

// ExecProjectEulerProblem - Problem 25
func ExecProjectEulerProblem() {
	const limitDigits = 1000
	println(xteFibonacciNumber(limitDigits))
}

func xteFibonacciNumber(limitDigits int) uint64 {
	var currentFibonacci = big.NewInt(1)
	var tmp1 = big.NewInt(1)
	var tmp2 = big.NewInt(1)

	var index uint64 = 2

	digits := len(currentFibonacci.String())
	for digits < limitDigits {
		currentFibonacci.Add(tmp1, tmp2)
		tmp1.Set(tmp2)
		tmp2.Set(currentFibonacci)
		digits = len(currentFibonacci.String())
		index++
	}
	return index
}
