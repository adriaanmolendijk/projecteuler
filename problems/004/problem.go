package main

import (
	"fmt"
	"strconv"
)

func main() {

	maxProduct := 0
	for a := 100; a <= 999; a++ {
		for b := a + 1; b <= 999; b++ {
			product := a * b
			if product > maxProduct {
				if isPalindrome(strconv.Itoa(product)) {
					maxProduct = product
				}
			}
		}
	}
	fmt.Println(maxProduct)
}

// isPalindrome recursively validates a string to be palindrome.
func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return isPalindrome(s[1 : len(s)-1])
}
