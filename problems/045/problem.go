package main

import (
	"fmt"
	"os"
)

func main() {

	n := int64(1_000_000)
	T := map[int64]bool{}
	P := map[int64]bool{}
	H := map[int64]bool{}
	for i := int64(2); i <= n; i++ {
		j := int64(i)
		T[j*(j+1)/2] = true
		P[j*(3*j-1)/2] = true
		H[j*(2*j-1)] = true
	}

	for key, _ := range T {
		if P[key] && H[key] {
			fmt.Println(key)
			os.Exit(0)
		}
	}
}
