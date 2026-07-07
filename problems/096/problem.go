package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	Tests()
	Tests2()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	n := 9
	grid, row := make([][]int, n), 0
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if row%10 == 0 {
			row = 1
			continue
		} else {
			grid[row-1] = make([]int, n)
			for col := 1; col <= n; col++ {
				num, _ := strconv.ParseInt(line[col-1:col], 10, 32)
				grid[row-1][col-1] = int(num)
			}
			row++

			if row == 10 {
				sol := solveSudokuRecursive(grid)
				num := 100*sol[0][0] + 10*sol[0][1] + sol[0][2]
				sum += num
			}
		}
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Tests() {
	testSudoku := [][]int{
		{0, 0, 3, 0, 2, 0, 6, 0, 0},
		{9, 0, 0, 3, 0, 5, 0, 0, 1},
		{0, 0, 1, 8, 0, 6, 4, 0, 0},
		{0, 0, 8, 1, 0, 2, 9, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 8},
		{0, 0, 6, 7, 0, 8, 2, 0, 0},
		{0, 0, 2, 6, 0, 9, 5, 0, 0},
		{8, 0, 0, 2, 0, 3, 0, 0, 9},
		{0, 0, 5, 0, 1, 0, 3, 0, 0},
	}

	// map[1:true 4:true 5:true 7:true 8:true 9:true]
	if len(getRowCandidates(1, testSudoku)) != 6 {
		os.Exit(1)
	}

	// map[1:true 2:true 3:true 4:true 5:true 6:true]
	if len(getColCandidates(1, testSudoku)) != 6 {
		os.Exit(1)
	}

	// map[2:true 4:true 5:true 6:true 7:true 8:true]
	if len(getSubCandidates(1, 1, testSudoku)) != 6 {
		os.Exit(1)
	}

	// map[4:true 5:true]
	if len(getRowColSubCandidates(1, 1, testSudoku)) != 2 {
		os.Exit(1)
	}

	//Sudoku solvable without recursion
	//sol := solveSudokuNonRecursive(testSudoku)
}

func Tests2() {
	testSudoku := [][]int{
		{2, 0, 0, 0, 8, 0, 3, 0, 0},
		{0, 6, 0, 0, 7, 0, 0, 8, 4},
		{0, 3, 0, 5, 0, 0, 2, 0, 9},
		{0, 0, 0, 1, 0, 5, 4, 0, 8},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{4, 0, 2, 7, 0, 6, 0, 0, 0},
		{3, 0, 1, 0, 0, 7, 0, 4, 0},
		{7, 2, 0, 0, 4, 0, 0, 6, 0},
		{0, 0, 4, 0, 1, 0, 0, 0, 3},
	}

	// map[1:true 4:true 5:true 6:true 7:true 9:true]
	if len(getRowCandidates(1, testSudoku)) != 6 {
		os.Exit(1)
	}

	// map[1:true 5:true 6:true 8:true 9:true]
	if len(getColCandidates(1, testSudoku)) != 5 {
		os.Exit(1)
	}

	// map[1:true 4:true 5:true 7:true 8:true 9:true]
	if len(getSubCandidates(1, 2, testSudoku)) != 6 {
		os.Exit(1)
	}

	// map[1:true 4:true 5:true 7:true 9:true]
	if len(getRowColSubCandidates(1, 2, testSudoku)) != 5 {
		os.Exit(1)
	}

	//Sudoku solvable with recursion
	//sol := solveSudokuRecursive(testSudoku)
}

func solveSudokuRecursive(grid [][]int) [][]int {
	// Exhaust non-recursive solution
	solveSudokuNonRecursive(grid)

	// Find best guess
	minI, minJ, minCandidates := 10, 10, 10
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if grid[i-1][j-1] == 0 {
				candidates := getRowColSubCandidates(i, j, grid)
				if len(candidates) < minCandidates {
					minCandidates = len(candidates)
					minI, minJ = i, j
				}
			}
		}
	}
	if gridIsFilled(grid) {
		return grid
	}

	candidates := getRowColSubCandidates(minI, minJ, grid)
	for candidate := range candidates {
		duplicate := duplicateGrid(grid)
		duplicate[minI-1][minJ-1] = candidate
		duplicate = solveSudokuRecursive(duplicate)
		if gridIsFilled(duplicate) {
			return duplicate
		}
	}

	return grid
}

func solveSudokuNonRecursive(grid [][]int) [][]int {
	finished := false
	for !finished {
		finished = true
		for i := 1; i <= 9; i++ {
			for j := 1; j <= 9; j++ {
				if grid[i-1][j-1] == 0 {
					candidates := getRowColSubCandidates(i, j, grid)
					if len(candidates) == 1 {
						for candidate := range candidates {
							grid[i-1][j-1] = candidate
							finished = false
						}
					}
				}
			}
		}
	}

	return grid
}

func getRowColSubCandidates(row int, col int, grid [][]int) map[int]bool {
	checkEntryAlreadyFilled(row, col, grid)

	rowCandidates := getRowCandidates(row, grid)
	colCandidates := getColCandidates(col, grid)
	subCandidates := getSubCandidates(row, col, grid)

	candidates := map[int]bool{}
	for i := 1; i <= 9; i++ {
		if rowCandidates[i] && colCandidates[i] && subCandidates[i] {
			candidates[i] = true
		}
	}
	return candidates
}

func getRowCandidates(row int, grid [][]int) map[int]bool {
	candidates := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
		7: true,
		8: true,
		9: true,
	}

	for col := 1; col <= 9; col++ {
		delete(candidates, grid[row-1][col-1])
	}
	return candidates
}

func getColCandidates(col int, grid [][]int) map[int]bool {
	candidates := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
		7: true,
		8: true,
		9: true,
	}
	for row := 1; row <= 9; row++ {
		delete(candidates, grid[row-1][col-1])
	}
	return candidates
}

func getSubCandidates(row int, col int, grid [][]int) map[int]bool {
	checkEntryAlreadyFilled(row, col, grid)

	candidates := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
		7: true,
		8: true,
		9: true,
	}
	subRow := (row - 1) / 3
	subCol := (col - 1) / 3
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			delete(candidates, grid[subRow*3+r][subCol*3+c])
		}
	}
	return candidates
}

func checkEntryAlreadyFilled(row, col int, grid [][]int) {
	if grid[row-1][col-1] != 0 {
		fmt.Println("Row", row, "Col", col, "is already filled with value", grid[row-1][col-1])
		os.Exit(1)
	}
}

func gridIsFilled(grid [][]int) bool {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if grid[i-1][j-1] == 0 {
				return false
			}
		}
	}
	return true
}

func duplicateGrid(grid [][]int) [][]int {
	n := len(grid)
	duplicate := make([][]int, n)
	for i := 1; i <= n; i++ {
		duplicate[i-1] = make([]int, n)
		for j := 1; j <= n; j++ {
			duplicate[i-1][j-1] = grid[i-1][j-1]
		}
	}
	return duplicate
}

func verifySolution(grid [][]int) bool {
	for i := 1; i <= 9; i++ {
		candidates := getRowCandidates(i, grid)
		if len(candidates) > 0 {
			return false
		}

		candidates = getColCandidates(i, grid)
		if len(candidates) > 0 {
			return false
		}
	}
	return true
}
