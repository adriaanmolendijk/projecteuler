package main

import (
	"fmt"
	"strconv"
)

func main() {

	count := 0
	for i := 1; i <= 1_000_000; i++ {
		if isPrime(int64(i)) && isCircularPrime(i) {
			count++
		}
	}
	fmt.Println(count)
}

func isCircularPrime(n int) bool {
	num := strconv.Itoa(n)
	for i := 1; i < len(num); i++ {
		circNum, _ := strconv.ParseInt(num[i:]+num[0:i], 10, 64)
		if !isPrime(circNum) {
			return false
		}
	}

	return true
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
