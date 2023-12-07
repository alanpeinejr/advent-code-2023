package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main(){
	//part 1/part 2 if altered as comments describe
	fmt.Println(score(parseLines(readInput())))

}
func score(hands []Hand) int{
	slices.SortFunc(hands, compare)
	score := 0
	for i, v := range hands {
		score+= (i+1) * v.Bid
	}
	return score
}

func compare(a, b Hand) int {
	//a < b negative
	//a >b pos
	aValue := handLevel(a)
	bValue := handLevel(b)
	if aValue < bValue {
		return -1
	} else if aValue > bValue {
		return 1
	}else {
		return highCardCompare(a, b)
	}

}
func handLevel(a Hand) int {
	switch  {
	case fiveAKind(a):
		return 6
	case fourAKind(a):
		return 5
	case fullHouse(a):
		return 4
	case threeAKind(a):
		return 3
	case twoPair(a):
		return 2
	case pair(a):
		return 1
	default:
		return 0

	}
}
func fiveAKind(hand Hand) bool{
	cards := hand.distinctCardOccurrences()
	for _, value := range cards {
		if value == 5 {
			return true
		}
	}
	return false
}
func fourAKind(hand Hand) bool {
	cards := hand.distinctCardOccurrences()
	for _, value := range cards {
		if value == 4 {
			return true
		}
	}
	return false
}
func fullHouse(hand Hand) bool {
	cards:= hand.distinctCardOccurrences()
	//doesn't loop but dont see an easy iterator/first key in go
	for _, v := range cards{
		return len(cards) == 2 && (v == 2 || v == 3)
	}
	return false
}
func threeAKind(hand Hand) bool{
	cards := hand.distinctCardOccurrences()
	for _, v := range cards {
		if v == 3 {
			return true
		}
	}
	return false
}
func twoPair (hand Hand) bool{
	cards := hand.distinctCardOccurrences()
	for _, v := range cards{
		if len(cards) == 3 && v == 2{
			return true
		}
	}
	return false
}
func pair(hand Hand) bool {
	cards := hand.distinctCardOccurrences()
	for _, v := range cards{
		if len(cards) == 4 && v == 2{
			return true
		}
	}
	return false
}
func highCardCompare(a, b Hand) int {
	for i, _ := range a.Cards {
		if CardValues[rune(a.Cards[i])] < CardValues[rune(b.Cards[i])] {
			return -1
		}else if CardValues[rune(a.Cards[i])] > CardValues[rune(b.Cards[i])]{
			return 1
		}
	}
	return 0
}

func (this Hand) Occurs(letter uint8, instances int) bool{
	return strings.Count(this.Cards, string(letter)) == instances
}
func (this Hand) distinctCardOccurrences() map[uint8]int {
	letters := map[uint8]int{}
	for i:=0; i < len(this.Cards); i++ {
		letters[this.Cards[i]]+=1
	}
	//comment out for part 1
	if letters['J'] > 0 {
		return wildCard(letters)
	}

	return letters
}

func wildCard(occurrences map[uint8]int) map[uint8]int {
	for key, _ := range occurrences {
		if key != 'J' {
			occurrences[key] += occurrences['J']

		}
	}
	// j as a different card is more powerful in every instance, except 5 J's
	if len(occurrences) > 1 {
		delete(occurrences, 'J')
	}
	return occurrences
}

func parseLines(input string) []Hand {
	hands := make([]Hand, 0)
	for _, line := range strings.Split(input, "\n") {
		handBidArray := strings.Split(line, " ")
		hands = append(hands, Hand {handBidArray[0], stringToint(handBidArray[1])})
	}
	return hands

}
var CardValues =map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7':7,
	'8':8,
	'9':9,
	'T':10,
	'J':0,//11 if part 1
	'Q':12,
	'K':13,
	'A':14,
}

type (
	Hand struct {
		Cards string
		Bid   int
	}
)
func stringToint(this string) int {
	value, _ := strconv.Atoi(this)
	return int(value)
}

func sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum+= value
	}
	return sum
}

func readInput() string {
	var filename string
	if len(os.Args) < 2 {
		fmt.Println("Assuming local file input.txt")
		filename = "./input.txt"
	}else{
		filename = os.Args[1]
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read file:", filename)
		panic(err)
	}

	//return and account for windows
	return strings.ReplaceAll(string(data), "\r\n", "\n")
}
