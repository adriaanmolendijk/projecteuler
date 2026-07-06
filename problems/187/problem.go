package main

import "fmt"

func main() {
	n := 100_000_000
	primes := SieveOfEratosthenes(n)

	exactly2Primes := map[int64]bool{}
	for _, p := range primes {
		for _, q := range primes {
			prod := int64(p) * int64(q)
			if prod > int64(n) {
				break
			}
			exactly2Primes[prod] = true
		}
	}
	fmt.Println(len(exactly2Primes))
}

// SieveOfEratosthenes returns all prime numbers up to and including n
// Source: https://reintech.io/blog/sieve-of-eratosthenes-in-go-tutorial
func SieveOfEratosthenes(n int) []int {
	if n < 2 {
		return []int{}
	}

	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for multiple := p * p; multiple <= n; multiple += p {
				isPrime[multiple] = false
			}
		}
	}

	primes := make([]int, 0, n/10) // Approximate capacity
	for num := 2; num <= n; num++ {
		if isPrime[num] {
			primes = append(primes, num)
		}
	}
	return primes
}
