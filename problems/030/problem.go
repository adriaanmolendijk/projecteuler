package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum := 0
	for i := 2; i <= 100_000_000; i++ {
		if i == sumFifthPowerDigits(strconv.Itoa(i)) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func sumFifthPowerDigits(n string) int {
	sum := 0
	fifths := [10]int{0, 1, 32, 243, 1024, 3125, 7776, 16807, 32768, 59049}
	for _, rune := range n {
		digit := rune - '0'
		sum += fifths[int(digit)]
	}

	return sum
}
