package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main(){
	//part 1
	trail, start, end := parseLines(readInput())
	fmt.Println(findLength(trail, start, end, true))


	//part 2
	fmt.Println(findLength(trail, start, end, false))

}

func hike(nodes map[Point][]Cost, current, end Point, visited map[Point]bool, cost, max int) int {
	if current == end {
		if cost > max {
			max = cost
		}
		return max
	}

	visited[current] = true

	for _, to := range nodes[current] {
		if !visited[to.Point] {
			max = hike(nodes, to.Point, end, visited, cost + to.Cost, max)
		}
	}
	visited[current] = false
	return max
}

func findLength(trail map[Point]Trail, start, end Point, steepSlopes bool) int {
	cost := 0
	nodes := findNodes(trail, start, steepSlopes)

	//skip end
	if !steepSlopes && len(nodes[end]) > 0 {
		cost = nodes[end][0].Cost
		end = nodes[end][0].Point
	}

	visited := map[Point]bool{}
	return hike(nodes, start, end, visited, cost, 0)
}

func findNodes(trail map[Point]Trail, start Point, steepSlopes bool) map[Point][]Cost {
	costs := map[Point][]Cost{}
	queue := []Point{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			next := current.add(dir)
			if trail[next].char != '#' {
				cost, valid := nodesPath(trail, current, next, 1, steepSlopes)
				if valid && !slices.Contains(costs[current], cost ) {
					costs[current] = append(costs[current], cost)
					queue = append(queue, cost.Point)
				}
			}
		}
	}
	return costs
}

func nodesPath(trail map[Point]Trail, previous, current Point, cost int, steepSlopes bool) (Cost, bool) {
	options := 0
	for _, dir := range directions {
		next := current.add(dir)
		if trail[next].char != '#' {
			options += 1
		}
	}
	//node point
	if options > 2 {
		return Cost{current, cost}, true
	}

	if steepSlopes {
		char := trail[current].char
		if char != '.' {
			if current.X > previous.X && char != '>' {
				return Cost{}, false
			}
			if current.X < previous.X && char != '<' {
				return Cost{}, false
			}
			if current.Y > previous.Y && char != 'v' {
				return Cost{}, false
			}
			if current.Y < previous.Y && char != '^' {
				return Cost{}, false
			}
		}
	}

	//follow path
	for _, dir := range directions {
		next := current.add(dir)
		if trail[next].char != '#' && next != previous {
			return nodesPath(trail, current, next, cost + 1, steepSlopes)
		}
	}

	return Cost{current, cost}, true
}


//func print(visited map[Point]bool, trail map[Point]Trail) {
//	for y:= 0; y < 24; y ++ {//24 for test input,142
//		for x:=0; x < 24; x++ {//24 for test input,142
//			point := Point{x,y}
//			if _, exists := visited[point]; exists {
//				fmt.Print("O")
//			}else {
//				fmt.Print(string(trail[point].char))
//			}
//		}
//		fmt.Println()
//	}
//}


func (this Point) add(that Point) Point {
	return Point{this.X + that.X, this.Y + that.Y}
}
var directions = []Point{North, East, South, West}
var (
	North = Point{0, -1}
	East = Point{1, 0}
	South = Point{0, 1}
	West = Point{-1, 0}

)
func forceInbounds(trail map[Point]Trail, maxX, maxY int) {
	for x := -1; x <= maxX; x++ {
		north := Point{x, -1}
		south := Point{x, maxY}
		trail[north] = Trail{north, '#'}
		trail[south] = Trail{south, '#'}
	}
	for y := -1; y <= maxY; y++ {
		west := Point{-1, y}
		east := Point{maxX, y}
		trail[west] = Trail{west, '#'}
		trail[east] = Trail{east, '#'}
	}
}

func parseLines(lines string) (trail map[Point]Trail, start, end Point){
	trail = map[Point]Trail{}
	rows := strings.Split(lines, "\n")
	for y, row := range rows{
		for x, char := range row {
			point := Point{x,y}
			trail[point] = Trail{point, char}
			if y == 0 && char == '.'{
				start = point
			}
			if y == len(row) - 1 && char == '.'{
				end = point
			}
		}

	}
	//make a wall
	forceInbounds(trail, len(rows[0]), len(rows))

	return
}

type (
	Point struct {
		X,Y int
	}
	Cost struct {
		Point
		Cost int
	}
	Trail struct {
		Point
		char rune
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
