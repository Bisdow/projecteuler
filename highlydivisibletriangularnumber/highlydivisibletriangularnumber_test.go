package highlydivisibletriangularnumber

import "testing"

/*
1	->	1: 1										=> 1
2	->	3: 1,3										=> 2
3	->	6: 1,2,3,6									=> 4
4	->	10: 1,2,5,10								=> 4
5	->	15: 1,3,5,15								=> 4
6	->	21: 1,3,7,21								=> 4
7	->	28: 1,2,4,7,14,28							=> 6
8	->	36: 1,2,3,4,6,9,12,18,36					=> 9
9	->	45: 1,3,5,9,15,45							=> 6
10	->	55: 1,5,11,55								=> 4
11	->	66:	1,2,3,6,11,22,33,66						=> 8
12	->	78:	1,2,3,6,13,26,39,78						=> 8
13	->	91:	1,7,13,91								=> 4
14	->	105: 1,3,5,7,15,21,35,105					=> 8
15	->	120: 1,2,3,4,5,6,8,10,12,15,20,24,30,40,60,120	=> 16
*/
func TestAmountOfDivisorsOf(t *testing.T) {
	testAmountOfDivisorOf(t, 1, 1)
	testAmountOfDivisorOf(t, 3, 2)
	testAmountOfDivisorOf(t, 6, 4)
	testAmountOfDivisorOf(t, 10, 4)
	testAmountOfDivisorOf(t, 15, 4)
	testAmountOfDivisorOf(t, 21, 4)
	testAmountOfDivisorOf(t, 28, 6)
	testAmountOfDivisorOf(t, 36, 9)
	testAmountOfDivisorOf(t, 45, 6)
	testAmountOfDivisorOf(t, 55, 4)
	testAmountOfDivisorOf(t, 66, 8)
	testAmountOfDivisorOf(t, 78, 8)
	testAmountOfDivisorOf(t, 91, 4)
	testAmountOfDivisorOf(t, 105, 8)
	testAmountOfDivisorOf(t, 120, 16)
}

var result int

func benchmarkHighlydivisibletriangularnumber(input int, b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		value = Highlydivisibletriangularnumber(input)
	}
	result = value
}

func BenchmarkHighlydivisibletriangularnumber10(b *testing.B) {
	benchmarkHighlydivisibletriangularnumber(10, b)
}
func BenchmarkHighlydivisibletriangularnumber100(b *testing.B) {
	benchmarkHighlydivisibletriangularnumber(100, b)
}
func BenchmarkHighlydivisibletriangularnumber1000(b *testing.B) {
	benchmarkHighlydivisibletriangularnumber(1000, b)
}

func testAmountOfDivisorOf(t *testing.T, input int, expected int) {
	result := amountOfDivisorsOf(input)
	if result != expected {
		t.Error("Input", input, "Expected", expected, "Result", result)
	}
}
