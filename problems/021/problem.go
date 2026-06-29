package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 10_000; i++ {
		divs := sumDivisors(i)
		if divs != i && sumDivisors(divs) == i {
			sum += i
		}
	}
	fmt.Println(sum)
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
