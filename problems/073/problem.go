package main

import "fmt"

func main() {
	limit := 12_000
	count := 0
	for d := 2; d <= limit; d++ {
		for n := 1; n <= d; n++ {
			if gcd(d, n) == 1 {
				f := float64(n) / float64(d)
				if f > 1/3.0 && f < 1/2.0 {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
