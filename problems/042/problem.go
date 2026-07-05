package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	count := 0
	for _, name := range names {
		value := val(name)
		if isTriangle(value) {
			count++
		}
	}
	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func val(s string) int {
	out := 0
	for _, r := range s {
		out += int(r - '@')
	}
	return out
}

func isTriangle(n int) bool {
	for i := 1; i <= n; i++ {
		j := i * (i + 1) / 2
		if j > n {
			return false
		} else if j == n {
			return true
		}
	}
	return false
}
