package main

import (
	"fmt"
	"strconv"
)

func main() {
	count := 0
	for i := 1; i <= 10_000_000; i++ {
		if squareDigitChainEndsAt89(i) {
			count++
		}
	}
	fmt.Println(count)
}

func squareDigitChainEndsAt89(n int) bool {
	if n == 89 {
		return true
	}
	if n == 1 {
		return false
	}
	return squareDigitChainEndsAt89(sumSquareDigits(strconv.Itoa(n)))
}

func sumSquareDigits(n string) int {
	sum := 0
	squares := []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}
	for _, rune := range n {
		digit := rune - '0'
		sum += squares[int(digit)]
	}

	return sum
}
