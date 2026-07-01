package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxVal, maxLine := float64(0), 0
	lineNr := 0
	for scanner.Scan() {
		lineNr++
		line := scanner.Text()
		numbers := strings.Split(line, ",")
		a, _ := strconv.ParseInt(numbers[0], 10, 64)
		b, _ := strconv.ParseInt(numbers[1], 10, 64)
		val := float64(b) * math.Log10(float64(a))
		if val > maxVal {
			maxVal = val
			maxLine = lineNr
		}
	}
	fmt.Println(maxLine)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
