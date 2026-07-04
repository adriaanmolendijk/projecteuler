package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Card struct {
	value int
	suit  string
}

func main() {

	unitTests()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hand1 := make([]Card, 5)
	hand2 := make([]Card, 5)
	player1Wins := 0
	for scanner.Scan() {
		line := scanner.Text()
		cards := strings.Split(line, " ")
		for i := 1; i <= 5; i++ {
			hand1[i-1] = Card{getCardValue(rune(cards[i-1][0])), cards[i-1][1:2]}
			hand2[i-1] = Card{getCardValue(rune(cards[i+4][0])), cards[i+4][1:2]}
		}
		res := winningHand(hand1, hand2)
		if res == 1 {
			player1Wins++
		}
	}
	fmt.Println(player1Wins)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getCardValue(val rune) int {
	switch val {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(val - '0')
	}
}

func winningHand(hand1 []Card, hand2 []Card) int {
	// Royal Flush
	if royalFlush(hand1) > royalFlush(hand2) {
		return 1
	}
	if royalFlush(hand2) > royalFlush(hand1) {
		return 2
	}

	// Straight Flush
	if straightFlush(hand1) > straightFlush(hand2) {
		return 1
	}
	if straightFlush(hand2) > straightFlush(hand1) {
		return 2
	}

	// 40Ak
	if fourOfAKind(hand1) > fourOfAKind(hand2) {
		return 1
	}
	if fourOfAKind(hand2) > fourOfAKind(hand1) {
		return 2
	}

	// Full House
	if fullHouse(hand1) > fullHouse(hand2) {
		return 1
	}
	if fullHouse(hand2) > fullHouse(hand1) {
		return 2
	}

	// Flush
	if flush(hand1) > flush(hand2) {
		return 1
	}
	if flush(hand2) > flush(hand1) {
		return 2
	}

	// Straight
	if straight(hand1) > straight(hand2) {
		return 1
	}
	if straight(hand2) > straight(hand1) {
		return 2
	}

	// 3OAK
	if threeOfAKind(hand1) > threeOfAKind(hand2) {
		return 1
	}
	if threeOfAKind(hand2) > threeOfAKind(hand1) {
		return 2
	}

	// Two Pair
	if twoPair(hand1) > twoPair(hand2) {
		return 1
	}
	if twoPair(hand2) > twoPair(hand1) {
		return 2
	}

	// Pair
	if pair(hand1) > pair(hand2) {
		return 1
	}
	if pair(hand2) > pair(hand1) {
		return 2
	}

	return decideOnHighCard(hand1, hand2)
}

func decideOnHighCard(hand1 []Card, hand2 []Card) int {
	cards1, cards2 := map[int]int{}, map[int]int{}
	for _, card := range hand1 {
		cards1[card.value]++
	}
	for _, card := range hand2 {
		cards2[card.value]++
	}

	for i := 14; i >= 1; i-- {
		if cards1[i] > cards2[i] {
			return 1
		}
		if cards1[i] < cards2[i] {
			return 2
		}
	}
	os.Exit(1)
	return 1
}

func pair(hand []Card) int {
	counts := map[int]int{}
	for _, card := range hand {
		counts[card.value]++
	}
	for val, count := range counts {
		if count == 2 {
			return val
		}
	}

	return 0
}

func twoPair(hand []Card) int {
	counts := map[int]int{}
	for _, card := range hand {
		counts[card.value]++
	}
	foundDouble := false
	for _, count := range counts {
		if count == 2 {
			foundDouble = true
		}
	}

	// There is a 3OAK and a pair
	if foundDouble && len(counts) == 3 {
		return 1
	}

	return 0
}

func fullHouse(hand []Card) int {
	counts := map[int]int{}
	for _, card := range hand {
		counts[card.value]++
	}
	foundTriple := false
	for _, count := range counts {
		if count == 3 {
			foundTriple = true
		}
	}

	// There is a 3OAK and a pair
	if foundTriple && len(counts) == 2 {
		return 1
	}

	return 0
}

func threeOfAKind(hand []Card) int {
	counts := map[int]int{}
	for _, card := range hand {
		counts[card.value]++
	}
	for val, count := range counts {
		if count == 3 {
			return val
		}
	}
	return 0
}

func fourOfAKind(hand []Card) int {
	counts := map[int]int{}
	for _, card := range hand {
		counts[card.value]++
	}
	for val, count := range counts {
		if count == 4 {
			return val
		}
	}
	return 0
}

func flush(hand []Card) int {
	for i := 1; i < len(hand); i++ {
		if hand[i].suit != hand[0].suit {
			return 0
		}
	}

	maxVal := 0
	for _, card := range hand {
		maxVal = max(maxVal, card.value)
	}
	return maxVal
}

func straight(hand []Card) int {
	minVal, maxVal := 1<<31, 0
	cards := map[int]bool{}
	for _, card := range hand {
		minVal = min(minVal, card.value)
		maxVal = max(maxVal, card.value)
		cards[card.value] = true
	}
	if cards[minVal] && cards[minVal+1] && cards[minVal+2] && cards[minVal+3] && cards[minVal+4] {
		return maxVal
	}

	// exception for A 2 3 4 5
	if maxVal == 14 {
		if cards[2] && cards[3] && cards[4] && cards[5] {
			return 5
		}
	}

	return 0
}

func straightFlush(hand []Card) int {
	if flush(hand) > 0 && straight(hand) > 0 {
		return straight(hand)
	}
	return 0
}

func royalFlush(hand []Card) int {
	if straightFlush(hand) == 14 {
		return 14
	}
	return 0
}

func unitTests() {
	isFlush()
	isStraight()
	isStraightFlush()
	isRoyalFlush()
	isFourOfAKind()
	isThreeOfAKind()
	isFullHouse()
	isTwoPair()
	isPair()
}

func isFullHouse() {
	hand := []Card{
		Card{10, "D"},
		Card{10, "C"},
		Card{10, "H"},
		Card{11, "S"},
		Card{11, "D"},
	}

	if fullHouse(hand) != 1 {
		os.Exit(1)
	}

	hand2 := []Card{
		Card{9, "D"},
		Card{10, "C"},
		Card{10, "H"},
		Card{11, "S"},
		Card{11, "D"},
	}

	if fullHouse(hand2) > 0 {
		os.Exit(1)
	}
}

func isPair() {
	hand := []Card{
		Card{10, "D"},
		Card{10, "C"},
		Card{7, "H"},
		Card{8, "S"},
		Card{9, "D"},
	}

	if pair(hand) != 10 {
		os.Exit(1)
	}
}

func isTwoPair() {
	hand := []Card{
		Card{10, "D"},
		Card{10, "C"},
		Card{9, "H"},
		Card{11, "S"},
		Card{11, "D"},
	}

	if twoPair(hand) != 1 {
		os.Exit(1)
	}
}

func isFourOfAKind() {
	hand := []Card{
		Card{10, "D"},
		Card{10, "C"},
		Card{10, "H"},
		Card{10, "S"},
		Card{11, "D"},
	}

	if fourOfAKind(hand) != 10 {
		os.Exit(1)
	}
}

func isThreeOfAKind() {
	hand := []Card{
		Card{10, "D"},
		Card{10, "C"},
		Card{10, "H"},
		Card{9, "S"},
		Card{8, "D"},
	}

	if threeOfAKind(hand) != 10 {
		os.Exit(1)
	}
}

func isRoyalFlush() {
	hand := []Card{
		Card{10, "D"},
		Card{11, "D"},
		Card{12, "D"},
		Card{13, "D"},
		Card{14, "D"},
	}

	if royalFlush(hand) != 14 {
		os.Exit(1)
	}

	hand2 := []Card{
		Card{2, "D"},
		Card{3, "D"},
		Card{4, "D"},
		Card{5, "D"},
		Card{6, "D"},
	}

	if royalFlush(hand2) > 0 {
		os.Exit(1)
	}
}

func isStraightFlush() {
	hand := []Card{
		Card{2, "D"},
		Card{3, "D"},
		Card{4, "D"},
		Card{5, "D"},
		Card{6, "D"},
	}

	if straightFlush(hand) != 6 {
		os.Exit(1)
	}
}

func isStraight() {
	hand := []Card{
		Card{2, "D"},
		Card{3, "D"},
		Card{4, "D"},
		Card{5, "D"},
		Card{6, "D"},
	}

	if straight(hand) != 6 {
		os.Exit(1)
	}

	hand2 := []Card{
		Card{2, "D"},
		Card{3, "D"},
		Card{4, "D"},
		Card{5, "D"},
		Card{14, "D"},
	}

	if straight(hand2) != 5 {
		os.Exit(1)
	}

	hand3 := []Card{
		Card{2, "D"},
		Card{2, "S"},
		Card{4, "D"},
		Card{5, "D"},
		Card{6, "D"},
	}

	if straight(hand3) > 0 {
		os.Exit(1)
	}
}

func isFlush() {
	hand := []Card{
		Card{2, "D"},
		Card{3, "D"},
		Card{4, "D"},
		Card{5, "D"},
		Card{6, "D"},
	}

	if flush(hand) != 6 {
		os.Exit(1)
	}

	hand2 := []Card{
		Card{2, "D"},
		Card{3, "D"},
		Card{4, "D"},
		Card{5, "D"},
		Card{6, "C"},
	}

	if flush(hand2) > 0 {
		os.Exit(1)
	}
}
