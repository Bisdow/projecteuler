/*
1^2 + 2^2 					= 1 + 4 				= 5
1^2 + 2^2 + 3^2 			= 1 + 4 + 9 			= 14
1^2 + 2^2 + 3^2 + 4^2 		= 1 + 4 + 9 + 16		= 30
1^2 + 2^2 + 3^2 + 4^2 + 5^2	= 1 + 4 + 9 + 16 + 25	= 55

(1 + 2)^2				= 3^2	=	9
(1 + 2 + 3)^2			= 6^2	=	36
(1 + 2 + 3 + 4)^2		= 10^2	=	100
(1 + 2 + 3 + 4 + 5)^2	= 15^2	=	225

2 => 9 - 5 = 4
3 => 36 - 14 = 22
4 => 100 - 30 = 70
5 => 225 - 55 = 170

Gaussformel
(1 + 2 + ... + n) => n (n+1)/2
n = 100 => 101 * 50
*/

package problem006sumsquaredifference

import (
	"fmt"
	"math"
)

func ExecProjectEulerProblem() {
	fmt.Printf("Ergebnis (Differenz): %f\n", CalculateDelta(100))
}

func CalculateDelta(number float64) float64 {
	result := quadratDerSumme(number) - summeDerQuadrate(number)
	return result
}

func summeDerQuadrate(value float64) float64 {
	var result float64
	var number float64 = 1
	for ; number <= value; number++ {
		result += math.Pow(number, 2)
	}
	return result
}

func quadratDerSumme(value float64) float64 {
	var sum float64
	// Bruteforece
	/*var number float64 = 1
	for ; number <= value; number++ {
		sum += number
	}
	*/
	// Gausformel
	sum = (value + 1) * (value / 2)
	return math.Pow(sum, 2)
}
