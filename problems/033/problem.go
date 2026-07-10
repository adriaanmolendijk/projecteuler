package main

import "fmt"

func main() {
	frac, frac2 := 0.0, 0.0
	prodN, prodD := 1, 1
	for d := 10; d <= 99; d++ {
		for n := 1; n <= d; n++ {
			if n%11 == 0 {
				// ignoring all trivial the 11/22 22/33 etc fractions
				continue
			}
			frac = float64(n) / float64(d)
			frac2 = float64(n/10) / float64(d%10)

			if frac == frac2 {
				if n%10 == d/10 {
					prodN *= n
					prodD *= d
				}
			}
		}
	}
	fmt.Println(float64(prodN) / float64(prodD))
}
