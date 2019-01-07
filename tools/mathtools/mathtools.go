package mathtools

import (
	"math"
	"sort"

	"github.com/Bisdow/projecteuler/tools/arraytools"
	"github.com/Bisdow/projecteuler/tools/primenumbers"
)

// FindDivisorsOf - Returns an Array with all Divisors of the einput
func FindDivisorsOf(value int) []int {
	var primeArr = primenumbers.GetPrimeArray()
	var divisors []int
	if value == 0 {
		return []int{0}
	}
	if value == 1 {
		return []int{1}
	}
	divisors = []int{1, value}
	var maxDividable = getMaxDividable(value)
	for maxDividable > arraytools.LastElementInt(primeArr) {
		primeArr = primenumbers.NextPrime()
	}
	for i := 0; i < len(primeArr) && primeArr[i] <= maxDividable; i++ {
		prime := primeArr[i]
		if value%prime == 0 {
			divisor := prime
			for divisor <= maxDividable {
				if value%divisor == 0 && !arraytools.ElementInListInt(divisor, divisors) {
					divisors = append(divisors, divisor)
					if divisor*divisor != value {
						divisors = append(divisors, value/divisor)
					}
				}
				divisor += prime
			}
		}
	}
	sort.Ints(divisors)
	return divisors
}

// Factorial - returns the factorial of the input => X!
func Factorial(value uint64) uint64 {
	if value == 0 || value == 1 {
		return 1
	}

	result := uint64(2)
	for i := uint64(3); i <= value; i++ {
		result *= i
	}
	return result
}

func getMaxDividable(triangleSum int) int {
	return int(math.Sqrt(float64(triangleSum)))
}

func lastElement(arr []int) int {
	return arr[len(arr)-1]
}
