package main

import "fmt"

func main() {
	n := 1_000_000
	maxRatio, maxN := 1.0, 1
	// jumps of 210 to ensure prime factorization contains 2,3,5,7
	for i := 210; i <= n; i += 210 {
		ratio := float64(i) / float64(phi(int64(i)))
		if ratio > maxRatio {
			maxRatio = ratio
			maxN = i
		}
	}
	fmt.Println(maxN)
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
