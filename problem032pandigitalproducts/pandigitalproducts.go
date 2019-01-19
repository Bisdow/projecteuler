package problem032pandigitalproducts

import (
	"strconv"

	"github.com/Bisdow/projecteuler/tools/arraytools"
)

// ExecProjectEulerProblem -- Problem 32
func ExecProjectEulerProblem() {
	println(findPandigitalProductsSum())
}

func findPandigitalProductsSum() int {
	results := []int{}
	result := 0
	for i := 1; i < 9876; i++ {
		for j := i; j < 9876; j++ {
			if isPandigitalProduct(i, j) {
				value := i * j
				if !arraytools.ElementInListInt(value, results) {
					results = append(results, value)
					result += value
				}
			}
		}
	}
	return result
}

func isPandigitalProduct(a, b int) bool {
	if !checkNumber(a) || !checkNumber(b) {
		return false
	}
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	for i := 0; i < len(aStr); i++ {
		digitA := aStr[i : i+1]
		for j := 0; j < len(bStr); j++ {
			digitB := bStr[j : j+1]
			if digitA == digitB {
				return false
			}
		}
	}
	product := a * b
	if !checkNumber(product) {
		return false
	}
	productStr := strconv.Itoa(product)
	if len(aStr)+len(bStr)+len(productStr) != 9 {
		return false
	}
	for i := 0; i < len(aStr); i++ {
		digitA := aStr[i : i+1]
		for j := 0; j < len(productStr); j++ {
			digitProduct := productStr[j : j+1]
			if digitA == digitProduct {
				return false
			}
		}
	}
	for i := 0; i < len(bStr); i++ {
		digitB := bStr[i : i+1]
		for j := 0; j < len(productStr); j++ {
			digitProduct := productStr[j : j+1]
			if digitB == digitProduct {
				return false
			}
		}
	}
	return true
}

func checkNumber(a int) bool {
	aStr := strconv.Itoa(a)
	for i := 0; i < len(aStr); i++ {
		digitA := aStr[i : i+1]
		if digitA == "0" {
			return false
		}
		for j := i + 1; j < len(aStr); j++ {
			if digitA == aStr[j:j+1] {
				return false
			}
		}
	}
	return true
}
