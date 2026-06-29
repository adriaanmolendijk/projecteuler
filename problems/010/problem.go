package main

import "fmt"

func main() {
	sum := int64(0)

	for i := 1; i <= 2_000_000; i++ {
		if isPrime(int64(i)) {
			sum += int64(i)
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
