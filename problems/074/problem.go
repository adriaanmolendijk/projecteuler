package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Tests
	if digitFactorialChain(69) != 5 {
		os.Exit(1)
	}
	if digitFactorialChain(78) != 4 {
		os.Exit(1)
	}
	if digitFactorialChain(540) != 2 {
		os.Exit(1)
	}

	count := 0
	for i := 1; i <= 1_000_000; i++ {
		if digitFactorialChain(i) == 60 {
			count++
		}
	}
	fmt.Println(count)
}

func digitFactorialChain(n int) int {
	chain := map[int]bool{}
	chain[n] = true
	next := sumFactorialDigits(strconv.Itoa(n))
	for !chain[next] {
		chain[next] = true
		next = sumFactorialDigits(strconv.Itoa(next))
	}
	return len(chain)
}

func sumFactorialDigits(n string) int {
	sum := 0
	factorials := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	for _, rune := range n {
		digit := rune - '0'
		sum += factorials[int(digit)]
	}

	return sum
}
