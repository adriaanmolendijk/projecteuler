package main

import (
	"fmt"
	"os"
)

func main() {

	primeCount := 0
	for i := 1; i <= 10_000_000; i++ {
		if isPrime(int64(i)) {
			primeCount++
			if primeCount == 10_001 {
				fmt.Println(i)
				os.Exit(0)
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
