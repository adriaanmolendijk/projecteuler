package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Find the unique positive integer whose square has the form 1_2_3_4_5_6_7_8_9_0,
	// where each “_” is a single digit.

	// Given n^2 = 1_2_3_4_5_6_7_8_9_0
	// 10 divides n because it ends with a 0
	// n = 10m => n^2 = (10m)^2 = 100m^2 = 1_2_3_4_5_6_7_8_9_0 => m^2 = 1_2_3_4_5_6_7_8_9

	// m = 100_000l+k
	// m^2 = (100_000l+k)^2 = 100_000^2l^2 + 200_000l*m + k^2 = 1_2_3_4_5_6_7_8_9
	// Thus k^2 has to end with 7_8_9 (as this is smaller than 100_00)
	// We gather such values for k in a list
	targets := []int64{}
	for k := 1; k <= 100_000; k++ {
		j := k * k
		if j%10 == 9 && (j/100)%10 == 8 && (j/10_000)%10 == 7 {
			targets = append(targets, int64(k))
		}
	}

	for l := 1; l <= 10_000; l++ {
		for _, target := range targets {
			m := 100_000*int64(l) + target
			j := m * m
			if isConcealedSquare(strconv.FormatInt(j, 10)) {
				fmt.Println(m * 10)
			}
		}
	}
}

func isConcealedSquare(n string) bool {
	if len(n) != 17 {
		return false
	}
	if n[0] != '1' || n[2] != '2' || n[4] != '3' || n[6] != '4' {
		return false
	}
	if n[8] != '5' || n[10] != '6' {
		return false
	}
	return true
}
