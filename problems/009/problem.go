package main

import (
	"fmt"
	"os"
)

func main() {

	for a := 1; a < 334; a++ {
		for b := a + 1; b < 500; b++ {
			c := 1000 - b - a
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				os.Exit(0)
			}
		}
	}
}
