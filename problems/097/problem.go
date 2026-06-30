package main

import (
	"fmt"
	"math/big"
)

func main() {
	num := big.NewInt(1)
	for i := 1; i <= 7830457; i++ {
		num.Mul(num, big.NewInt(2))
		num.Mod(num, big.NewInt(int64(100_000_000_000)))
	}
	num.Mul(num, big.NewInt(28433))
	num.Add(num, big.NewInt(1))
	fmt.Println(num.String()[len(num.String())-10:])
}
