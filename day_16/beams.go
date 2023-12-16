package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	//part 1
	fmt.Println(len(removeDoubleVisits(directLight(-1, 0, East, parseLines(readInput()), map[string]int{}))))
	//part 2
	cave := parseLines(readInput())
	max := 0
	for x := 0; x < len(cave[0]); x++ {
		energized := len(removeDoubleVisits(directLight(x, -1, South, cave, map[string]int{})))
		if max  <  energized{
			max = energized
		}
	}
	for x := 0; x < len(cave[0]); x++ {
		energized := len(removeDoubleVisits(directLight(x, len(cave) - 1, North, cave, map[string]int{})))
		if max  <  energized{
			max = energized
		}
	}
	for y := 0; y < len(cave); y++ {
		energized := len(removeDoubleVisits(directLight(-1, y, East, cave, map[string]int{})))
		if max  <  energized{
			max = energized
		}
	}
	for y := 0; y < len(cave); y++ {
		energized := len(removeDoubleVisits(directLight(len(cave[0]) -1, y, West, cave, map[string]int{})))
		if max  <  energized{
			max = energized
		}
	}
	fmt.Println(max)

}
func removeDoubleVisits(visited map[string]int) map[string]int{
	deduped := map[string]int{}
	for k, v := range visited{
		deduped[k[:len(k) -2]]+= v
	}
	return deduped
}
func directLight(x, y, direction int, cave [][]rune, pointsVisited map[string]int) map[string]int {
	shineOn := true
	for shineOn {
		x, y = nextPoint(x, y, direction)
		if x < 0 || y < 0 || x > len(cave[0]) - 1 || y > len(cave) - 1{
			break
		}else if pointsVisited[toString(x, y, direction)] > 1 {
			return pointsVisited
		}
		pointsVisited[toString(x, y, direction)] += 1
		switch cave[y][x] {
		case '.':
		case '\\':
			direction = nextDirection(direction, '\\')
		case '/':
			direction = nextDirection(direction, '/')
		case '|':
			if direction == North || direction == South {
				continue
			}
			shineOn = false
			directLight(x, y, North, cave, pointsVisited)
			directLight(x, y, South, cave, pointsVisited)
		case '-':
			if direction == East || direction == West {
				continue
			}
			shineOn = false
			directLight(x, y, East, cave, pointsVisited)
			directLight(x, y, West, cave, pointsVisited)
		}
	}

	return pointsVisited
}

func toString(x,y, direction int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, direction)
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
	return -1, -1
}
func nextDirection(direction int, mirror rune) int {
	if mirror == '/' {
		switch direction {
		case North:
			return East
		case East:
			return North
		case South:
			return West
		case West:
			return South
		}
	}else {
		// \
		switch direction {
		case North:
			return West
		case West:
			return North
		case South:
			return East
		case East:
			return South
		}
	}
	panic("whoops")
}
func parseLines(input string) [][]rune {
	cave := make([][]rune, 0)
	for _, row := range strings.Split(input, "\n"){
		chars := make([]rune, len(row))
		for i, char := range row {
			chars[i] = char
		}
		cave = append(cave, chars)
	}

	return cave
}

const (
	North = 0
	East = 1
	South = 2
	West = 3
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
