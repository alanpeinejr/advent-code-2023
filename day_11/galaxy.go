package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	//part 1
	galaxies, xCosts, yCosts := parseLines(readInput(), 2)
	pairs := createPairs(galaxies)
	fmt.Println(calculateMinLengthsSum(pairs, xCosts, yCosts))

	//part 2
	//part 1
	galaxies, xCosts, yCosts = parseLines(readInput(), 1000000)
	pairs = createPairs(galaxies)
	fmt.Println(calculateMinLengthsSum(pairs, xCosts, yCosts))
	fmt.Println()
}
func calculateMinLengthsSum(pairs []Pair, xCosts map[int]int, yCosts map[int]int) int {
	minLengthSums := 0
	for i, _ := range pairs {
		pairs[i].Distance = pairs[i].calculateDistance(xCosts, yCosts)
		minLengthSums += pairs[i].Distance
	}
	return minLengthSums
}

func (this Pair) calculateDistance(xCosts, yCosts map[int]int) int{
	cost :=0
	//no diagonal distances so down then over is easiest

	//first why is always less than, or equal to
	for y := this.A.Y; y < this.B.Y; y++{
		//moving INTO a zone is what costs movement
		cost += yCosts[y+1]
	}
	//first x could be less so find which is further left
	greaterX := this.A
	lesserX := this.B
	if greaterX.X < lesserX.X {
		greaterX = this.B
		lesserX = this.A
	}

	for x := lesserX.X; x < greaterX.X; x++{
		//moving INTO a zone is what costs movement
		cost += xCosts[x+1]
	}
	return cost
}

func createPairs(galaxies []Galaxy) []Pair {
	pairs := make([]Pair, 0)
	for i, _ := range galaxies {
		for j := i+1; j < len(galaxies); j++ {
			pairs = append(pairs, Pair{galaxies[i], galaxies[j], 0})
		}
	}

	return pairs
}

func parseLines(lines string, expansionFactor int) (galaxies []Galaxy, xCosts map[int]int, yCosts map[int]int) {
	galaxies = make([]Galaxy, 0)
	linesArray := strings.Split(lines, "\n")

	//we 'expand' it by recoding the step cost instead
	yCosts, xCosts = map[int]int{}, map[int]int{}
	for xIndex, _ := range linesArray[0] {
		xCosts[xIndex] = expansionFactor
	}
	for yIndex, _ := range linesArray {
		yCosts[yIndex] = expansionFactor
	}

	//parse lines and expand vertically
	for y, line := range linesArray {
		row := make([]rune, 0)
		for x, char := range line {
			row = append(row, char)
			if char == '#' {
				xCosts[x]=1
				yCosts[y]=1
				galaxies = append(galaxies, Galaxy{x, y})
			}
		}
	}

	return galaxies, xCosts, yCosts
}

type (
	Galaxy struct {
		X, Y int
	}
	Pair struct {
		A,B Galaxy
		Distance int
	}
)
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
