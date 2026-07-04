package main

import "fmt"

func main() {
	limit := 1_500_000
	triangles := map[int]int{}
	count := 0
	for m := 1; m*m <= limit; m++ {
		for n := 1; n <= m; n++ {
			if validPair(n, m) {
				a := m*m - n*n
				b := 2 * m * n
				c := m*m + n*n
				sum := a + b + c
				if sum > limit {
					break
				}
				triangles[sum]++
				for k := 2; k <= limit; k++ {
					sumK := sum * k
					if sumK > limit {
						break
					}
					triangles[sum*k]++
				}
			}
		}
	}

	for i := 1; i <= limit; i++ {
		if triangles[i] == 1 {
			count++
		}
	}
	fmt.Println(count)
}

func validPair(n, m int) bool {
	if n%2 == 1 && m%2 == 1 {
		return false
	}
	primeFac1 := getPrimeFactorization(int64(n))
	primeFac2 := getPrimeFactorization(int64(m))
	for p1, _ := range primeFac1 {
		if primeFac2[p1] {
			return false
		}
	}
	return true
}

func getPrimeFactorization(n int64) map[int64]bool {
	if n <= 2 {
		out := map[int64]bool{n: true}
		return out
	}

	// n >= 1
	primes := map[int64]bool{}
	divisor := int64(2)
	for divisor < n {
		if n%divisor == 0 {
			primes[divisor] = true
			n /= divisor
		} else {
			divisor++
		}

		if divisor == n {
			primes[n] = true
		}
	}

	return primes
}
