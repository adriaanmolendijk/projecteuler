package main

import (
	"fmt"
	"strconv"
)

func main() {

	n := 50_000
	primes := map[int]bool{}
	for i := 1; i < n; i++ {
		if isPrime(int64(i)) {
			primes[i] = true
		}
	}

	minRatio := 10.0
	minK := -1
	for p, _ := range primes {
		for r, _ := range primes {
			q := p * r
			if q > 10_000_000 {
				continue
			}
			phiQ := phi(int64(q))
			if isPermutation(strconv.Itoa(q), strconv.Itoa(int(phiQ))) {
				ratio := float64(q) / float64(phiQ)
				if ratio < minRatio {
					minRatio = ratio
					minK = q
				}
			}
		}
	}
	fmt.Println(minK)
}

func phi(n int64) int64 {
	primeFactorization := getPrimeFactorization(n)
	prod := int64(1)
	for p, k := range primeFactorization {
		for i := 1; i <= k-1; i++ {
			prod *= p
		}
		prod *= p - 1
	}
	return prod
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

func isPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// s1, s2 have same length
	s1Count := map[rune]int{}
	s2Count := map[rune]int{}
	for _, rune := range s1 {
		s1Count[rune]++
	}
	for _, rune := range s2 {
		s2Count[rune]++
	}

	for key, val := range s1Count {
		if s2Count[key] != val {
			return false
		}
	}
	return true
}
