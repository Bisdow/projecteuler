package problem020factorialdigitsum

import (
	"github.com/Bisdow/projecteuler/tools/bignumbers"
)

func ExecProjectEulerProblem() {
	const number = 100

	println("The result is: ", Factorialdigitsum(number))
}

// Factorialdigitsum - Calculates the factorial of the input and then the sum of every digit of the result
func Factorialdigitsum(number int) int {
	value := bignumbers.ParseInt(number)
	factorial := bignumbers.FactorialInt(value)

	result := 0
	for i := 0; i < len(factorial); i++ {
		result += factorial[i]
	}
	return result
}
