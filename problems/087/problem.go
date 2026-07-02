package main

import "fmt"

func main() {
	n := 50_000_000

	primes2, primes3, primes4 := []int{}, []int{}, []int{}
	for i := 1; i*i <= n; i++ {
		if isPrime(int64(i)) {
			primes2 = append(primes2, i*i)
		}
	}

	for j := 1; j*j*j <= n; j++ {
		if isPrime(int64(j)) {
			primes3 = append(primes3, j*j*j)
		}
	}

	for k := 1; k*k*k*k <= n; k++ {
		if isPrime(int64(k)) {
			primes4 = append(primes4, k*k*k*k)
		}
	}

	sums := map[int]bool{}
	for _, p2 := range primes2 {
		for _, p3 := range primes3 {
			sum := p2 + p3
			if sum > n {
				break
			}
			for _, p4 := range primes4 {
				newSum := sum + p4
				if newSum > n {
					break
				}
				sums[newSum] = true
			}
		}
	}
	fmt.Println(len(sums))

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
