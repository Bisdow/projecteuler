package problem031coinsums

// ExecProjectEulerProblem - Lösung für Problem 31
func ExecProjectEulerProblem() {
	println(calcWaysOfCoin(0, 200, []int{200, 100, 50, 20, 10, 5, 2, 1}))
}

func calcWaysOfCoin(sum int, limit int, coins []int) int {
	coinValue := coins[0]
	possibleWays := 0
	for i := 0; sum+i*coinValue <= limit; i++ {
		newSum := sum + i*coinValue
		if newSum < limit {
			if len(coins) > 1 {
				possibleWays += calcWaysOfCoin(newSum, limit, coins[1:])
			}
		}
		if newSum == limit {
			possibleWays++
		}
	}
	return possibleWays
}
