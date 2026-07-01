package main

import "fmt"

var target = 200
var coins []int
var count int

func main() {
	count = 0
	coins = []int{1, 2, 5, 10, 20, 50, 100, 200}
	generateMoney(map[int]int{}, 200)
	fmt.Println(count)
}

func generateMoney(spend map[int]int, limit int) {
	total := getTotalSpend(spend)
	if total > target {
		return
	} else if total == target {
		count++
		return
	} else {
		// perform recursion
		for _, coin := range coins {
			if coin > limit {
				continue
			}
			newSpend := map[int]int{}
			for spend, amount := range spend {
				newSpend[spend] = amount
			}
			newSpend[coin]++
			generateMoney(newSpend, min(limit, coin))
		}
	}
}

func getTotalSpend(spend map[int]int) int {
	total := 0
	for coin, amount := range spend {
		total += coin * amount
	}
	return total
}
