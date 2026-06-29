package main

import "fmt"

func main() {
	sum1 := int64(0)
	sum2 := int64(0)

	n := 100
	for i := 1; i <= n; i++ {
		sum1 += int64(i) * int64(i)
		sum2 += int64(i)
	}
	sum2 *= sum2

	fmt.Println(sum2 - sum1)
}
