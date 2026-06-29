package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum := 0
	for i := 3; i <= 10_000_000; i++ {
		out := sumFactorialDigits(strconv.Itoa(i))
		if i == out {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println(sum)
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
