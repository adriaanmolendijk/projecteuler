package main

import (
	"fmt"
	"os"
)

func main() {
	side := 100_001
	current := 1
	jump := 0
	sum := int64(1)
	nrPrimes, nrTotal := 0, 1
	for i := 1; i <= (side-1)/2; i++ {
		jump += 2
		for j := 1; j <= 4; j++ {
			current += jump
			sum += int64(current)
			if isPrime(int64(current)) {
				nrPrimes++
			}
			nrTotal++
		}
		ratio := float64(nrPrimes) / float64(nrTotal)
		if ratio < 0.1 {
			fmt.Println(jump + 1)
			os.Exit(0)
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
