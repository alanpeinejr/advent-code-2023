package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	//part 1
	fmt.Println(sumWeights(roll(parseLines(readInput()), North)))

	//part 2
	platform := parseLines(readInput())
	scores := map[int]int{}
	var cycle, cycleStart int
	for i := 0; i < 200; i++{
		score := uniqueScore(spin(platform))
		if scores[score]!= 0{
			cycle = i - scores[score]
			cycleStart = scores[score]
			break;
		}
		scores[score] = i
	}
	fmt.Println(cycle, cycleStart)
	//get to next point in cycle that matches, easier than reloading input
	//god it took me forever to realize I needed to add 1 because I zero indexed
	for i:= 0; i < (1000000000-(cycleStart + cycle + 1)) % cycle; i++{
		spin(platform)
	}
	fmt.Println(sumWeights(platform))

}
func uniqueScore(platform [][]rune) int {
	score := 0
	for y, v :=range platform {
		for x, char := range v {
			switch char {
			case 'O':
				//best idea of the day. I started with just trying the weights but they repeated out of cycle
				//primes and non zero based to avoid noncycle duplicates
				score += 7 * ((y +1)*17) * ((x+1)*23)
			}
		}
	}
	return score
}

func sumWeights(platform [][]rune) int {
	sum := 0
	for y, row := range platform {
		for _, char := range row {
			if char == 'O'{
				sum += len(platform) - y
			}
		}
	}

	return sum
}

func spin(platform [][]rune) [][]rune {
	return roll(roll(roll(roll(platform, North), West), South), East)
}

func roll(platform [][]rune, direction int) [][]rune{
	var start, lastRow, iterateValue, outerLoopLength int
	switch direction {
	case North:
		start = 0
		lastRow = len(platform)
		iterateValue = 1
		outerLoopLength = len(platform[0])
	case South:
		start = len(platform) - 1
		lastRow = -1
		iterateValue = -1
		outerLoopLength = len(platform[0])
	case West:
		start = 0
		lastRow = len(platform[0])
		iterateValue = 1
		outerLoopLength = len(platform)
	case East:
		start = len(platform[0]) - 1
		lastRow = -1
		iterateValue = -1
		outerLoopLength = len(platform)
	}
	for _, x := range loopRange(0, outerLoopLength, 1) {
		lastSolid := start - iterateValue
		for _, y := range loopRange(start, lastRow, iterateValue) {
			//changes position and helps keep track of where we can instantly roll to
			lastSolid = changePosition(x, y, direction, lastSolid, iterateValue, platform)
		}
	}
	return platform
}
func changePosition(outer, inner, direction, lastSolid, iterateValue int, platform [][]rune) int {
	var x,y, newSolid, rollingX, rollingY int
	if North == direction || South == direction {
		x = outer
		y = inner
		newSolid = y
		rollingX = x
		rollingY = lastSolid + iterateValue

	}else {
		x = inner
		y = outer
		newSolid = x
		rollingX = lastSolid + iterateValue
		rollingY = y
	}

	switch platform[y][x] {
	case '#':
		return newSolid
	case 'O':
		platform[y][x] = '.'
		platform[rollingY][rollingX] = 'O'
		return lastSolid + iterateValue
	default:
		return lastSolid
	}

}

func columnIterate(start int, iteration int) func(int) bool{
	step := start
	//end exclusive?
	return func(end int) bool{
		step+=iteration
		return step != end
	}
}

func loopRange(start, end, iteration int) []int {
	loop := make([]int, 0)
	for i := start; i != end; i+=iteration {
		loop = append(loop, i)
	}
	return loop
}

func rollNorth(platform [][]rune) [][]rune{
	for x := 0; x < len(platform[0]); x++ {
		lastSolid := -1
		for y := 0; y < len(platform); y++ {
			switch platform[y][x] {
			case '#':
				lastSolid = y
			case 'O':
				lastSolid = lastSolid + 1
				platform[y][x] = '.'
				platform[lastSolid][x] = 'O'
			default:
				//do nothing
			}
		}
	}
	return platform

}
const (
	North = 0
	East = 1
	South = 2
	West = 3
)
func parseLines(input string) [][]rune {
	platform := make([][]rune, 0)
	for _, lineString := range strings.Split(input, "\n") {
		line := make([]rune, len(lineString))
		for x, char := range lineString{
			line[x] = char
		}
		platform = append(platform, line)
	}

	return platform
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
