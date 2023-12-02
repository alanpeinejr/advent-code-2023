package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	games := parseInput(readInput())
	//part 1
	validGameIds := disqualify(Pull{12, 13, 14}, games)
	fmt.Println(sum(validGameIds))
	//part 2
	powers := minimize(games)
	fmt.Println(sum(powers))

}

func sum(validGames []int) int {
	sum := 0
	for _, value := range validGames {
		sum+= value
	}
	return sum
}

func minimize(games [][]Pull) []int {
	powers := make([]int, len(games))
	for i, game := range games {
		min := Pull{0, 0, 0}
		for _, pull := range game {
			if pull.Red > min.Red {
				min.Red = pull.Red
			}
			if pull.Green > min.Green {
				min.Green = pull.Green
			}
			if pull.Blue > min.Blue {
				min.Blue = pull.Blue
			}
		}
		powers[i] = min.Red * min.Green * min.Blue
	}
	return powers
}

//returns an array of valid game ids
func disqualify(totalCubes Pull, games [][]Pull) []int{
	validGames := make([]int, 0)
	for i, game := range games {
		valid:= true
		for _, pull := range game {
			if pull.Red > totalCubes.Red || pull.Green > totalCubes.Green || pull.Blue > totalCubes.Blue {
				valid = false
			}
		}
		if valid {
			validGames = append(validGames, i+1)

		}
	}
	return validGames
}

type (
	Pull struct {
		Red int
		Green int
		Blue int
	}
)
func parseInput(input string) [][]Pull {
	games := make([][]Pull, 0)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		games = append(games, parseLine(line))
	}

	return games
}
func parseLine(input string) []Pull {
	game := make([]Pull, 0)
	_, noGameLabel, _ := strings.Cut(input, ": ")
	pullStrings := strings.Split(noGameLabel, "; ")
	for _, pullString := range pullStrings {
		game = append(game, parsePull(pullString))
	}

	return game
}
func parsePull(input string) Pull{
	red, green, blue := 0,0,0
	for _, cubeAmountString := range strings.Split(input, ",") {
		if strings.Contains(cubeAmountString, " red") {
			red = cutColorLabel(cubeAmountString, " red")
		}
		if strings.Contains(cubeAmountString, " green") {
			green = cutColorLabel(cubeAmountString, " green")
		}
		if strings.Contains(cubeAmountString, " blue") {
			blue = cutColorLabel(cubeAmountString, " blue")
		}
	}
	return Pull{red, green, blue}
}
func cutColorLabel(valueLabel string, label string) int {
	fmt.Println(valueLabel)
	value, _ := strings.CutSuffix(valueLabel, label)
	return stringToint(strings.TrimSpace(value))
}

func stringToint(this string) int {
	value, _ := strconv.Atoi(this)
	return int(value)
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
