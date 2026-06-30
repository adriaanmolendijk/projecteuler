package main

import (
	"fmt"
	"strconv"
)

func main() {
	targets := map[int]bool{}
	for a := 1; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			c := a * b
			if len(strconv.Itoa(c)) > 9 {
				break
			}
			concat := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c)
			if isPermutation(concat, "123456789") {
				targets[c] = true
			}
		}
	}
	sum := 0
	for key, _ := range targets {
		sum += key
	}
	fmt.Println(sum)
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
