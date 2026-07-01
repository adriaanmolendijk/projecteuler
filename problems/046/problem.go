package main

import (
	"fmt"
	"os"
)

func main() {
	n := int64(10_000)
	primes := map[int64]bool{}

	for i := int64(1); i < n; i++ {
		if isPrime(i) {
			primes[i] = true
		}
	}

	nums := map[int64]bool{}
	for prime, _ := range primes {
		for j := int64(1); j < n; j++ {
			num := prime + 2*j*j
			nums[num] = true
		}
	}

	for i := int64(2); i < n; i++ {
		// odd
		if i%2 == 1 {
			// composite
			if !isPrime(i) {
				// cannot be written as sum of prime and twice a square
				if !nums[i] {
					fmt.Println(i)
					os.Exit(0)
				}
			}
		}
	}

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
