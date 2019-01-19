package maximumpathsumii

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ExecProjectEulerProblem - Problem 67
func ExecProjectEulerProblem() {
	const filename = "./maximumpathsumii/triangle.txt"
	println(Maximumpathsumii(filename))
}

func Maximumpathsumii(filename string) int {
	triangle := readDataAndPrepare(filename)

	return findMaximumTotal(triangle)
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
	var length = len(next)
	var result []int
	for i := 0; i < length; i++ {
		if i == 0 {
			result = append(result, base[i]+next[i])
		} else if i == length-1 {
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

func readDataAndPrepare(filename string) [][]int {
	f, err := os.Open(filename)
	check(err)
	defer func() {
		err = f.Close()
		check(err)
	}()

	var triangle [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var line []int
		stringLine := scanner.Text()
		stringArr := strings.Split(stringLine, " ")
		for i := 0; i < len(stringArr); i++ {
			value, err := strconv.Atoi(stringArr[i])
			if err != nil {
				panic("ERROR while reading")
			}
			line = append(line, value)
		}
		triangle = append(triangle, line)
	}
	return triangle
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
