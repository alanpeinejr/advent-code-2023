package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	//part 1
	fmt.Println(sumPatterns(parseLines(readInput()), 0))

	//part 2
	fmt.Println(sumPatterns(parseLines(readInput()), 1))
}

func sumPatterns(patterns [][][]rune, smudges int) int {
	total := 0
	for _, pattern := range patterns {
		subTotal := 0
		subTotal += findPattern(pattern, rowMisMatches, len(pattern), smudges) * 100
		subTotal += findPattern(pattern, columnMisMatches, len(pattern[0]), smudges)

		total += subTotal
	}
	return total
}
func findPattern(patternGroup [][]rune, misMatchCounter MisMatchCounter, directionLength int, tolerance int) int {
	for i:=0; i < directionLength - 1; i++ {
		misMatches := misMatchCounter(i, i + 1, patternGroup)
		if misMatches <= tolerance {
			for j := 1; i + 1 + j < directionLength && i - j >= 0; j++ {
				misMatches += misMatchCounter(i - j, i + 1 + j, patternGroup)
				if misMatches > tolerance {
					break
				}
			}
			if misMatches == tolerance {
				//because its 1 index'd
				return i + 1
			}
		}
	}
	return 0
}

//if i found a way to swap the [a][j]/[j][a] i could do both directions the same
func columnMisMatches(a, b int, patternGroup [][]rune) int {
	misMatches := len(patternGroup)
	for j := 0; j < len(patternGroup); j++ {
		if patternGroup[j][a] == patternGroup[j][b]{
			misMatches -= 1
		}
	}
	return misMatches
}
func rowMisMatches(a, b int, patternGroup [][]rune) int {
	misMatches := len(patternGroup[0])
	for j := 0; j < len(patternGroup[0]); j++ {
		if patternGroup[a][j] == patternGroup[b][j]{
			misMatches -= 1
		}
	}
	return misMatches
}

type MisMatchCounter func(a, b int, patterngroup [][]rune) int

//func misMatches(a, b, directionLength int, patternGroup [][]rune) int {
//
//}
func parseLines(input string) [][][]rune {
	//get note pattern groups
	patternGroups := make([][][]rune, 0)
	for _, patternGroupString := range strings.Split(input, "\n\n") {
		patternRow := make([][]rune, 0)
			for _, rowString := range strings.Split(patternGroupString, "\n") {
				chars := make([]rune, 0)
				for _, char := range rowString {
					chars = append(chars, char)
				}
				patternRow = append(patternRow, chars)
			}
			patternGroups = append(patternGroups, patternRow)
	}
	return patternGroups
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
