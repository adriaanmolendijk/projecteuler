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

	// Read message into ascii slice
	var message []string
	ascii := []rune{}
	for scanner.Scan() {
		line := scanner.Text()
		message = strings.Split(line, ",")
	}

	for _, m := range message {
		num, _ := strconv.Atoi(m)
		ascii = append(ascii, rune(num))
	}

	// Tests
	testKey := generateKey(ascii, "abc")
	if testKey[0] != 'a' || testKey[1] != 'b' || testKey[2] != 'c' {
		os.Exit(1)
	}
	if len(testKey) != 1455 {
		os.Exit(1)
	}
	if xor(65, 42) != 107 || xor(107, 42) != 65 {
		os.Exit(1)
	}

	// The following code snippet will find 4 keys: [exa,exf,exg,exp], of which exp is translating the ciphertext to plaintext
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, a := range alphabet {
		for _, b := range alphabet {
			for _, c := range alphabet {
				try := string(a) + string(b) + string(c)
				key := generateKey(ascii, try)
				plainText := decryptedMessage(ascii, key)
				countE := countRatioChar(plainText, 'e')
				if countE > 0.1 {
					//fmt.Println(try, countE)
				}
			}
		}
	}

	ans := decryptedMessage(ascii, generateKey(ascii, "exp"))
	sum := 0
	for _, a := range ans {
		sum += int(a)
	}
	fmt.Println(sum)
}

//func printMessage(message []rune) string {
//	out := ""
//	for _, m := range message {
//		out += string(m)
//	}
//	return out
//}

func countRatioChar(message []rune, char rune) float64 {
	count := 0
	n := len(message)
	for i := 0; i < n; i++ {
		if message[i] == char {
			count++
		}
	}
	return float64(count) / float64(n)
}

func decryptedMessage(ascii []rune, key []rune) []rune {
	n := len(ascii)
	if n != len(key) {
		fmt.Print("Length of ascii slice and key slice are not equal.")
		os.Exit(1)
	}

	out := make([]rune, n)
	for i := 0; i < n; i++ {
		out[i] = xor(ascii[i], key[i])
	}
	return out
}

func generateKey(ascii []rune, key string) []rune {
	if len(key) != 3 {
		fmt.Println("Incorrect key input of length not equal to 3.")
		os.Exit(1)
	}

	// len(key) == 3
	n := len(ascii)
	out := make([]rune, n)
	for i := 1; i <= n; i++ {
		index := (i - 1) % 3
		out[i-1] = rune(key[index])
	}
	return out
}

func xor(a, b rune) rune {
	return a ^ b
}
