package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 1; i <= 1_000_000; i++ {
		primeFac := getPrimeFactorization(int64(i))
		if len(primeFac) == 4 {
			if len(getPrimeFactorization(int64(i+1))) == 4 {
				if len(getPrimeFactorization(int64(i+2))) == 4 {
					if len(getPrimeFactorization(int64(i+3))) == 4 {
						fmt.Println(i)
						os.Exit(0)
					}
				}
			}
		}
	}

}

func getPrimeFactorization(n int64) map[int64]bool {
	if n <= 2 {
		out := map[int64]bool{n: true}
		return out
	}

	// n >= 1
	primes := map[int64]bool{}
	divisor := int64(2)
	for divisor < n {
		if n%divisor == 0 {
			primes[divisor] = true
			n /= divisor
		} else {
			divisor++
		}

		if divisor == n {
			primes[n] = true
		}
	}

	return primes
}
