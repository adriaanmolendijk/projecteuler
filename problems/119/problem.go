package main

import (
	"fmt"
	"math/big"
	"sort"
)

func main() {
	n := 75
	powers := []big.Int{}
	for a := 2; a <= n; a++ {
		for b := 2; b <= n; b++ {
			power := pow(a, b)
			sum := sumDigits(power)
			if a == sum {
				powers = append(powers, *power)
			}
		}
	}
	sort.Slice(powers, func(i, j int) bool {
		return powers[i].Cmp(&powers[j]) < 0
	})
	fmt.Println(powers[30-1])
}

func pow(a, b int) *big.Int {
	result := big.NewInt(1)
	for i := 1; i <= b; i++ {
		result.Mul(result, big.NewInt(int64(a)))
	}
	return result
}

func sumDigits(n *big.Int) int {
	sum := 0
	for _, rune := range n.String() {
		digit := rune - '0'
		sum += int(digit)
	}
	return sum
}
