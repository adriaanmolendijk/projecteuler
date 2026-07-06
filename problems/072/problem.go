package main

import "fmt"

func main() {
	sum := int64(0)
	primes := []int{}
	for d := 2; d <= 1_000_000; d++ {
		if isPrime(int64(d)) {
			primes = append(primes, d)
		}
		phiD := phi(int64(d), primes)
		sum += phiD
	}
	fmt.Println(sum)
}

func phi(n int64, primes []int) int64 {
	primeFactorization := getPrimeFactorization(n, primes)
	prod := int64(1)
	for p, k := range primeFactorization {
		for i := 1; i <= k-1; i++ {
			prod *= int64(p)
		}
		prod *= int64(p - 1)
	}
	return prod
}

func getPrimeFactorization(n int64, primes []int) map[int]int {
	if n <= 2 {
		out := map[int]int{int(n): 1}
		return out
	}

	// n >= 1
	factorization := map[int]int{}
	for index, p := range primes {
		for n%int64(p) == 0 {
			factorization[p]++
			n /= int64(p)
			index--
		}

		if int64(p) == n {
			factorization[p]++
		}
	}

	return factorization
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
