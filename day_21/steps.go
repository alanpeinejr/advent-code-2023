package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	//part 1
	start, field := parseLines(readInput())
	reachable, firstVisits := stepToCount(130, start, field)
	fmt.Println(len(firstVisits))

	//print(reachable, field)
	fmt.Println(len(reachable))

	//part 2
	//2356 # in my input
	//131x131
	// 7451 even always
	// 7344 odd always
	//14795 first visits
	//10 unreachable??
	//printUnreachable(firstVisits, field)///yep, 10 unreachable
	//at exactly 130, we visit every reachable point once
	fmt.Println(part2(firstVisits))
}

func part2(firsts map[Point]int) int{
	evenFull := 0
	evenPastHalf := 0
	oddFull := 0
	oddPastHalf := 0
	for _, value := range firsts{
		if value % 2 == 0 {
			evenFull+=1
			if value > 130/2 {
				evenPastHalf +=1
			}
		}else {
			oddFull+=1
			if value > 130/2 {
				oddPastHalf +=1
			}
		}
	}
	//   O
	//  OEO
	// OEOEO
	//OEOEOEO
	// OEOEO
	//  OEO
	//   O
	//if even cycle # Os is cycle+1^2, Es # = cycle^2
	//26501365 is steps required, 131 is length of field, 130/2 is length to fill half
	cycle := (26501365 - 130/2)/131
	fmt.Println(cycle, evenFull, evenPastHalf, oddFull, oddPastHalf)//evens fulls match 499/500 step counts i did above
	totalFilled := (cycle * cycle * evenFull) + ((cycle + 1) * (cycle + 1) * oddFull)
	cornersRemoved := (cycle + 1) * oddPastHalf
	cornersAdded := cycle * (evenPastHalf - 1)//i dont know why - 1, S maybe?, I was off by one cycle length and this is only place where that math could work out
	return totalFilled - cornersRemoved + cornersAdded
}

func print(points map[Point]int, field [][]rune){
	for y, row := range field{
		rowString := make([]rune, len(row))
		for x, char := range row{
			rowString[x] = char
			point := Point{x,y}
			if _, exists :=points[point]; exists{
				rowString[x] = 'O'
			}
			fmt.Print(string(rowString[x]))
		}
		fmt.Println()
	}
}
func printUnreachable(firsts map[Point]int, field [][]rune) {
	for y, row := range field{
		for x, _ := range row{
			point := Point{x,y}
			if _, exists :=firsts[point]; !exists && field[y][x] != '#'{
				fmt.Println(x,y)
			}
		}
	}
}
func stepToCount(steps int, start Point, field [][]rune) (map[Point]int, map[Point]int) {
	firstVisit := map[Point]int{}
	positions := map[Point]int{start:0}
	for i := 0; i < steps; i++{
		newPositions := map[Point]int{}
		for point, _ := range positions{
			for _, step := range point.step(field) {
				newPositions[step]=i + 1
				if _, exists := firstVisit[step]; !exists{
					firstVisit[step] = i + 1
				}
			}
		}
		positions = newPositions
	}
	return positions, firstVisit
}

func (this Point) step(field [][]rune) []Point{
	possibleSteps := make([]Point, 0)
	for _, move := range moves {
		possibility := this.add(move)
		if possibility.inBounds(field) && field[possibility.Y][possibility.X] != '#' {
			possibleSteps = append(possibleSteps, possibility)
		}
	}
	return possibleSteps
}

var moves = []Point{{0,1}, {1,0}, {-1, 0}, {0, -1}}

func (this Point) add(that Point) Point {
	return Point{this.X + that.X, this.Y + that.Y}
}
func (this Point) inBounds(field [][]rune) bool {
	return this.X >= 0 && this.X < len(field[0]) && this.Y >= 0 && this.Y < len(field)
}
func parseLines(lines string) ( start Point, field [][]rune) {
	field = make([][]rune, 0)
	for i, line := range strings.Split(lines, "\n") {
		row := make([]rune, 0)
		for j, char := range line {
			if char == 'S' {
				start = Point{j, i}
			}
			row = append(row, char)

		}
		field = append(field, row)
	}
	return start, field
}
type (
	Point struct {
		X,Y int
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
