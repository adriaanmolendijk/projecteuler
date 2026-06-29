package main

import "fmt"

func main() {
	sum := int64(2)
	f1 := 1
	f2 := 2

	limit := 4_000_000
	for n := 3; n < limit; n++ {
		fn := f1 + f2
		if fn > limit {
			break
		}
		f1 = f2
		f2 = fn
		if fn%2 != 0 {
			sum += int64(fn)
		}
	}
	fmt.Println(sum)
}
