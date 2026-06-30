package main

import "fmt"

func main() {

	abundants := map[int]bool{}
	all := map[int]bool{}
	for i := 1; i <= 28123; i++ {
		all[i] = true
		if isAbundant(i) {
			abundants[i] = true
		}
	}

	for a, _ := range abundants {
		for a2, _ := range abundants {
			all[a+a2] = false
		}
	}

	sum := 0
	for a, _ := range all {
		if all[a] == true {
			sum += a
		}
	}
	fmt.Println(sum)
}

func isAbundant(n int) bool {
	return sumDivisors(n) > n
}

func sumDivisors(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
