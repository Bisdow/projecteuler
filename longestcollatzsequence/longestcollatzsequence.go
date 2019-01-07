package longestcollatzsequence

import "fmt"

// ExecProjectEulerProblem - Calculate Solution for Problem 14
func ExecProjectEulerProblem() {
	const startPoint int = 1000000
	chainStart, chainLength := GetLongestChain(startPoint)
	fmt.Println("Chain-Start: ", chainStart, " Länge: ", chainLength)
}

// GetLongestChain - Get Longest Collatz Sequenz
func GetLongestChain(startPoint int) (int, int) {
	maxChainLength := 0
	maxStartingPoint := 0
	for i := startPoint; i > 0; i-- {
		chainLength := getChainLength(i)
		if chainLength > maxChainLength {
			maxChainLength = chainLength
			maxStartingPoint = i
		}
	}
	return maxStartingPoint, maxChainLength
}

func getChainLength(startPoint int) int {
	var tmp = startPoint
	var counter = 1
	for tmp > 1 {
		tmp = newValue(tmp)
		counter++
	}
	return counter
}

func newValue(value int) int {
	if value%2 == 0 {
		return value / 2
	} else {
		return 3*value + 1
	}
}
