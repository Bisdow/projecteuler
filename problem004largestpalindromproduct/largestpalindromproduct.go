package problem004largestpalindromproduct

import (
	"fmt"
)

func ExecProjectEulerProblem() {
	const lowLimit int = 100
	const maxLimit int = 999
	var maxX, maxY = FindHighestPalindrom(lowLimit, maxLimit)
	fmt.Println(maxX, " * ", maxY, "=", maxX*maxY)
}

func FindHighestPalindrom(lowLimit int, maxLimit int) (int, int) {
	x := maxLimit
	y := maxLimit
	maxX := 0
	maxY := 0
	maxProduct := 0
	for ; x >= lowLimit; x-- {
		y = maxLimit
		for ; y >= x; y-- {
			product := x * y
			if product%11 != 0 {
				continue
				/*
				 Palindrom => abccba
				 => 1a + 10b + 100c + 1000c + 10000b + 100000a
				 => 100001a + 10010b + 1100c
				 => 11 * (9091a + 910b + 100c)
				*/
			}
			if product > maxProduct {
				if checkPalindrom(product) {
					maxX = x
					maxY = y
					maxProduct = product
				}
			}
		}
	}
	return maxX, maxY
}

func checkPalindrom(number int) bool {
	numberArr := getNumberArray(number)
	if len(numberArr) < 2 {
		return false
	}

	begin := 0
	end := len(numberArr) - 1

	for begin < end {
		if numberArr[begin] != numberArr[end] {
			return false
		}
		begin++
		end--
	}

	return true
}

func getNumberArray(number int) []int {
	highestDivider := 0
	for i := 1; i <= number; i *= 10 {
		highestDivider = i
	}

	var numberArr []int
	tmp := number
	for i := highestDivider; i >= 1; i /= 10 {
		numberArr = append(numberArr, tmp/i)
		tmp = number % i
	}
	return numberArr
}
