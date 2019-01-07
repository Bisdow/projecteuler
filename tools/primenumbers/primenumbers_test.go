package primenumbers

/*
20.10.2018
BenchmarkNextPrime100-8                	  100000	     13355 ns/op	    2032 B/op	       7 allocs/op
BenchmarkNextPrime500-8                	   10000	    127247 ns/op	    8176 B/op	       9 allocs/op
BenchmarkNextPrime1000-8               	    5000	    334954 ns/op	   16368 B/op	      10 allocs/op
BenchmarkNextPrime10000-8              	     200	   7665521 ns/op	  386288 B/op	      19 allocs/op
BenchmarkGetFactors123-8               	 2000000	       886 ns/op	     264 B/op	       6 allocs/op
BenchmarkGetFactors1234-8              	  100000	     16228 ns/op	    2056 B/op	       9 allocs/op
BenchmarkGetFactors12345-8             	  100000	     23105 ns/op	    4136 B/op	      11 allocs/op
BenchmarkGetFactors123456-8            	  100000	     17335 ns/op	    2152 B/op	      11 allocs/op
BenchmarkGetPrimeNumbersUpTo100-8      	 1000000	      1855 ns/op	     496 B/op	       5 allocs/op
BenchmarkGetPrimeNumbersUpTo1000-8     	   50000	     27642 ns/op	    4080 B/op	       8 allocs/op
BenchmarkGetPrimeNumbersUpTo10000-8    	    3000	    452481 ns/op	   26608 B/op	      11 allocs/op
BenchmarkGetPrimeNumbersUpTo100000-8   	     200	   7357403 ns/op	  386288 B/op	      19 allocs/op
*/

import (
	"testing"

	"github.com/Bisdow/projecteuler/tools/arraytools"
)

func TestNextPrime(t *testing.T) {
	var result []int
	var expectedResult []int
	ResetPrimenumbers()
	result = GetPrimeArray()
	expectedResult = []int{2, 3}
	arrayCheck(t, "NextPrime()", result, expectedResult)
	result = NextPrime()
	expectedResult = []int{2, 3, 5}
	arrayCheck(t, "NextPrime()", result, expectedResult)
	result = NextPrime()
	expectedResult = []int{2, 3, 5, 7}
	arrayCheck(t, "NextPrime()", result, expectedResult)
	result = NextPrime()
	expectedResult = []int{2, 3, 5, 7, 11}
	arrayCheck(t, "NextPrime()", result, expectedResult)
	result = NextPrime()
	expectedResult = []int{2, 3, 5, 7, 11, 13}
	arrayCheck(t, "NextPrime()", result, expectedResult)
	result = NextPrime()
	expectedResult = []int{2, 3, 5, 7, 11, 13, 17}
	arrayCheck(t, "NextPrime()", result, expectedResult)
	result = NextPrime()
	expectedResult = []int{2, 3, 5, 7, 11, 13, 17, 19}
	arrayCheck(t, "NextPrime()", result, expectedResult)
	result = NextPrime()
	expectedResult = []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	arrayCheck(t, "NextPrime()", result, expectedResult)
	ResetPrimenumbers()
}

func TestGetFactors(t *testing.T) {
	var result []int
	var expected []int
	ResetPrimenumbers()
	result = GetFactors(2)
	expected = []int{2}
	arrayCheck(t, "GetFactors(2)", result, expected)
	ResetPrimenumbers()
	result = GetFactors(3)
	expected = []int{3}
	arrayCheck(t, "GetFactors(3)", result, expected)
	ResetPrimenumbers()
	result = GetFactors(4)
	expected = []int{2, 2}
	arrayCheck(t, "GetFactors(4)", result, expected)
	ResetPrimenumbers()
	result = GetFactors(5)
	expected = []int{5}
	arrayCheck(t, "GetFactors(5)", result, expected)
	ResetPrimenumbers()
	result = GetFactors(6)
	expected = []int{2, 3}
	arrayCheck(t, "GetFactors(6)", result, expected)
	ResetPrimenumbers()
	result = GetFactors(7)
	expected = []int{7}
	arrayCheck(t, "GetFactors(7)", result, expected)
	ResetPrimenumbers()
	result = GetFactors(8)
	expected = []int{2, 2, 2}
	arrayCheck(t, "GetFactors(8)", result, expected)
	ResetPrimenumbers()
	result = GetFactors(9)
	expected = []int{3, 3}
	arrayCheck(t, "GetFactors(9)", result, expected)
	ResetPrimenumbers()
	result = GetFactors(10)
	expected = []int{2, 5}
	arrayCheck(t, "GetFactors(10)", result, expected)
	ResetPrimenumbers()
}

