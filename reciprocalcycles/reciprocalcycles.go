package reciprocalcycles

import (
	"github.com/Bisdow/projecteuler/tools/arraytools"
)

func ExecProjectEulerProblem() {
	const limit = 1000
	println(calculateReciprocalCycle(limit))
}

func calculateReciprocalCycle(limit int) int {
	var maxLengthOfCycle = 0
	var dividerWithMaxCycle = 0
	for i := 2; i < limit; i++ {
		cycleLength := findLengthOfRecurringCycle(i)
		if cycleLength > maxLengthOfCycle {
			maxLengthOfCycle = cycleLength
			dividerWithMaxCycle = i
		}
	}
	return dividerWithMaxCycle
}

func findLengthOfRecurringCycle(divider int) int {
	var decimalDigits = []int{}
	var bisherigeReste = []int{1}
	var rest = 1
	for rest != 0 {
		if rest >= divider || rest <= 0 {
			println("PANIC PANIC PANIC --- Rest = ", rest)
			panic("Rest außerhalb von 1-9 ist sehr unerwartet")
		}
		rest *= 10
		divided := rest / divider
		for divided == 0 {
			decimalDigits = append(decimalDigits, rest/divider)
			bisherigeReste = append(bisherigeReste, 0)
			rest *= 10
			divided = rest / divider
		}
		decimalDigits = append(decimalDigits, rest/divider)
		rest = rest % divider

		if rest == 0 {
			continue
		}
		indexInReste := arraytools.FindIndexOfElementInt(rest, bisherigeReste)
		if indexInReste == -1 {
			bisherigeReste = append(bisherigeReste, rest)
			continue
		}

		lengthOfCycle := len(bisherigeReste) - indexInReste
		return lengthOfCycle
	}
	return 0 // no recurring cycle
}
