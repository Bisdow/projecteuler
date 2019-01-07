package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Bisdow/projecteuler/amicablenumbers"
	"github.com/Bisdow/projecteuler/circularprimes"
	"github.com/Bisdow/projecteuler/coinsums"
	"github.com/Bisdow/projecteuler/countingsundays"
	"github.com/Bisdow/projecteuler/digitcancellingfractions"
	"github.com/Bisdow/projecteuler/digitfactorials"
	"github.com/Bisdow/projecteuler/digitfifthpowers"
	"github.com/Bisdow/projecteuler/distinctpowers"
	"github.com/Bisdow/projecteuler/doublebasepalindromes"
	"github.com/Bisdow/projecteuler/evenfibonaccinumbers"
	"github.com/Bisdow/projecteuler/factorialdigitsum"
	"github.com/Bisdow/projecteuler/highlydivisibletriangularnumber"
	"github.com/Bisdow/projecteuler/largestpalindromproduct"
	"github.com/Bisdow/projecteuler/largestprimefactor"
	"github.com/Bisdow/projecteuler/largestproductinagrid"
	"github.com/Bisdow/projecteuler/largestproductinaseries"
	"github.com/Bisdow/projecteuler/largesum"
	"github.com/Bisdow/projecteuler/latticepath"
	"github.com/Bisdow/projecteuler/lexicographicpermutations"
	"github.com/Bisdow/projecteuler/longestcollatzsequence"
	"github.com/Bisdow/projecteuler/maximumpathsumi"
	"github.com/Bisdow/projecteuler/maximumpathsumii"
	"github.com/Bisdow/projecteuler/multiplesofxy"
	"github.com/Bisdow/projecteuler/namesscores"
	"github.com/Bisdow/projecteuler/nonabundantsums"
	"github.com/Bisdow/projecteuler/numberlettercounts"
	"github.com/Bisdow/projecteuler/numberspiraldiagonals"
	"github.com/Bisdow/projecteuler/pandigitalproducts"
	"github.com/Bisdow/projecteuler/pentagonnumbers"
	"github.com/Bisdow/projecteuler/powerdigitsum"
	"github.com/Bisdow/projecteuler/quadraticprimes"
	"github.com/Bisdow/projecteuler/reciprocalcycles"
	"github.com/Bisdow/projecteuler/smallestmultiple"
	"github.com/Bisdow/projecteuler/specialpythagoreantriplet"
	"github.com/Bisdow/projecteuler/summationofprimes"
	"github.com/Bisdow/projecteuler/sumsquaredifference"
	"github.com/Bisdow/projecteuler/xdigitfibonaccinumber"
	"github.com/Bisdow/projecteuler/xteprime"
)

type EulerProblem struct {
	number       int
	name         string
	findSolution func()
}

