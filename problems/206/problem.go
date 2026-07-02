package main

import "fmt"

func main() {
	// Find the unique positive integer whose square has the form 1_2_3_4_5_6_7_8_9_0,
	// where each “_” is a single digit.

	// Given n^2 = 1_2_3_4_5_6_7_8_9_0
	// 10 divides n because it ends with a 0
	// n = 10m => n^2 = (10m)^2 = 100m^2 = 1_2_3_4_5_6_7_8_9_0 => m^2 = 1_2_3_4_5_6_7_8_9
	// Looking at i^2 mod 10, we see m = 10k + 3 or m = 10k + 7, because m^2 ends with a 9.

	for k := 1; k <= 100_000_000; k++ {
		m := 10*k + 3
		fmt.Println(m * m)
		m = 10*k + 7
		fmt.Println(m * m)
	}
}
