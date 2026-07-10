package main

import (
	"fmt"
	"os"
)

func main() {
	if cycleLength(7) != 6 {
		os.Exit(1)
	}

	maxD, maxCycle := 0, 0
	for d := 1; d < 1_000; d++ {
		if cycle := cycleLength(d); cycle > maxCycle {
			maxD, maxCycle = d, cycle
		}
	}
	fmt.Println(maxD)
}

func cycleLength(d int) int {
	remainder := map[int]int{}
	a := 1
	t := 0
	for {
		if _, ok := remainder[a]; ok {
			break
		}
		remainder[a] = t
		a = a % d * 10
		t++
	}
	return t - remainder[a]
}
