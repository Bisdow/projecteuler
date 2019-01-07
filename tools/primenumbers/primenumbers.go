package primenumbers

import (
	"math"

	"github.com/Bisdow/projecteuler/tools/arraytools"
)

var primeArray = []int{2, 3}

// ResetPrimenumbers - Setze alle ermittelten Primzahlen zurück
func ResetPrimenumbers() {
	primeArray = []int{2, 3}
}

// GetPrimeArray - Returns Array with known Prime numbers
func GetPrimeArray() []int {
	return primeArray
}

// GetPrimeArraySlice - Returns the amount of prime numbers
func GetPrimeArraySlice(high int) []int {
	for len(primeArray) < high {
		NextPrime()
	}
	return primeArray[0:high]
}

// GetPrimeNumbersUpTo - Returns the Prime numbers up to at least the needed value
func GetPrimeNumbersUpTo(value int) []int {
	var i = 0
	for primeArray[i] <= value {
		if i == len(primeArray)-1 {
			NextPrime()
		}
		i++
	}

	if primeArray[i-1] < value {
		i++
		if i > len(primeArray)-1 {
			NextPrime()
		}
	}
	return primeArray[0:i]
}

// NextPrime - Add the next Prime number to the PrimeArray and returns it
func NextPrime() []int {
	var result = lastElement(primeArray)
	var istPrime = false
	for istPrime == false {
		result += 2
		istPrime = isPrime(result)
	}
	primeArray = append(primeArray, result)
	return primeArray
}

// GetFactors - Ermittelt die Primefaktoren einer Zahl
func GetFactors(value int) []int {
	var factors []int
	var tmp = value
	for i := 0; i < len(primeArray) && primeArray[i] <= value; i++ {
		for tmp > 0 && tmp%primeArray[i] == 0 {
			factors = append(factors, primeArray[i])
			tmp = tmp / primeArray[i]
		}
		if i == len(primeArray)-1 && tmp > 1 {
			NextPrime()
		}
	}
	return factors
}

// IsPrime - Checks if the given value is a prime number
func IsPrime(value int) bool {
	for value > lastElement(primeArray) {
		NextPrime()
	}
	return arraytools.ElementInListInt(value, primeArray)
}

func isPrime(value int) bool {
	isPrime := true
	for i := 0; primeArray[i] <= maxDivisor(value); i++ {
		if value%primeArray[i] == 0 {
			isPrime = false
			break
		}
	}
	return isPrime
}

func maxDivisor(value int) int {
	return int(math.Sqrt(float64(value)))
}

func lastElement(arr []int) int {
	return arr[len(arr)-1]
}
