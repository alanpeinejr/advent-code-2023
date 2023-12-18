package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//part 1
	city := parseLines(readInput())
	fmt.Println(findPointFrom(0,0, len(city[0]) - 1, len(city) - 1, city, 1, 3 ))
	//part 2
	fmt.Println(findPointFrom(0,0, len(city[0]) - 1, len(city) - 1, city, 4, 10 ))

}

func findPointFrom(startX, startY, endX, endY int, city [][]int, min, max int) int{

	queue := NewPriorityQueue[QueueItem](func(a, b QueueItem) int {
		return a.Cost - b.Cost
	})

	queue.Push(QueueItem{Point:Point{startX, startY, East, 1}})
	queue.Push(QueueItem{Point:Point{startX , startY, South, 1}})

	visited := map[Point]int{}

	for !queue.isEmpty() {
		current := queue.Pop()

		cost := city[current.Y][current.X] + current.Cost
		//if its the end, we're done
		if current.X == endX && current.Y == endY{
			return cost
		}

		if v, exists := visited[current.Point]; exists {
			if v <= cost {
				continue
			}
		}
		visited[current.Point] = cost

		for _, next := range current.possibleDirections(city, min, max) {
			queue.Push(QueueItem{next, cost})
		}


	}
	panic("Whoops")

}

func (this Point) possibleDirections(city [][]int, min, max int) []Point {
	possibilities := make([]Point, 0)
	for _, possibility := range this.optionFilter(city, min, max) {
		nextX, nextY := nextPoint(this.X, this.Y, possibility)
		straightTime := 1
		if this.direction == possibility {
			straightTime = this.StraightTime + 1
		}
		possibilities = append(possibilities, Point{nextX, nextY, possibility, straightTime})
	}

	return possibilities
}
func (this Point) optionFilter(city [][]int, min, max int) []int {
	options := make([]int, 0)
	if this.StraightTime >= min {
		options = append(options, turnDirections(this.direction)...)
	}
	if this.StraightTime < max {
		options = append(options, this.direction)
	}
	withinBounds := make([]int, 0)
	for _, possibility := range options {
		nextX, nextY := nextPoint(this.X, this.Y, possibility)
		if nextX >= 0 && nextX < len(city[0]) && nextY >= 0 && nextY < len(city) {
			withinBounds = append(withinBounds, possibility)
		}
	}
	return withinBounds
}
func turnDirections(direction int) []int {
	switch direction {
	case North:
		return []int{East, West}
	case East:
		return []int{North, South}
	case South:
		return []int{East, West}
	case West:
		return []int{North, South}
	}
	panic("whoops")
}
func nextPoint(x, y, direction int) (int, int){
	switch direction {
	case North:
		return x, y - 1
	case East:
		return x + 1, y
	case South:
		return x , y + 1
	case West:
		return x - 1, y
	}
	panic("whoops")
	return -1, -1
}
func (this Point) toString() string {
	return fmt.Sprintf("%d,%d", this.X, this.Y)
}

func parseLines(input string) [][]int {
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
	North = 0
	East = 1
	South = 2
	West = 3
)

type (
	Point struct {
		X    int
		Y    int
		direction int
		StraightTime int
	}
	QueueItem struct {
		Point
		Cost int
	}
)
func stringToInt(this string) int {
	value, _ := strconv.Atoi(this)
	return value
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