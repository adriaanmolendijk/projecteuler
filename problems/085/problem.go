package main

import "fmt"

func main() {
	width2, height2, closests := int64(0), int64(0), int64(2_000_000)
	rectangles := int64(0)
	for width := 1; width <= 100; width++ {
		for height := 1; height <= 100; height++ {
			rectangles = getRectangles(width, height)
			if max(rectangles-2_000_000, 2_000_000-rectangles) < closests {
				closests = max(rectangles-2_000_000, 2_000_000-rectangles)
				width2 = int64(width)
				height2 = int64(height)
			}
		}
	}
	fmt.Println(width2 * height2)
}

func getRectangles(width, height int) int64 {
	sum := int64(0)
	for i := 1; i <= height; i++ {
		for j := 1; j <= width; j++ {
			sum += int64(width-j+1) * int64(height-i+1)
		}
	}
	return sum
}
