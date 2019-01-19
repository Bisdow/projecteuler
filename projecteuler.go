package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	problem001 "github.com/Bisdow/projecteuler/problem001multiplesofxy"
	"github.com/Bisdow/projecteuler/problem002evenfibonaccinumbers"
	"github.com/Bisdow/projecteuler/problem003largestprimefactor"
	"github.com/Bisdow/projecteuler/problem004largestpalindromproduct"
	"github.com/Bisdow/projecteuler/problem005smallestmultiple"
	"github.com/Bisdow/projecteuler/problem006sumsquaredifference"
	"github.com/Bisdow/projecteuler/problem007largestproductinaseries"
	"github.com/Bisdow/projecteuler/problem007xteprime"
	"github.com/Bisdow/projecteuler/problem009specialpythagoreantriplet"
	"github.com/Bisdow/projecteuler/problem010summationofprimes"
	"github.com/Bisdow/projecteuler/problem011largestproductinagrid"
	"github.com/Bisdow/projecteuler/problem012highlydivisibletriangularnumber"
	"github.com/Bisdow/projecteuler/problem013largesum"
	"github.com/Bisdow/projecteuler/problem014longestcollatzsequence"
	"github.com/Bisdow/projecteuler/problem015latticepath"
	"github.com/Bisdow/projecteuler/problem016powerdigitsum"
	"github.com/Bisdow/projecteuler/problem017numberlettercounts"
	"github.com/Bisdow/projecteuler/problem018maximumpathsumi"
	"github.com/Bisdow/projecteuler/problem019countingsundays"
	"github.com/Bisdow/projecteuler/problem020factorialdigitsum"
	"github.com/Bisdow/projecteuler/problem021amicablenumbers"
	"github.com/Bisdow/projecteuler/problem022namesscores"
	"github.com/Bisdow/projecteuler/problem023nonabundantsums"
	"github.com/Bisdow/projecteuler/problem024lexicographicpermutations"
	"github.com/Bisdow/projecteuler/problem025xdigitfibonaccinumber"
	"github.com/Bisdow/projecteuler/problem026reciprocalcycles"
	"github.com/Bisdow/projecteuler/problem027quadraticprimes"
	"github.com/Bisdow/projecteuler/problem028numberspiraldiagonals"
	"github.com/Bisdow/projecteuler/problem029distinctpowers"
	"github.com/Bisdow/projecteuler/problem030digitfifthpowers"
	"github.com/Bisdow/projecteuler/problem031coinsums"
	"github.com/Bisdow/projecteuler/problem032pandigitalproducts"
	"github.com/Bisdow/projecteuler/problem033digitcancellingfractions"
	"github.com/Bisdow/projecteuler/problem034digitfactorials"
	"github.com/Bisdow/projecteuler/problem035circularprimes"
	"github.com/Bisdow/projecteuler/problem036doublebasepalindromes"
	"github.com/Bisdow/projecteuler/problem044pentagonnumbers"
	"github.com/Bisdow/projecteuler/problem067maximumpathsumii"
)

type EulerProblem struct {
	number       int
	name         string
	findSolution func()
}

