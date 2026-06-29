package main

import "fmt"

func main() {
	fmt.Println(getPrimeFactors(600851475143))
}

func getPrimeFactors(n int64) []int64 {
	if n <= 2 {
		return []int64{n}
	}

	// n >= 1
	primes := make([]int64, 0)
	divisor := int64(2)
	for divisor < n {
		if n%divisor == 0 {
			primes = append(primes, divisor)
			n /= divisor
		} else {
			divisor++
		}

		if divisor == n {
			primes = append(primes, n)
		}
	}

	return primes
}
