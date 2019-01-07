package coinsums

import "testing"

func TestCalcWaysOfCoin(t *testing.T) {
	var limit int
	var coinValues []int
	var expected int

	limit = 1
	coinValues = []int{1}
	expected = 1
	checkCalcWaysOfCoin(limit, coinValues, expected, t)

	limit = 2
	coinValues = []int{1}
	expected = 1
	checkCalcWaysOfCoin(limit, coinValues, expected, t)

	limit = 3
	coinValues = []int{1, 2}
	expected = 2
	checkCalcWaysOfCoin(limit, coinValues, expected, t)

	limit = 4
	coinValues = []int{1, 2}
	expected = 3
	checkCalcWaysOfCoin(limit, coinValues, expected, t)
}

func checkCalcWaysOfCoin(limit int, coinValues []int, expected int, t *testing.T) {
	result := calcWaysOfCoin(0, limit, coinValues)
	if result != expected {
		t.Errorf("Limit: %v, CoinValues: %v, Expected: %v", limit, coinValues, expected)
	}
}
