package main

import (
	"fmt"
	"math/big"
)

func main() {
	digitNums := map[*big.Int]bool{}
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			power := pow(a, b)
			if len(power.String()) == b {
				digitNums[power] = true
			}
		}
	}
	fmt.Println(len(digitNums))
}

func pow(a int, b int) *big.Int {
	power := big.NewInt(1)
	for i := 1; i <= b; i++ {
		power.Mul(power, big.NewInt(int64(a)))
	}
	return power
}
