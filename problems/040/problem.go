package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	champernownes := 1
	prod := 1
	target := 10
	for i := 2; i <= 100_000; i++ {
		num := strconv.Itoa(i)
		if champernownes < target && target < champernownes+len(num) {
			index := target - champernownes
			digit, _ := strconv.Atoi(num[index-1 : index])
			prod *= digit
			target *= 10
			if target == 10_000_000 {
				os.Exit(0)
			}
		}
		champernownes += len(num)
	}
	fmt.Println(prod)
}
