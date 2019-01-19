package problem024lexicographicpermutations

import (
	"fmt"
	"sort"
)

// ExecProjectEulerProblem - Problem 24
func ExecProjectEulerProblem() {
	const position = 1000000
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(lexicographicPermutations(position, numbers))
}

func lexicographicPermutations(position int, numbers []int) []int {
	sort.Ints(numbers)
	for i := 1; i < position; i++ {
		numbers = doPermutation(numbers)
	}
	return numbers
}

func doPermutation(numbers []int) []int {
	// find rightest number, that has a biger right neighbor
	var value1 int
	var index1 int
	for i := len(numbers) - 2; i >= 0; i-- {
		value1 = numbers[i]
		index1 = i
		if numbers[i] < numbers[i+1] {
			break
		}
	}
	// find smallest number, bigger than the number we found before on the right side of the before found number
	var index2 int
	var minValue int
	for i := index1 + 1; i < len(numbers); i++ {
		if numbers[i] > value1 {
			if minValue == 0 || minValue > numbers[i] {
				minValue = numbers[i]
				index2 = i
			}
		}
	}
	if index2 == 0 {
		index2 = len(numbers) - 1
	}
	// Swap both found numbers
	numbers[index1], numbers[index2] = numbers[index2], numbers[index1]
	// sort the part right from the original index of the first found number
	sort.Ints(numbers[index1+1:])
	return numbers
}
