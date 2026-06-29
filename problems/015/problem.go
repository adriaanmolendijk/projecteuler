package main

import "fmt"

func main() {

	n := 21
	grid := make([][]int, n)

	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		grid[i][0] = 1
		grid[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	fmt.Println(grid)
}
