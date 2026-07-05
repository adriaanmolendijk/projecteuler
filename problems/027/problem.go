package main

import (
	"fmt"
	"os"
)

func main() {
	if quadraticPrimes(1, 41) != 40 {
		os.Exit(1)
	}
	if quadraticPrimes(-79, 1601) != 80 {
		os.Exit(1)
	}

	maxA, maxB, maxQuad := 0, 0, 0
	for a := -1000; a <= 1000; a++ {
		for b := 1; b <= 1000; b++ {
			quadratic := quadraticPrimes(a, b)
			if quadratic > maxQuad {
				maxQuad = quadratic
				maxA = a
				maxB = b
			}
		}
	}
	fmt.Println(maxA * maxB)
}

func quadraticPrimes(a, b int) int {
	for n := 0; n <= 1_000; n++ {
		function := f(a, b, n)
		if !isPrime(int64(function)) {
			return n
		}
	}
	return -1
}

func f(a, b, n int) int {
	return n*n + a*n + b
}

func isPrime(n int64) bool {
	if n < 2 {
		return false
	}
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
