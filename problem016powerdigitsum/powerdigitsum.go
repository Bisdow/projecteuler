package problem016powerdigitsum

import (
	"github.com/Bisdow/projecteuler/tools/bignumbers"
)

func ExecProjectEulerProblem() {
	const base = 2
	const potent = 1000

	println(calcSumOfDigits(base, potent))
}

func calcSumOfDigits(base int, potent int) int {
	bigBase := bignumbers.ParseInt(base)
	bigNumber := bigBase
	for i := 1; i < potent; i++ {
		bigNumber = bignumbers.MultiplyInt(bigNumber, bigBase)
	}
	bigNumberLength := len(bigNumber)
	result := 0
	for i := 0; i < bigNumberLength; i++ {
		result += bigNumber[i]
	}
	return result
}
