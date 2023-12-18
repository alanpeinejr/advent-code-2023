package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main_old(){
	//part 1
	city := parseLines_old(readInput_old())
	fmt.Println(findPointFrom_old(0,0, len(city[0]) - 1, len(city) - 1, city, 1, 3 ))
	//part 2
	fmt.Println(findPointFrom_old(0,0, len(city[0]) - 1, len(city) - 1, city, 4, 10))

}

func findPointFrom_old(startX, startY, endX, endY int, city [][]int, min, max int) int{
	startOption1 := Point_old{min, startY, East_old, 4}
	startOption2 := Point_old{startX, min, South_old, 4}

	//startOption2 := Point_old{startX + 1, startY, East_old, 1}
	unvisited := preFill_old(city, min, max)
	unvisited[startOption1] = 0
	unvisited[startOption2] = 0
	visited := map[Point_old]int{}
	fmt.Println(len(unvisited))
	for len(unvisited) > 0 {
		current := smallestUnvisited_old(unvisited)
		fmt.Println(current)
		cost := city[current.Y][current.X] + unvisited[current]
		visited[current] = cost//?? check if less?
		delete(unvisited, current)

		if current.X == endX && current.Y == endY{
			return visited[current]
		}
		// each next option, is it a valid move, has it been exploredCheaper? enqueue
		for _, next := range current.possibleDirections_old(city, min, max){
			if value, exists := unvisited[next]; exists {
				if value > cost {
					unvisited[next] = cost
				}
			}
		}
	}
	panic("Whoops")

}

func (this Point_old) possibleDirections_old(city [][]int, min, max int) []Point_old {
	possibilities := make([]Point_old, 0)
	for _, possibility := range this.optionFilter_old(city, min, max) {
		nextX, nextY := nextPoint_old_old(this.X, this.Y, possibility, min)
		straightTime := min
		if this.direction == possibility {
			straightTime = this.StraightTime + 1
		}
		possibilities = append(possibilities, Point_old{nextX, nextY,possibility, straightTime})
	}

	return possibilities
}

func (this Point_old) optionFilter_old(city [][]int, min, max int) []int {
	options := make([]int, 0)
	if this.StraightTime >= min {
		options = append(options, turnDirections_old(this.direction)...)
	}
	if this.StraightTime < max {
		options = append(options, this.direction)
	}
	withinBounds := make([]int, 0)
	for _, possibility := range options {
		nextX, nextY := nextPoint_old_old(this.X, this.Y, possibility, min)
		if nextX >= 0 && nextX < len(city[0]) && nextY >= 0 && nextY < len(city) {
			withinBounds = append(withinBounds, possibility)
		}
	}
	return withinBounds
}
func turnDirections_old(direction int) []int {
	switch direction {
	case North_old:
		return []int{East_old, West_old}
	case East_old:
		return []int{North_old, South_old}
	case South_old:
		return []int{East_old, West_old}
	case West_old:
		return []int{North_old, South_old}
	}
	panic("whoops")
}
func nextPoint_old_old(x, y, direction, min int) (int, int){
	switch direction {
	case North_old:
		return x, y - min
	case East_old:
		return x + min, y
	case South_old:
		return x , y + min
	case West_old:
		return x - min, y
	}
	panic("whoops")
}

func smallestUnvisited_old(djiktstras map[Point_old]int) (Point_old) {
	minCost := math.MaxInt
	key := Point_old{-1,-1,-1,-1}
	for k, v := range djiktstras {
		if minCost > v {
			minCost = v
			key = k
		}
	}
	return key
}
func preFill_old(city [][]int, min, max int) (prefill map[Point_old]int){
	prefill = map[Point_old]int{}
	for y, row := range city {
		for x, _ := range row {
			for direction:=0; direction < 4; direction++{
				for i := min; i <= max; i++ {
					point := Point_old{x, y, direction, i}
					prefill[point] = math.MaxInt
				}
			}

		}
	}
	return prefill
}

func parseLines_old(input string) [][]int {
	city := make([][]int, 0)
	for _, row := range strings.Split(input, "\n"){
		chars := make([]int, len(row))
		for i, char := range strings.Split(row, "") {
			chars[i] = stringToInt(char)
		}
		city = append(city, chars)
	}

	return city
}

const (
	North_old = 0
	East_old = 1
	South_old = 2
	West_old = 3
)

type (
	Point_old struct {
		X    int
		Y    int
		direction int
		StraightTime int
	}
)
func stringToInt_old(this string) int {
	value, _ := strconv.Atoi(this)
	return value
}
func readInput_old() string {
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
