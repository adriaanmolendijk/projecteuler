package main

import (
	"fmt"
	"strconv"
)

func main() {
	for a := 1000; a <= 3340; a++ {
		b := a + 3330
		if isPrime(a) && isPermutation(strconv.Itoa(a), strconv.Itoa(b)) {
			c := b + 3330
			if isPrime(b) && isPermutation(strconv.Itoa(b), strconv.Itoa(c)) {
				if isPrime(c) {
					fmt.Println(strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c))
				}
			}
		}
	}
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

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
