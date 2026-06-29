package main

import "fmt"

func main() {
	lcm := int64(1)
	lcm *= 2 * 2 * 2 * 2 * 3 * 3 * 5 * 7 * 11 * 13 * 17 * 19
	fmt.Println(lcm)
}
