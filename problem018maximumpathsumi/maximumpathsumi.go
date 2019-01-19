package problem018maximumpathsumi

import "fmt"

func ExecProjectEulerProblem() {
	var triangle = [][]int{
		[]int{75},
		[]int{95, 64},
		[]int{17, 47, 82},
		[]int{18, 35, 87, 10},
		[]int{20, 4, 82, 47, 65},
		[]int{19, 1, 23, 75, 3, 34},
		[]int{88, 2, 77, 73, 7, 63, 67},
		[]int{99, 65, 4, 28, 6, 16, 70, 92},
		[]int{41, 41, 26, 56, 83, 40, 80, 70, 33},
		[]int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		[]int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		[]int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		[]int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		[]int{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		[]int{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23}}

	// var triangle = [][]int{
	// 	[]int{3},
	// 	[]int{7, 4},
	// 	[]int{2, 4, 6},
	// 	[]int{8, 5, 9, 3}} // Expected 3+7+4+9 = 23

	fmt.Println(findMaximumTotal(triangle))
}

func findMaximumTotal(triangle [][]int) int {
	var length = len(triangle)
	lastLine := triangle[0]
	for i := 1; i < length; i++ {
		lastLine = calcNextLevel(lastLine, triangle[i])
	}
	length = len(lastLine)
	max := 0
	for i := 0; i < length; i++ {
		if lastLine[i] > max {
			max = lastLine[i]
		}
	}
	return max
}

func calcNextLevel(base []int, next []int) []int {
	var len = len(next)
	var result []int
	for i := 0; i < len; i++ {
		if i == 0 {
			result = append(result, base[i]+next[i])
		} else if i == len-1 {
			result = append(result, base[i-1]+next[i])
		} else {
			left := base[i-1] + next[i]
			right := base[i] + next[i]
			if left > right {
				result = append(result, left)
			} else {
				result = append(result, right)
			}
		}
	}
	return result
}
