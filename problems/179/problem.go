package main

import "fmt"

func main() {
	n := 10_000_000
	primes := SieveOfEratosthenes(n)

	divs := map[int]int{}
	for i := 2; i < n; i++ {
		divisors := getNumDivisors(int64(i), primes)
		divs[i] = divisors
	}

	count := 0
	for i := 2; i < n-1; i++ {
		if divs[i] == divs[i+1] {
			count++
		}
	}
	fmt.Println(count)
}

func getNumDivisors(n int64, primes []int) int {
	factorization := getPrimeFactorization(n, primes)
	numDivs := 1
	for _, power := range factorization {
		numDivs *= (power + 1)
	}
	return numDivs
}

func getPrimeFactorization(n int64, primes []int) map[int64]int {
	if n <= 2 {
		out := map[int64]int{n: 1}
		return out
	}

	// n >= 1
	factorization := map[int64]int{}
	for _, prime := range primes {
		if int64(prime) > n {
			break
		}
		for n%int64(prime) == 0 {
			factorization[int64(prime)]++
			n /= int64(prime)
		}
		if n == 1 {
			break
		}
	}

	return factorization
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
