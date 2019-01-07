package numberlettercounts

func ExecProjectEulerProblem() {
	const numberLimit = 1000
	initHashMap()
	println(getWordLengthAllNumbers(numberLimit))
}

func getWordLengthAllNumbers(limit int) int {
	result := 0
	for i := 1; i <= limit; i++ {
		result += getWordLengthOfNumber(i)
	}
	return result
}

var numberHash = make(map[int]int)

func getWordLengthOfNumber(number int) int {
	result := 0
	value := number
	if value >= 1000 && value < 10000 {
		result += numberHash[value/1000]
		result += numberHash[1000]
		value %= 1000
	}
	if value >= 100 && value < 1000 {
		// 154 => 100
		result += numberHash[value/100]
		result += numberHash[100]
		value %= 100
		if value > 0 {
			result += 3 // and
		}
	}
	if value >= 20 && value < 100 {
		// 21 => 20
		result += numberHash[value-value%10]
		value %= 10
	}
	if value < 20 && value > 0 {
		result += numberHash[value]
	}
	return result
}

func initHashMap() {
	numberHash[1] = 3    // one
	numberHash[2] = 3    // two
	numberHash[3] = 5    // three
	numberHash[4] = 4    // four
	numberHash[5] = 4    // five
	numberHash[6] = 3    // six
	numberHash[7] = 5    // seven
	numberHash[8] = 5    // eight
	numberHash[9] = 4    // nine
	numberHash[10] = 3   // ten
	numberHash[11] = 6   // eleven
	numberHash[12] = 6   // twelve
	numberHash[13] = 8   // thirteen
	numberHash[14] = 8   // fourteen
	numberHash[15] = 7   // fifteen
	numberHash[16] = 7   // sixteen
	numberHash[17] = 9   // seventeen
	numberHash[18] = 8   // eighteen
	numberHash[19] = 8   // nineteen
	numberHash[20] = 6   // 	twenty*
	numberHash[30] = 6   // 	thirty*
	numberHash[40] = 5   // 	forty*
	numberHash[50] = 5   // 	fifty*
	numberHash[60] = 5   // 	sixty*
	numberHash[70] = 7   // 	seventy*
	numberHash[80] = 6   // 	eighty*
	numberHash[90] = 6   // 	ninety*
	numberHash[100] = 7  // 	*hundred
	numberHash[1000] = 8 //		*thousand
}
