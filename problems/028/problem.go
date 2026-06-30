package main

import "fmt"

func main() {
	side := 1001
	current := 1
	jump := 0
	sum := int64(1)
	for i := 1; i <= (side-1)/2; i++ {
		jump += 2
		for j := 1; j <= 4; j++ {
			current += jump
			sum += int64(current)
		}
	}
	fmt.Println(sum)
}
