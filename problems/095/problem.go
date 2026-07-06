package main

import (
	"fmt"
	"os"
)

func main() {
	n := 1_000_000

	sumDivs := map[int]int{}
	for i := 1; i <= n; i++ {
		sumDivs[i] = sumDivisors(i)
	}

	// Tests
	if amicableChain(220, sumDivs) != 2 {
		os.Exit(1)
	}
	if amicableChain(284, sumDivs) != 2 {
		os.Exit(1)
	}
	if amicableChain(12496, sumDivs) != 5 {
		os.Exit(1)
	}

	maxAmicable, maxChain := 0, 0
	for i := 1; i <= n; i++ {
		amicable := amicableChain(i, sumDivs)
		if amicable > maxChain {
			maxChain = amicable
			maxAmicable = i
		}

	}
	fmt.Println(maxAmicable)
}

func amicableChain(n int, sumDivs map[int]int) int {
	chain := map[int]int{}
	chain[n] = 1
	next := sumDivs[n]

	for i := 2; i <= 100; i++ {
		if next > 1_000_000 {
			return -1
		}
		if chain[next] > 0 {
			if chain[next] == 1 {
				return i - 1
			} else {
				return -1
			}
		}
		chain[next] = i
		next = sumDivs[next]
	}

	return -1
}

func sumDivisors(n int) int {
	sum := 1
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			sum += i
			d := n / i
			if d != i {
				sum += d
			}
		}
	}
	return sum
}
