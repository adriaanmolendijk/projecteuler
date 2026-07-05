package main

import (
	"fmt"
	"strconv"
	"strings"
)

var perms map[string]bool

func main() {
	perms = map[string]bool{}
	generatePerms("", "0123456789")
	sum := int64(0)
	for perm := range perms {
		num, _ := strconv.ParseInt(perm, 10, 64)
		sum += num
	}
	fmt.Println(sum)
}

func generatePerms(s string, chars string) {
	if !divibilityCheck(s) {
		return
	}
	if chars == "" {
		perms[s] = true
	}

	// chars left
	for _, char := range chars {
		generatePerms(s+string(char), strings.ReplaceAll(chars, string(char), ""))
	}
}

func divibilityCheck(s string) bool {
	divisors := []int64{2, 3, 5, 7, 11, 13, 17}
	for i := 1; i <= 7; i++ {
		if len(s) >= i+3 {
			num, _ := strconv.ParseInt(s[i:i+3], 10, 64)
			if num%divisors[i-1] != 0 {
				return false
			}
		}
	}

	return true
}
