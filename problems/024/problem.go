package main

import (
	"fmt"
	"os"
	"strings"
)

var perms map[string]bool
var permCounter int

func main() {
	perms = map[string]bool{}
	permCounter = 0
	generatePerms("", "0123456789")
}

func generatePerms(s string, chars string) {
	if chars == "" {
		permCounter++
		perms[s] = true
		if permCounter == 1_000_000 {
			fmt.Println(s)
			os.Exit(0)
		}
	}

	// chars left
	for _, char := range chars {
		generatePerms(s+string(char), strings.ReplaceAll(chars, string(char), ""))
	}
}
