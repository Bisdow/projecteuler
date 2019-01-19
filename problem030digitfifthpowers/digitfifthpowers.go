package problem030digitfifthpowers

import (
	"math"
	"strconv"
)

func ExecProjectEulerProblem() {
	const power = 5
	const limit = 1000000
	println(findDigitPowerSum(power, limit))
}

func findDigitPowerSum(power int, limit uint64) uint64 {
	var powerArray [10]uint64
	for i := 0; i <= 9; i++ {
		powerArray[i] = uint64(math.Pow(float64(i), float64(power)))
	}

	result := uint64(0)
	for i := uint64(2); i <= limit; i++ {
		if isEqualDigitPower(i, powerArray) {
			result += i
		}
	}
	return result
}

func isEqualDigitPower(value uint64, powerArray [10]uint64) bool {
	valueStr := strconv.FormatUint(value, 10)
	result := uint64(0)
	for i := 0; i < len(valueStr); i++ {
		digit, err := strconv.Atoi(valueStr[i : i+1])
		if err != nil {
			panic("Oh oh")
		}
		result += powerArray[digit]
	}
	// fmt.Printf("%v => %v\n", value, result)
	return result == value
}
