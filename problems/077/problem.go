package main

import (
	"fmt"
	"os"
)

var n = 100

func main() {
	// number of ways i can be written as sum using max j
	// s(i,j) = sum_k=1^j s(i-k, min(i-k,k))
	sums := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		sums[i] = make([]int64, n+1)
	}

	sums[0][0] = 1
	sums[0][1] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			sums[i][j] = dynamicSum(i, j, sums)
		}

		if sums[i][i] > 5_000 {
			fmt.Println(i)
			os.Exit(0)
		}
	}
}

func dynamicSum(i, j int, sums [][]int64) int64 {
	sum := int64(0)
	for k := 1; k <= j; k++ {
		if isPrime(k) {
			limit := min(k, i-k)
			sum += sums[i-k][limit]
		}
	}
	return sum
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := int(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
