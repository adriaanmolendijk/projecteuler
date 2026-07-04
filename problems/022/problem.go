package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	names := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ",")
		for _, word := range words {
			names = append(names, word[1:len(word)-1])
		}
	}
	slices.Sort(names)
	score := int64(0)
	for i := 1; i <= len(names); i++ {
		score += int64(i) * val(names[i-1])
	}
	fmt.Println(score)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func val(s string) int64 {
	out := int64(0)
	for _, r := range s {
		out += int64(r - '@')
	}
	return out
}
