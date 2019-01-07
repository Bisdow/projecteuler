package digitfactorials

import (
	"strconv"

	"github.com/Bisdow/projecteuler/tools/mathtools"
)

func ExecProjectEulerProblem() {
	println(digitfactorials())
}

/* 145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
Find the sum of all numbers which are equal to the sum of the factorial of their digits.
Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/

var factorials0To9 [10]int

func digitfactorials() int {
	factorials0To9 = calcFactorials()
	const limit = 2540160 // 9! has 6 digits => with 7digits * 9! it grows less then then the number
	result := 0
	for i := 3; i <= limit; i++ {
		if checkNumber(i) {
			result += i
		}
	}
	return result
}

func checkNumber(value int) bool {
	sum := 0
	str := strconv.Itoa(value)
	for i := 0; i < len(str); i++ {
		numberStr := str[i : i+1]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic("Could not convert to number")
		}
		sum += factorials0To9[number]
		if sum > value {
			return false
		}
	}
	return value == sum
}

func calcFactorials() [10]int {
	result := [10]int{}
	for i := 0; i <= 9; i++ {
		result[i] = int(mathtools.Factorial(uint64(i)))
	}
	return result
}
