package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var gons map[string]bool

func main() {
	gons = map[string]bool{}
	for i := 14; i <= 19; i++ {
		generatePerms("", "0123456789", i)
		for gon, _ := range gons {
			sol := gonSolToString(gon)
			if sol != "" {
				fmt.Println(gonSolToString(gon))
			}
		}
		gons = map[string]bool{}
	}
}

func generatePerms(s string, chars string, limit int) {
	if chars == "" {
		if validateGon(s, limit) {
			gons[s] = true
		}
		return
	}

	// chars left
	for _, char := range chars {
		generatePerms(s+string(char), strings.ReplaceAll(chars, string(char), ""), limit)
	}
}

func validateGon(s string, limit int) bool {
	if len(s) != 10 {
		return false
	}
	if getGonRowValue(s, 1, 2, 3) != limit {
		return false
	}
	if getGonRowValue(s, 4, 3, 5) != limit {
		return false
	}
	if getGonRowValue(s, 6, 5, 7) != limit {
		return false
	}
	if getGonRowValue(s, 8, 7, 9) != limit {
		return false
	}
	if getGonRowValue(s, 10, 9, 2) != limit {
		return false
	}

	return true
}

func getGonRowValue(s string, node1, node2, node3 int) int {
	return getGonNodeValue(s, node1) + getGonNodeValue(s, node2) + getGonNodeValue(s, node3)
}

func getGonNodeValue(s string, gonPos int) int {
	run := s[gonPos-1]
	if run != '0' {
		return int(run - '0')
	}
	// run == '0' representing 10
	return 10
}

func gonSolToString(s string) string {
	l := len(s)
	if l != 10 {
		os.Exit(1)
	}

	out := ""
	out += strconv.Itoa(getGonNodeValue(s, 1))
	out += strconv.Itoa(getGonNodeValue(s, 2))
	out += strconv.Itoa(getGonNodeValue(s, 3))

	out += strconv.Itoa(getGonNodeValue(s, 4))
	out += strconv.Itoa(getGonNodeValue(s, 3))
	out += strconv.Itoa(getGonNodeValue(s, 5))

	out += strconv.Itoa(getGonNodeValue(s, 6))
	out += strconv.Itoa(getGonNodeValue(s, 5))
	out += strconv.Itoa(getGonNodeValue(s, 7))

	out += strconv.Itoa(getGonNodeValue(s, 8))
	out += strconv.Itoa(getGonNodeValue(s, 7))
	out += strconv.Itoa(getGonNodeValue(s, 9))

	out += strconv.Itoa(getGonNodeValue(s, 10))
	out += strconv.Itoa(getGonNodeValue(s, 9))
	out += strconv.Itoa(getGonNodeValue(s, 2))

	// Get rid of symmetric solutions
	if getGonNodeValue(s, 4) < getGonNodeValue(s, 1) {
		return ""
	}
	if getGonNodeValue(s, 6) < getGonNodeValue(s, 1) {
		return ""
	}
	if getGonNodeValue(s, 8) < getGonNodeValue(s, 1) {
		return ""
	}
	if getGonNodeValue(s, 10) < getGonNodeValue(s, 1) {
		return ""
	}

	// Get rid of 17-digit string solutions
	if len(out) > 16 {
		return ""
	}
	return out
}
