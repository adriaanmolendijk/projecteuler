package main

import (
	"fmt"
	"math/big"
)

func main() {

	num := big.NewInt(1)
	maxSum := 0
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			num = big.NewInt(1)
			for j := 1; j <= b; j++ {
				num.Mul(num, big.NewInt(int64(a)))
			}
			sum := sumDigits(num.String())
			maxSum = max(maxSum, sum)
		}
	}
	fmt.Println(maxSum)
}

func sumDigits(n string) int {
	sum := 0
	for _, rune := range n {
		digit := rune - '0'
		sum += int(digit)
	}
	return sum
}
