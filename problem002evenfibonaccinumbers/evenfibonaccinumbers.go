package problem002evenfibonaccinumbers

import (
	"fmt"
)

func ExecProjectEulerProblem() {
	const limit int = 4000000
	result := EvenValuedFibonacciNumbers(limit)
	fmt.Printf("Result is %v\n", result)
}

func EvenValuedFibonacciNumbers(limit int) int {
	if limit <= 1 {
		return 0
	}
	var result int
	var tmp int
	var tmp1 = 1
	var tmp2 = 1

	for tmp < limit {
		tmp = tmp1 + tmp2
		tmp1 = tmp2
		tmp2 = tmp
		if tmp%2 == 0 {
			result += tmp2
		}
	}
	return result
}
