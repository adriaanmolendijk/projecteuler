package main

import (
	"fmt"
	"strconv"
)

func main() {

	sum := 0
	for i := 10; i <= 1_000_000; i++ {
		if isPrime(int64(i)) && isTruncPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)

}

func isTruncPrime(n int) bool {
	num := strconv.Itoa(n)
	for i := 1; i < len(num); i++ {
		subNum, _ := strconv.Atoi(num[i:len(num)])
		if !isPrime(int64(subNum)) {
			return false
		}

		subnum2, _ := strconv.Atoi(num[0 : len(num)-i])
		if !isPrime(int64(subnum2)) {
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
