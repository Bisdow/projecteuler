package specialpythagoreantriplet

import (
	"math"
)

const targetSum = 1000

func ExecProjectEulerProblem() {
	const target = 1000
	println(SpecialPythagoreanTriplet(target))
}

// SpecialPythagoreanTriplet - Returns the product of a * b * c for a^2 + b^2 = c^2 with a+b+c = INPUT abcEqual
func SpecialPythagoreanTriplet(abcEqualTo int) int {
	for c := abcEqualTo - 3; c > 0; c-- {
		for b := 2; b < c; b++ {
			for a := 1; a < b; a++ {
				if (a + b + c) != abcEqualTo {
					continue
				}
				if checkPythagoreanTriplet(a, b, c) {
					return a * b * c
				}
			}
		}
	}
	return 0
}

func checkPythagoreanTriplet(a, b, c int) bool {
	return math.Pow(float64(a), 2)+math.Pow(float64(b), 2) == math.Pow(float64(c), 2)
}
