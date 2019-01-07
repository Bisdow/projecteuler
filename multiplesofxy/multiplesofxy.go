package multiplesofxy

import (
	"fmt"
)

func ExecProjectEulerProblem() {
	const limit int = 1000
	const firstDivider int = 3
	const secondDivider int = 5
	var result = SumOfMultiples(limit, firstDivider, secondDivider)
	fmt.Printf("Result is %v\n", result)
}

func SumOfMultiples(limit int, firstDivider int, secondDivider int) int {
	var result int
	for i := 0; i < limit; i++ {
		firstResult := i % firstDivider
		if firstResult == 0 {
			result += i
		} else {
			secondResult := i % secondDivider
			if secondResult == 0 {
				result += i
			}
		}
	}
	return result
}
