package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i := 1_000; i <= 10_000_000; i++ {
		for mult := 2; mult <= 6; mult++ {
			j := i * mult
			if !isPermutation(strconv.Itoa(i), strconv.Itoa(j)) {
				break
			}

			if mult == 6 {
				fmt.Println(i, i*2, i*3, i*4, i*5, i*6)
				os.Exit(0)
			}
		}
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
