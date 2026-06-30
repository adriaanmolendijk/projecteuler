package main

import (
	"fmt"
	"strconv"
)

func main() {

	oneToNine := "123456789"
	maxPandigital := 1
	for i := 1; i <= 10_000_000; i++ {
		num := strconv.Itoa(i)
		n := len(num)
		if isPermutation(num, oneToNine[0:n]) {
			if isPrime(int64(i)) {
				maxPandigital = i
			}
		}
	}
	fmt.Println(maxPandigital)
}

func isPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// s1, s2 have same length
	s1Count := map[rune]int{}
	s2Count := map[rune]int{}
	for _, rune := range s1 {
		s1Count[rune]++
	}
	for _, rune := range s2 {
		s2Count[rune]++
	}

	for key, val := range s1Count {
		if s2Count[key] != val {
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