func getEulerProblems() []EulerProblem {
	return []EulerProblem{
		{number: 1, name: "Multiples of 3 and 5", findSolution: multiplesofxy.ExecProjectEulerProblem},
		{number: 2, name: "Even Fibonacci number", findSolution: evenfibonaccinumbers.ExecProjectEulerProblem},
		{number: 2, name: "Even Fibonacci number", findSolution: evenfibonaccinumbers.ExecProjectEulerProblem},
		{number: 3, name: "Largest Prim Factor", findSolution: largestprimefactor.ExecProjectEulerProblem},
		{number: 4, name: "Largest Palindrome Product", findSolution: largestpalindromproduct.ExecProjectEulerProblem},
		{number: 5, name: "Smallest Multiple", findSolution: smallestmultiple.ExecProjectEulerProblem},
		{number: 6, name: "Sum Square Difference", findSolution: sumsquaredifference.ExecProjectEulerProblem},
		{number: 7, name: "1001ste Prime number", findSolution: xteprime.ExecProjectEulerProblem},
		{number: 8, name: "Largest Product in a Series", findSolution: largestproductinaseries.ExecProjectEulerProblem},
		{number: 9, name: "Special Pythagorean triplet", findSolution: specialpythagoreantriplet.ExecProjectEulerProblem},
		{number: 10, name: "Summation of primes", findSolution: summationofprimes.ExecProjectEulerProblem},
		{number: 11, name: "Largest product in a grid", findSolution: largestproductinagrid.ExecProjectEulerProblem},
		{number: 12, name: "Highly divisible triangular number", findSolution: highlydivisibletriangularnumber.ExecProjectEulerProblem},
		{number: 13, name: "Large sum", findSolution: largesum.ExecProjectEulerProblem},
		{number: 14, name: "Longest Collatz Sequence", findSolution: longestcollatzsequence.ExecProjectEulerProblem},
		{number: 15, name: "Lattice paths", findSolution: latticepath.ExecProjectEulerProblem},
		{number: 16, name: "Power digit sum", findSolution: powerdigitsum.ExecProjectEulerProblem},
		{number: 17, name: "Number letter counts", findSolution: numberlettercounts.ExecProjectEulerProblem},
		{number: 18, name: "Maximum path sum I", findSolution: maximumpathsumi.ExecProjectEulerProblem},
		{number: 19, name: "Counting Sundays", findSolution: countingsundays.ExecProjectEulerProblem},
		{number: 20, name: "Factorial digit sum", findSolution: factorialdigitsum.ExecProjectEulerProblem},
		{number: 21, name: "Amicable numbers", findSolution: amicablenumbers.ExecProjectEulerProblem},
		{number: 22, name: "Names scores", findSolution: namesscores.ExecProjectEulerProblem},
		{number: 23, name: "Non-abundant sums", findSolution: nonabundantsums.ExecProjectEulerProblem},
		{number: 24, name: "Lexicographic permutations", findSolution: lexicographicpermutations.ExecProjectEulerProblem},
		{number: 25, name: "1000-digit Fibonacci number", findSolution: xdigitfibonaccinumber.ExecProjectEulerProblem},
		{number: 26, name: "Reciprocal cycles", findSolution: reciprocalcycles.ExecProjectEulerProblem},
		{number: 27, name: "Quadratic primes", findSolution: quadraticprimes.ExecProjectEulerProblem},
		{number: 28, name: "Number spiral diagonals", findSolution: numberspiraldiagonals.ExecProjectEulerProblem},
		{number: 29, name: "Distinct powers", findSolution: distinctpowers.ExecProjectEulerProblem},
		{number: 30, name: "Digit fifth powers", findSolution: digitfifthpowers.ExecProjectEulerProblem},
		{number: 31, name: "Coin sums", findSolution: coinsums.ExecProjectEulerProblem},
		{number: 32, name: "Pandigital products", findSolution: pandigitalproducts.ExecProjectEulerProblem},
		{number: 33, name: "Digital cancelling fractions", findSolution: digitcancellingfractions.ExecProjectEulerProblem},
		{number: 34, name: "Digit factorials", findSolution: digitfactorials.ExecProjectEulerProblem},
		{number: 35, name: "Circular primes", findSolution: circularprimes.ExecProjectEulerProblem},
		{number: 36, name: "Double-base palindromes", findSolution: doublebasepalindromes.ExecProjectEulerProblem},
		{number: 44, name: "Pentagon numbers", findSolution: pentagonnumbers.ExecProjectEulerProblem},
		{number: 67, name: " Maximum path sum II", findSolution: maximumpathsumii.ExecProjectEulerProblem}}
}

func main() {
	var problems = getEulerProblems()
	println("Please choose your Solution by number. Exit programm wit 'exit'")
	for _, problem := range problems {
		fmt.Printf("%3d - %v\n", problem.number, problem.name)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("-----------------------------------")
	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		// convert CRLF to LF
		input = normalizeInput(input)
		checkProgrammExit(input)
		problemnumber, err := strconv.Atoi(input)
		if err != nil {
			println("Bitte geben Sie eine Zahl oder exit ein")
		}
		problem, problemExists := findProblemSolutionMethod(problemnumber, problems)
		if !problemExists {
			fmt.Printf("Es ist keine Lösung für ein Problem mit der Nummer %3v vorhanden\n", problemnumber)
			continue
		} else {
			fmt.Printf("Lösung für das Problem \"%v - %v\":\n", problem.number, problem.name)
			problem.findSolution()
		}
	}
}

func checkProgrammExit(input string) {
	if input == "exit" {
		os.Exit(0)
	}
}

func normalizeInput(input string) string {
	input = strings.Replace(input, "\n", "", -1)
	input = strings.ToLower(input)
	input = strings.Trim(input, " ")
	input = strings.TrimLeft(input, "0")
	return input
}

func findProblemSolutionMethod(problem int, problems []EulerProblem) (solution EulerProblem, found bool) {
	found = false
	for _, checkProblem := range problems {
		if problem == checkProblem.number {
			solution = checkProblem
			found = true
		}
	}
	return solution, found
}
