package main

import (
	"fmt"
	"os"
)

var numbers20 map[int]string
var numbersTen map[int]string

func main() {
	numbers20 = map[int]string{}
	numbers20[1] = "one"
	numbers20[2] = "two"
	numbers20[3] = "three"
	numbers20[4] = "four"
	numbers20[5] = "five"
	numbers20[6] = "six"
	numbers20[7] = "seven"
	numbers20[8] = "eight"
	numbers20[9] = "nine"
	numbers20[10] = "ten"
	numbers20[11] = "eleven"
	numbers20[12] = "twelve"
	numbers20[13] = "thirteen"
	numbers20[14] = "fourteen"
	numbers20[15] = "fifteen"
	numbers20[16] = "sixteen"
	numbers20[17] = "seventeen"
	numbers20[18] = "eighteen"
	numbers20[19] = "nineteen"

	numbersTen = map[int]string{}
	numbersTen[2] = "twenty"
	numbersTen[3] = "thirty"
	numbersTen[4] = "forty"
	numbersTen[5] = "fifty"
	numbersTen[6] = "sixty"
	numbersTen[7] = "seventy"
	numbersTen[8] = "eighty"
	numbersTen[9] = "ninety"

	if len(num2String(342)) != 23 {
		os.Exit(1)
	}
	if len(num2String(115)) != 20 {
		os.Exit(1)
	}
	if len(num2String(100)) != len("onehundred") {
		os.Exit(1)
	}

	out := ""
	for i := 1; i <= 1_000; i++ {
		out += num2String(i)
	}
	fmt.Println(len(out))
}

func num2String(num int) string {
	if num == 1_000 {
		return "onethousand"
	}
	if num%100 == 0 {
		return numbers20[num/100] + "hundred"
	}
	if num > 100 {
		return numbers20[num/100] + "hundredand" + num2String(num%100)
	}
	if num >= 20 {
		return numbersTen[num/10] + numbers20[num%10]
	}
	return numbers20[num]
}
