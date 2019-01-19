package problem028numberspiraldiagonals

// ExecProjectEulerProblem - Problem 28
func ExecProjectEulerProblem() {
	const targetSize = 1001
	println(calcDiagonals(targetSize))
}

func calcDiagonals(targetSize int) uint64 {
	diagonals := uint64(1)
	addValue := uint64(0)
	currentValue := uint64(1)
	for currentSize := 3; currentSize <= targetSize; currentSize += 2 {
		addValue += 2
		for i := 0; i < 4; i++ {
			currentValue += addValue
			diagonals += currentValue
		}
	}
	return diagonals
}
