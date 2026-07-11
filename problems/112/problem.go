package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	count := 0
	ratio := 0.0
	for i := 1; i <= 10_000_000; i++ {
		if isBouncy(strconv.Itoa(i)) {
			count++
			ratio = float64(count) / float64(i)
			if ratio >= 0.99 {
				fmt.Println(i)
				os.Exit(0)
			}
		}
	}
}

func isBouncy(num string) bool {
	return !isIncreasing(num) && !isDecreasing(num)
}

func isIncreasing(num string) bool {
	digit := 0
	for _, r := range num {
		newDigit := int(r - '0')
		if newDigit < digit {
			return false
		}
		digit = newDigit
	}
	return true
}

func isDecreasing(num string) bool {
	return isIncreasing(Reverse(num))
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
