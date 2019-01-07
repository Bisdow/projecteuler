package quadraticprimes

import "github.com/Bisdow/projecteuler/tools/primenumbers"

// ExecProjectEulerProblem - Problem 27
func ExecProjectEulerProblem() {
	maxA := 1000
	maxB := 1000
	println(findCoefficentsProduct(maxA, maxB))
}

func findCoefficentsProduct(maxA int, maxB int) int {
	var maximum int
	var maximumCoefficient int

	// Prefill primenumbers
	primes := primenumbers.GetPrimeNumbersUpTo(maxB)
	for a := -maxA; a <= maxA; a++ {
		for b := 0; b < len(primes); b++ {
			found := findConsecutivePrimes(a, -primes[b])
			if found > maximum {
				maximum = found
				maximumCoefficient = a * -primes[b]
			}
			found = findConsecutivePrimes(a, primes[b])
			if found > maximum {
				maximum = found
				maximumCoefficient = a * primes[b]
			}
		}
	}
	return maximumCoefficient
}

func findConsecutivePrimes(a int, b int) int {
	for n := 0; ; n++ {
		result := n*n + a*n + b
		if !primenumbers.IsPrime(result) {
			return n - 1
		}
	}
}