func TestGetPrimeArraySlice(t *testing.T) {
	var result []int
	var expected []int
	ResetPrimenumbers()
	result = GetPrimeArraySlice(1)
	expected = []int{2}
	arrayCheck(t, "GetPrimeArraySlice(1)", result, expected)
	ResetPrimenumbers()
	result = GetPrimeArraySlice(6)
	expected = []int{2, 3, 5, 7, 11, 13}
	arrayCheck(t, "GetPrimeArraySlice(6)", result, expected)
	ResetPrimenumbers()
	result = GetPrimeArraySlice(2)
	expected = []int{2, 3}
	arrayCheck(t, "GetPrimeArraySlice(2)", result, expected)
	ResetPrimenumbers()
	result = GetPrimeArraySlice(0)
	expected = []int{}
	arrayCheck(t, "GetPrimeArraySlice(0)", result, expected)
	ResetPrimenumbers()
	result = GetPrimeArraySlice(9)
	expected = []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	arrayCheck(t, "GetPrimeArraySlice(9)", result, expected)
	ResetPrimenumbers()
}

func TestGetPrimeNumbersUpTo(t *testing.T) {
	var result []int
	var expected []int
	ResetPrimenumbers()
	result = GetPrimeNumbersUpTo(2)
	expected = []int{2}
	arrayCheck(t, "GetPrimeNumbersUpTo(2)", result, expected)
	ResetPrimenumbers()
	result = GetPrimeNumbersUpTo(3)
	expected = []int{2, 3}
	arrayCheck(t, "GetPrimeNumbersUpTo(3)", result, expected)
	ResetPrimenumbers()
	result = GetPrimeNumbersUpTo(4)
	expected = []int{2, 3, 5}
	arrayCheck(t, "GetPrimeNumbersUpTo(4)", result, expected)
	ResetPrimenumbers()
	result = GetPrimeNumbersUpTo(5)
	expected = []int{2, 3, 5}
	arrayCheck(t, "GetPrimeNumbersUpTo(5)", result, expected)
	ResetPrimenumbers()
	result = GetPrimeNumbersUpTo(6)
	expected = []int{2, 3, 5, 7}
	arrayCheck(t, "GetPrimeNumbersUpTo(6)", result, expected)
	ResetPrimenumbers()
	result = GetPrimeNumbersUpTo(7)
	expected = []int{2, 3, 5, 7}
	arrayCheck(t, "GetPrimeNumbersUpTo(7)", result, expected)
	ResetPrimenumbers()
}

var nextPrimeResult []int

func benchmarkNextPrime(primeTarget int, b *testing.B) {
	var value []int
	for i := 0; i < b.N; i++ {
		ResetPrimenumbers()
		for j := 0; j < primeTarget; j++ {
			value = NextPrime()
		}
	}
	nextPrimeResult = value
}
func BenchmarkNextPrime100(b *testing.B)   { benchmarkNextPrime(100, b) }
func BenchmarkNextPrime500(b *testing.B)   { benchmarkNextPrime(500, b) }
func BenchmarkNextPrime1000(b *testing.B)  { benchmarkNextPrime(1000, b) }
func BenchmarkNextPrime10000(b *testing.B) { benchmarkNextPrime(10000, b) }

var getFactorsResult []int

func benchmarkGetFactors(primeTarget int, b *testing.B) {
	var value []int
	for i := 0; i < b.N; i++ {
		ResetPrimenumbers()
		value = GetFactors(primeTarget)
	}
	getFactorsResult = value
}
func BenchmarkGetFactors123(b *testing.B)    { benchmarkGetFactors(123, b) }
func BenchmarkGetFactors1234(b *testing.B)   { benchmarkGetFactors(1234, b) }
func BenchmarkGetFactors12345(b *testing.B)  { benchmarkGetFactors(12345, b) }
func BenchmarkGetFactors123456(b *testing.B) { benchmarkGetFactors(123456, b) }

var getPrimeNumbersUpToResult []int

func benchmarkGetPrimeNumbersUpTo(primeTarget int, b *testing.B) {
	var value []int
	for i := 0; i < b.N; i++ {
		ResetPrimenumbers()
		value = GetPrimeNumbersUpTo(primeTarget)
	}
	getPrimeNumbersUpToResult = value
}

func BenchmarkGetPrimeNumbersUpTo100(b *testing.B)    { benchmarkGetPrimeNumbersUpTo(100, b) }
func BenchmarkGetPrimeNumbersUpTo1000(b *testing.B)   { benchmarkGetPrimeNumbersUpTo(1000, b) }
func BenchmarkGetPrimeNumbersUpTo10000(b *testing.B)  { benchmarkGetPrimeNumbersUpTo(10000, b) }
func BenchmarkGetPrimeNumbersUpTo100000(b *testing.B) { benchmarkGetPrimeNumbersUpTo(100000, b) }

// Erwartungshaltung ist das die geforderten Primzahlen enthalten sind
func arrayCheck(t *testing.T, call string, result, expected []int) {
	if !arraytools.EqualInt(result, expected) {
		t.Errorf("%v -- Fehlerhaft\n---Erwartet: %v\n---Resultat: %v\n", call, expected, result)
	}
}
