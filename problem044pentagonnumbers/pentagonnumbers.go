package problem044pentagonnumbers

import (
	"fmt"
)

var pentaArr = []uint64{1}

func ExecProjectEulerProblem() {
	println("Result: ", FindMinimumPentagonalDifference())
}

func FindMinimumPentagonalDifference() uint64 {
	var minDifference uint64
	for i := 1; ; i++ {
		if len(pentaArr) <= i {
			calcNextPenta()
		}
		delta := pentaArr[i] - pentaArr[i-1]
		if minDifference != 0 && delta > minDifference {
			break
		}
		penta1 := pentaArr[i]
		for j := 0; j < i; j++ {
			penta2 := pentaArr[j]
			if penta2 < delta {
				continue
			}
			difference := penta1 - penta2
			if !elementIsPenta(difference) {
				continue
			}
			if !checkAdditionIsPenta(penta1, penta2) {
				continue
			}
			fmt.Println("Found ", difference, "from ", penta1, " - ", penta2, "with i = ", i)
			if minDifference == 0 || minDifference > difference {
				minDifference = difference
				// TODO: Replace with better "loop break rule". Currently it runs Very very VERY LONG
				return minDifference
			}
		}
	}
	return minDifference
}

func checkAdditionIsPenta(penta1 uint64, penta2 uint64) bool {
	sum := penta1 + penta2

	for sum > lastElement() {
		calcNextPenta()
	}
	return elementIsPenta(sum)
}

func elementIsPenta(value uint64) bool {
	found := false
	for i := 0; i < len(pentaArr); i++ {
		if pentaArr[i] == value {
			found = true
			break
		}
	}
	return found
}

func calcNextPenta() {
	maxPentaNr := len(pentaArr)
	newPentaNr := lastElement() + uint64(maxPentaNr)*uint64(3) + uint64(1)
	pentaArr = append(pentaArr, newPentaNr)
}

func lastElement() uint64 {
	return pentaArr[len(pentaArr)-1]
}

/*
Pentagonal Numbers: 1 , 5, 12, 22, 35, 51 , 70, 92
1 -> 2 =  +4 	= 1 + 1*3+1 = 5
2 -> 3 = +7 	= 5 + 2*3+1 = 12
3 -> 4 = +10	= 12 + 3*3+1 = 22
=> last + LastPosition * 3 +1 => next

5482660 = n*(3n-1)/2 => (3n^2 - n)/2 =>
10965320 = 3n^2 - n
1911 = n - (n/2)^1/2
*/
