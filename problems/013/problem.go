package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	n := big.NewInt(1)
	sum := big.NewInt(1)
	for scanner.Scan() {
		line := scanner.Text()
		n.SetString(line, 10)
		sum = sum.Add(sum, n)
		fmt.Println(n)
	}

	fmt.Println(sum.String()[0:10])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
