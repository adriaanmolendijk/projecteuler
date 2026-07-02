package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := 10_000
	cubes := map[int]int64{}
	for i := 1; i <= n; i++ {
		cubes[i] = int64(i) * int64(i) * int64(i)
	}

	counts := make([]int, n)
	for i := 1; i <= n; i++ {
		cube := cubes[i]
		count := 0
		for j := 1; j <= n; j++ {
			cube2 := cubes[j]
			num := strconv.FormatInt(cube, 10)
			num2 := strconv.FormatInt(cube2, 10)
			if isPermutation(num, num2) {
				count++
				if count == 5 {
					fmt.Println(cubes[i])
					os.Exit(0)
				}
			}
		}
		counts[i-1] = count
	}
}

func isPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// s1, s2 have same length
	s1Count := map[rune]int{}
	s2Count := map[rune]int{}
	for _, rune := range s1 {
		s1Count[rune]++
	}
	for _, rune := range s2 {
		s2Count[rune]++
	}

	for key, val := range s1Count {
		if s2Count[key] != val {
			return false
		}
	}
	return true
}
