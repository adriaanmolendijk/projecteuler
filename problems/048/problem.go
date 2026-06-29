package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := big.NewInt(0)

	n := 1_000
	for i := 1; i <= n; i++ {
		entry := big.NewInt(1)
		for j := 1; j <= i; j++ {
			entry.Mul(entry, big.NewInt(int64(i)))
		}
		sum = sum.Add(sum, entry)
	}
	fmt.Println(sum.String()[len(sum.String())-10:])
}
