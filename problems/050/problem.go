package main

import "fmt"

func main() {
	limit := 1_000_000
	primes := []int{}
	for i := 2; i <= int(limit); i++ {
		if isPrime(int64(i)) {
			primes = append(primes, i)
		}
	}

	n := len(primes)
	maxLength, maxI, maxJ := 0, 0, 0
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			sum := sumSlice(i, j, primes)
			if sum > int64(limit) {
				break
			} else {
				if isPrime(sum) {
					if maxLength < j-i {
						maxLength = j - i
						maxI = i
						maxJ = j
					}
				}
			}
		}
	}
	fmt.Println(sumSlice(maxI, maxJ, primes))

}

func sumSlice(i, j int, slice []int) int64 {
	sum := int64(0)
	for k := i; k <= j; k++ {
		sum += int64(slice[k-1])
	}
	return sum
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
