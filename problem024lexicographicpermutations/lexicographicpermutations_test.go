package problem024lexicographicpermutations

import (
	"testing"

	"github.com/Bisdow/projecteuler/tools/arraytools"
)

func TestDoPermutation(t *testing.T) {
	checkDoPermutation([]int{0, 1}, []int{1, 0}, t)
	checkDoPermutation([]int{0, 1, 2}, []int{0, 2, 1}, t)
	checkDoPermutation([]int{0, 2, 1}, []int{1, 0, 2}, t)
	checkDoPermutation([]int{1, 0, 2}, []int{1, 2, 0}, t)
	checkDoPermutation([]int{1, 2, 0}, []int{2, 0, 1}, t)
	checkDoPermutation([]int{2, 0, 1}, []int{2, 1, 0}, t)
}

func checkDoPermutation(input []int, expected []int, t *testing.T) {
	result := doPermutation(input)
	if !arraytools.EqualInt(result, expected) {
		t.Errorf("Input: %v, Result: %v, Expected: %v", input, result, expected)
	}
}

func TestLexicographicPermutations(t *testing.T) {
	checklexicographicPermutations(1, []int{0, 1, 2}, []int{0, 1, 2}, t)
	checklexicographicPermutations(2, []int{0, 1, 2}, []int{0, 2, 1}, t)
	checklexicographicPermutations(4, []int{0, 1, 2}, []int{1, 2, 0}, t)
	checklexicographicPermutations(6, []int{0, 1, 2}, []int{2, 1, 0}, t)
}

func checklexicographicPermutations(position int, numbers []int, expected []int, t *testing.T) {
	result := lexicographicPermutations(position, numbers)
	if !arraytools.EqualInt(result, expected) {
		t.Errorf("Input: %v, Result: %v, Expected: %v", position, result, expected)
	}
}

/*
0 1
1 0

0 1 2
0 2 1
1 0 2
1 2 0
2 0 1
2 1 0

0 1 2 3
0 1 3 2
0 2 1 3
0 2 3 1
0 3 1 2
0 3 2 1
1 0 2 3
1 0 3 2
1 2 0 3
1 2 3 0
1 3 0 2
1 3 2 0
2 0 1 3
2 0 3 1
2 1 0 3
2 1 3 0
2 3 0 1
2 3 1 0
3 0 1 2
3 0 2 1
3 1 0 2
3 1 2 0
3 2 0 1
3 2 1 0
*/
