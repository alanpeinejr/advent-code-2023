package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//part 1
	scratchCards := parseInput(readInput())
	fmt.Println(sum(scoreCards(scratchCards)))
	//part 2
	findWins(scratchCards)
	fmt.Println(sum(scams(scratchCards)))

}

func scams(cards []ScratchCard) []int {
	copies := make([]int, len(cards))
	for i, card := range cards {
		copies[i]+=1
		fmt.Println(i, card.Wins, copies[i])
		//for the next amount of wins, add the amount of copies
		for j:= i+1; j <= i + card.Wins && j < len(cards); j++ {
			copies[j]+= copies[i]
		}
	}

	return copies
}

func findWins(cards []ScratchCard) {
	for i, card :=  range cards {
		score := 0
		for _, winningNumber := range card.Winning {
			if contains(winningNumber, card.Numbers) {
				score+=1
			}
		}
		cards[i].Wins = score
	}
}
func scoreCards(cards []ScratchCard) []int {
	scores := make([]int, 0)
	for _, card :=  range cards {
		score := 0
		for _, winningNumber := range card.Winning {
			if contains(winningNumber, card.Numbers) {
				if score == 0 {
					score = 1
				}else {
					score *= 2
				}
			}
		}
		scores = append(scores, score)
	}
	return scores
}
func contains(number int, values []int) bool {
	for _, value := range values {
		if number == value {
			return true
		}
	}
	return false
}
func parseInput(input string) []ScratchCard {
	cards := make([]ScratchCard, 0)
	noDoubles := strings.ReplaceAll(input, "  ", " ")//single every double space
	for i, line := range strings.Split(noDoubles, "\n") {
		values := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " | ")
		winningNumberStringsArray := strings.Split(strings.TrimSpace(values[0]), " ")
		numberStringArray := strings.Split(strings.TrimSpace(values[1]), " ")

		winningnumbers := make([]int, len(winningNumberStringsArray))
		for i, value := range winningNumberStringsArray {
			winningnumbers[i] = stringToint(strings.TrimSpace(value))
		}
		numbers := make([]int, len(numberStringArray))
		for i, value := range numberStringArray {
			numbers[i] = stringToint(strings.TrimSpace(value))
		}
		cards = append(cards, ScratchCard{i,winningnumbers, numbers, 0})
	}

	return cards
}
type (
	ScratchCard struct {
		ID int
		Winning []int
		Numbers []int
		Wins int
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
