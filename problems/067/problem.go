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
	n := 100
	grid := make([][]int, n)
	path := make([][]int, n)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		grid[row] = make([]int, n)
		path[row] = make([]int, n)
		for j := 0; j <= row; j++ {
			grid[row][j], _ = strconv.Atoi(numbers[j])
		}
		row++
	}

	path[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				path[i][0] = grid[i][j] + path[i-1][0]
			} else {
				path[i][j] = grid[i][j] + max(path[i-1][j], path[i-1][j-1])
			}
		}
	}

	maxSum := 0
	for i := 0; i < n; i++ {
		maxSum = max(maxSum, path[n-1][i])
	}
	fmt.Println(maxSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
