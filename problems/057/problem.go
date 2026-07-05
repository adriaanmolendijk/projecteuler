package main

import (
	"fmt"
	"math/big"
)

func main() {
	numerator := big.NewInt(3)
	denominator := big.NewInt(2)
	tmp := new(big.Int)
	count := 0
	for i := 2; i <= 1_000; i++ {
		tmp.Set(denominator)
		denominator.Add(denominator, numerator)
		numerator.Add(denominator, tmp)
		if len(numerator.String()) > len(denominator.String()) {
			count++
		}
	}
	fmt.Println(count)
}
