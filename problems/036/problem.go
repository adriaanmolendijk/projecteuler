package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum := 0
	for i := 1; i <= 1_000_000; i++ {
		if isPalindrome(strconv.Itoa(i)) {
			binary := strconv.FormatInt(int64(i), 2)
			if isPalindrome(binary) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
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
