package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Rad struct {
	n     int
	value int
}

func CmpRad(a, b Rad) int {
	if a.value == b.value {
		return cmp.Compare(a.n, b.n)
	}
	return cmp.Compare(a.value, b.value)
}

func main() {
	n := 100_000
	primes := []int{}
	rads := []Rad{}
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
		rads = append(rads, Rad{i, rad(i, primes)})
	}
	slices.SortFunc(rads, CmpRad)
	fmt.Println(rads[9999].n)
}

func rad(n int, primes []int) int {
	if n <= 2 {
		return 2
	}

	// n >= 1
	rad := 1
	for index, p := range primes {
		multipliedP := false
		for n%p == 0 {
			if !multipliedP {
				rad *= p
				multipliedP = true
			}
			n /= p
			index--
		}

		if p == n {
			rad *= p
			break
		}
	}

	return rad
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
