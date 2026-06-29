package main

import (
	"fmt"
	"math/big"
)

func main() {
	fact := big.NewInt(1)
	for i := 1; i <= 100; i++ {
		fact = fact.Mul(fact, big.NewInt(int64(i)))
	}
	fmt.Println(sumDigits(fact.String()))
}

func sumDigits(n string) int {
	sum := 0
	for _, rune := range n {
		digit := rune - '0'
		sum += int(digit)
	}
	return sum
}
