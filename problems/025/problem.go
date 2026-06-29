package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {

	f1 := big.NewInt(1)
	f2 := big.NewInt(2)
	fn := big.NewInt(0)
	n := 3

	for n > 0 {
		fn.Add(f2, f1)

		f1.Add(f2, big.NewInt(0))
		f2.Add(fn, big.NewInt(0))
		n++

		if len(fn.String()) == 1_000 {
			fmt.Println(n)
			os.Exit(0)
		}
	}
}
