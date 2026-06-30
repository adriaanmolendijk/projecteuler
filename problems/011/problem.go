package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	row := 0
	n := 20
	grid := make([][]int, n)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		grid[row] = make([]int, n)
		for j := 0; j < n; j++ {
			grid[row][j], _ = strconv.Atoi(numbers[j])
		}
		row++
	}

	maxProduct := 1
	for i := 1; i <= 16; i++ {
		for j := 1; j <= 16; j++ {
			subGrid := make([][]int, 4)
			for k := 0; k < 4; k++ {
				subGrid[k] = make([]int, 4)
				for l := 0; l < 4; l++ {
					subGrid[k][l] = grid[i+k][j+l]
				}
			}

			prod := maxProduct44Grid(subGrid)
			maxProduct = max(maxProduct, prod)
		}
	}
	fmt.Println(maxProduct)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func maxProduct44Grid(grid [][]int) int {
	maxProduct, prod1, prod2 := 1, 1, 1

	// diagonals
	for i := 1; i <= 4; i++ {
		prod1 *= grid[i-1][i-1]
		prod2 *= grid[i-1][4-i]
	}
	maxProduct = max(maxProduct, prod1)
	maxProduct = max(maxProduct, prod2)

	for i := 1; i <= 4; i++ {
		prodRow, prodCol := 1, 1
		for j := 1; j <= 4; j++ {
			prodRow *= grid[i-1][j-1]
			prodCol *= grid[j-1][i-1]
		}
		maxProduct = max(maxProduct, prodRow)
		maxProduct = max(maxProduct, prodCol)
	}

	return maxProduct
}
