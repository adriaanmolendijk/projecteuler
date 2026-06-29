package main

import (
	"fmt"
	"math/big"
)

func main() {

	numbers := map[string]bool{}
	n := 100
	for a := 2; a <= n; a++ {
		bigval := big.NewInt(int64(a))
		for b := 2; b <= n; b++ {
			bigval.Mul(bigval, big.NewInt(int64(a)))
			numbers[bigval.String()] = true
		}
	}

	fmt.Println(len(numbers))
}
