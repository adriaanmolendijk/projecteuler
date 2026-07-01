package main

import (
	"fmt"
	"math/big"
)

func main() {
	count := 0
	for i := 1; i <= 10_000; i++ {
		if isLychrel(i) {
			count++
		}
	}
	fmt.Println(count)
}

func isLychrel(n int) bool {
	m := big.NewInt(int64(n))
	rev := big.NewInt(0)

	for i := 1; i <= 50; i++ {
		rev.SetString(reverse(m.String()), 10)
		m.Add(m, rev)
		if isPalindrome(m.String()) {
			return false
		}
	}

	return true
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return isPalindrome(s[1 : len(s)-1])
}
