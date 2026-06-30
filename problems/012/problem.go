package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 1; i <= 1_000_000; i++ {
		triangle := int64(i) * int64(i+1) / 2
		numDivs := getNumDivisors(triangle)
		if numDivs > 500 {
			fmt.Println(triangle, ": ", numDivs)
			os.Exit(0)
		}
	}

}

func getNumDivisors(n int64) int {
	factorization := getPrimeFactorization(n)
	numDivs := 1
	for _, power := range factorization {
		numDivs *= (power + 1)
	}
	return numDivs
}

func getPrimeFactorization(n int64) map[int64]int {
	if n <= 2 {
		out := map[int64]int{n: 1}
		return out
	}

	// n >= 1
	primes := map[int64]int{}
	divisor := int64(2)
	for divisor < n {
		if n%divisor == 0 {
			primes[divisor]++
			n /= divisor
		} else {
			divisor++
		}

		if divisor == n {
			primes[n]++
		}
	}

	return primes
}
