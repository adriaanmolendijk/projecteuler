package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	primes := []int{}
	for n := 1; n <= 10_000; n++ {
		if isPrime(int64(n)) {
			primes = append(primes, n)
		}
	}

	for _, p := range primes {
		for _, q := range primes {
			if q > p {
				if isPairPrime(p, q) {
					for _, r := range primes {
						if r > q {
							if isPairPrime(p, r) && isPairPrime(q, r) {
								for _, s := range primes {
									if s > r {
										if isPairPrime(p, s) && isPairPrime(q, s) && isPairPrime(r, s) {
											for _, t := range primes {
												if t > s {
													if isPairPrime(p, t) && isPairPrime(q, t) && isPairPrime(r, t) && isPairPrime(s, t) {
														fmt.Println(p + q + r + s + t)
														os.Exit(0)
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func isPrime(n int64) bool {
	if n < 2 {
		return false
	}
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPairPrime(p, q int) bool {
	concat1 := strconv.Itoa(p) + strconv.Itoa(q)
	concat2 := strconv.Itoa(q) + strconv.Itoa(p)
	num1, _ := strconv.ParseInt(concat1, 10, 64)
	num2, _ := strconv.ParseInt(concat2, 10, 64)
	return isPrime(num1) && isPrime(num2)
}
