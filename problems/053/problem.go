package main

import (
	"fmt"
	"math/big"
)

func main() {
	count := 0
	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			fact := factorial(n, r)
			if fact.Cmp(big.NewInt(1_000_000)) == 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func factorial(n int, r int) *big.Int {
	out := big.NewInt(1)
	for i := n - r + 1; i <= n; i++ {
		out.Mul(out, big.NewInt(int64(i)))
	}
	for j := 1; j <= r; j++ {
		out.Div(out, big.NewInt(int64(j)))
	}
	return out
}
