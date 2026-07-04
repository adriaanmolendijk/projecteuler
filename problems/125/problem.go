package main

import (
	"fmt"
	"strconv"
)

func main() {
	limit := int64(100_000_000)
	squares := []int{}
	for i := 1; int64(i)*int64(i) <= limit; i++ {
		squares = append(squares, i*i)
	}

	n := len(squares)
	sumPalindromes := map[int]bool{}
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			sum := sumSlice(i, j, squares)
			if sum > int64(limit) {
				break
			} else {
				num := strconv.FormatInt(sum, 10)
				if isPalindrome(num) {
					sumPalindromes[int(sum)] = true
				}
			}
		}
	}
	sum := 0
	for pal := range sumPalindromes {
		sum += pal
	}
	fmt.Println(sum)
}

func sumSlice(i, j int, slice []int) int64 {
	sum := int64(0)
	for k := i; k <= j; k++ {
		sum += int64(slice[k-1])
	}
	return sum
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
