package main

import "fmt"

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
	}
	fmt.Println(sums[n][n] - 1)

}

func dynamicSum(i, j int, sums [][]int64) int64 {
	sum := int64(0)
	for k := 1; k <= j; k++ {
		limit := min(k, i-k)
		sum += sums[i-k][limit]
	}
	return sum
}
