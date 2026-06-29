package main

import (
	"fmt"
	"os"
)

func main() {
	maxLength, maxIndex := 0, 0
	for i := 1; i <= 1_000_000; i++ {
		collatzI := collatz(i)
		if collatzI > maxLength {
			maxLength = collatzI
			maxIndex = i
		}
	}
	fmt.Println(maxIndex)
}

func collatz(n int) int {
	if n < 1 {
		fmt.Println("An error occured")
		os.Exit(1)
	} else if n == 1 {
		return 1
	}

	if n%2 == 0 {
		return 1 + collatz(n/2)
	} else {
		return 2 + collatz((3*n+1)/2)
	}
}
