package distinctpowers

import (
	"math/big"
)

// ExecProjectEulerProblem - Problem 29
func ExecProjectEulerProblem() {
	const maxA = 100
	const maxB = 100
	println(distinctPowers(maxA, maxB))
}

func distinctPowers(aMax, bMax int) int {
	solutions := []*big.Int{}
	for a := 2; a <= aMax; a++ {
		for b := 2; b <= bMax; b++ {
			aBig := big.NewInt(int64(a))
			bBig := big.NewInt(int64(b))
			result := big.NewInt(0)
			result.Exp(aBig, bBig, nil)
			allreadyFound := false
			for i := 0; i < len(solutions); i++ {
				if result.Cmp(solutions[i]) == 0 {
					allreadyFound = true
					break
				}
			}
			if !allreadyFound {
				solutions = append(solutions, result)
			}
		}
	}
	return len(solutions)
}
