package main

import "fmt"

func main() {
	// Let n = 10k + b
	// Consider b mod 10
	// If b is odd, then n^2 + 1 is even, so n^2 + 1 is not prime
	// if b = 2 or b = 8, then n^2 + 1 is divisible by 5, so n^2 + 1 is not prime
	// if b = 4 or b = 6, then n^2 + 9 is divisible by 5, so n^2 + 9 is not prime
	// This leaves b = 0 and therefore n = 10k

	limit, sum := 1_000_000, 0
	additions := []int{1, 3, 7, 9, 13, 27}
	for i := 1; i*10 <= limit; i++ {
		k := i * 10
		magicInteger := true
		for _, add := range additions {
			m := int64(k)*int64(k) + int64(add)
			if !isPrime(m) {
				magicInteger = false
				break
			}
		}

		if magicInteger {
			sum += k
			fmt.Println(k)
		}
	}
	fmt.Println(sum)
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
