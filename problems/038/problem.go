package main

import (
	"fmt"
	"strconv"
)

func main() {

	pandigital := ""
	for i := 1; i <= 100_000; i++ {
		concat := ""
		for j := 1; j <= 6; j++ {
			prod := i * j
			concat += strconv.Itoa(prod)

			if len(concat) > 9 {
				break
			} else if len(concat) == 9 {
				if isPermutation(concat, "123456789") {
					pandigital = concat
				}
			}
		}
	}
	fmt.Println(pandigital)
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
