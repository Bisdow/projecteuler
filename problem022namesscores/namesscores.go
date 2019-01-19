package problem022namesscores

import (
	"io/ioutil"
	"sort"
	"strings"
)

// ExecProjectEulerProblem - Problem 22
func ExecProjectEulerProblem() {
	const file = "./namesscores/names.txt"
	println(NamesScores(file))
}

func NamesScores(filename string) int64 {
	nameList := readDataAndPrepare(filename)

	return scoreSumOverArray(nameList)
}

func scoreSumOverArray(nameList []string) int64 {
	var result int64
	for i := 0; i < len(nameList); i++ {
		result += int64(i+1) * valueOfName(nameList[i])
	}
	return result
}

func readDataAndPrepare(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	check(err)
	str := strings.Replace(string(data), "\"", "", 0)
	nameList := strings.Split(str, ",")
	sort.Strings(nameList)
	return nameList
}

func valueOfName(name string) int64 {
	var result int64
	for i := 0; i < len(name); i++ {
		switch name[i : i+1] {
		case "A":
			result += 1
		case "B":
			result += 2
		case "C":
			result += 3
		case "D":
			result += 4
		case "E":
			result += 5
		case "F":
			result += 6
		case "G":
			result += 7
		case "H":
			result += 8
		case "I":
			result += 9
		case "J":
			result += 10
		case "K":
			result += 11
		case "L":
			result += 12
		case "M":
			result += 13
		case "N":
			result += 14
		case "O":
			result += 15
		case "P":
			result += 16
		case "Q":
			result += 17
		case "R":
			result += 18
		case "S":
			result += 19
		case "T":
			result += 20
		case "U":
			result += 21
		case "V":
			result += 22
		case "W":
			result += 23
		case "X":
			result += 24
		case "Y":
			result += 25
		case "Z":
			result += 26
		}
	}
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
