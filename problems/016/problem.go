package main

import (
	"fmt"
	"math/big"
)

func main() {
	bignum := big.NewInt(1)
	for i := 1; i <= 1_000; i++ {
		bignum.Mul(bignum, big.NewInt(2))
	}
	fmt.Println(sumDigits(bignum.String()))
}

func sumDigits(n string) int {
	sum := 0
	for _, rune := range n {
		digit := rune - '0'
		sum += int(digit)
	}
	return sum
}
