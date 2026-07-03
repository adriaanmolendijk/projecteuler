package main

import "fmt"

func main() {
	// Let n = 10k + b
	// Consider b mod 10
	// If b is odd, then n^2 + 1 is even, so n^2 + 1 is not prime
	// if b = 2 or b = 8, then n^2 + 1 is divisible by 5, so n^2 + 1 is not prime
	// if b = 4 or b = 6, then n^2 + 9 is divisible by 5, so n^2 + 9 is not prime
	// This leaves b = 0 and therefore n = 10k

	limit, sum := int64(150_000_000), int64(0)
	for i := int64(1); i*10 < limit; i++ {
		k := i * 10
		if isPrimePattern(k) {
			sum += k
			fmt.Println(k)
		}
	}
	fmt.Println(sum)
}

func isPrimePattern(n int64) bool {
	if n < 2 {
		return false
	}

	scorecard := map[int]bool{}
	for i := 1; i <= 27; i++ {
		scorecard[i] = true
	}

	for i := int64(2); i*i <= n*n+27; i++ {
		for s, _ := range scorecard {
			m := n*n + int64(s)
			if m%i == 0 {
				if s == 1 || s == 3 || s == 7 || s == 9 || s == 13 || s == 27 {
					return false
				}
				delete(scorecard, s)
			}
		}
	}
	delete(scorecard, 1)
	delete(scorecard, 3)
	delete(scorecard, 7)
	delete(scorecard, 9)
	delete(scorecard, 13)
	delete(scorecard, 27)

	return len(scorecard) == 0
}