func getEulerProblems() []EulerProblem {
	return []EulerProblem{
		{number: 1, name: "Multiples of 3 and 5", findSolution: problem001.ExecProjectEulerProblem},
		{number: 2, name: "Even Fibonacci number", findSolution: problem002evenfibonaccinumbers.ExecProjectEulerProblem},
		{number: 2, name: "Even Fibonacci number", findSolution: problem002evenfibonaccinumbers.ExecProjectEulerProblem},
		{number: 3, name: "Largest Prim Factor", findSolution: problem003largestprimefactor.ExecProjectEulerProblem},
		{number: 4, name: "Largest Palindrome Product", findSolution: problem004largestpalindromproduct.ExecProjectEulerProblem},
		{number: 5, name: "Smallest Multiple", findSolution: problem005smallestmultiple.ExecProjectEulerProblem},
		{number: 6, name: "Sum Square Difference", findSolution: problem006sumsquaredifference.ExecProjectEulerProblem},
		{number: 7, name: "1001ste Prime number", findSolution: problem007xteprime.ExecProjectEulerProblem},
		{number: 8, name: "Largest Product in a Series", findSolution: problem007largestproductinaseries.ExecProjectEulerProblem},
		{number: 9, name: "Special Pythagorean triplet", findSolution: problem009specialpythagoreantriplet.ExecProjectEulerProblem},
		{number: 10, name: "Summation of primes", findSolution: problem010summationofprimes.ExecProjectEulerProblem},
		{number: 11, name: "Largest product in a grid", findSolution: problem011largestproductinagrid.ExecProjectEulerProblem},
		{number: 12, name: "Highly divisible triangular number", findSolution: problem012highlydivisibletriangularnumber.ExecProjectEulerProblem},
		{number: 13, name: "Large sum", findSolution: problem013largesum.ExecProjectEulerProblem},
		{number: 14, name: "Longest Collatz Sequence", findSolution: problem014longestcollatzsequence.ExecProjectEulerProblem},
		{number: 15, name: "Lattice paths", findSolution: problem015latticepath.ExecProjectEulerProblem},
		{number: 16, name: "Power digit sum", findSolution: problem016powerdigitsum.ExecProjectEulerProblem},
		{number: 17, name: "Number letter counts", findSolution: problem017numberlettercounts.ExecProjectEulerProblem},
		{number: 18, name: "Maximum path sum I", findSolution: problem018maximumpathsumi.ExecProjectEulerProblem},
		{number: 19, name: "Counting Sundays", findSolution: problem019countingsundays.ExecProjectEulerProblem},
		{number: 20, name: "Factorial digit sum", findSolution: problem020factorialdigitsum.ExecProjectEulerProblem},
		{number: 21, name: "Amicable numbers", findSolution: problem021amicablenumbers.ExecProjectEulerProblem},
		{number: 22, name: "Names scores", findSolution: problem022namesscores.ExecProjectEulerProblem},
		{number: 23, name: "Non-abundant sums", findSolution: problem023nonabundantsums.ExecProjectEulerProblem},
		{number: 24, name: "Lexicographic permutations", findSolution: problem024lexicographicpermutations.ExecProjectEulerProblem},
		{number: 25, name: "1000-digit Fibonacci number", findSolution: problem025xdigitfibonaccinumber.ExecProjectEulerProblem},
		{number: 26, name: "Reciprocal cycles", findSolution: problem026reciprocalcycles.ExecProjectEulerProblem},
		{number: 27, name: "Quadratic primes", findSolution: problem027quadraticprimes.ExecProjectEulerProblem},
		{number: 28, name: "Number spiral diagonals", findSolution: problem028numberspiraldiagonals.ExecProjectEulerProblem},
		{number: 29, name: "Distinct powers", findSolution: problem029distinctpowers.ExecProjectEulerProblem},
		{number: 30, name: "Digit fifth powers", findSolution: problem030digitfifthpowers.ExecProjectEulerProblem},
		{number: 31, name: "Coin sums", findSolution: problem031coinsums.ExecProjectEulerProblem},
		{number: 32, name: "Pandigital products", findSolution: problem032pandigitalproducts.ExecProjectEulerProblem},
		{number: 33, name: "Digital cancelling fractions", findSolution: problem033digitcancellingfractions.ExecProjectEulerProblem},
		{number: 34, name: "Digit factorials", findSolution: problem034digitfactorials.ExecProjectEulerProblem},
		{number: 35, name: "Circular primes", findSolution: problem035circularprimes.ExecProjectEulerProblem},
		{number: 36, name: "Double-base palindromes", findSolution: problem036doublebasepalindromes.ExecProjectEulerProblem},
		{number: 44, name: "Pentagon numbers", findSolution: problem044pentagonnumbers.ExecProjectEulerProblem},
		{number: 67, name: " Maximum path sum II", findSolution: problem067maximumpathsumii.ExecProjectEulerProblem}}
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
