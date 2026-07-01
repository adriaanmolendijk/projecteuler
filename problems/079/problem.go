package main

import (
	"bufio"
	"fmt"
	"log"
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

	numbers := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.ParseInt(line, 10, 64)
		numbers = append(numbers, int(num))
	}

	for i := 1; i <= 100_000_000; i++ {
		num := strconv.Itoa(i)
		validGuess := true
		for _, number := range numbers {
			if !validateKey(strconv.Itoa(number), num) {
				validGuess = false
				break
			}
		}

		if validGuess {
			fmt.Println(i)
			os.Exit(0)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func validateKey(key string, guess string) bool {
	if key == "" {
		return true
	}

	r := key[0]
	index := strings.IndexRune(guess, rune(r))
	if index == -1 {
		return false
	}
	return validateKey(key[1:], guess[index:])
}
