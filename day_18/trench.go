package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//part 1
	//fmt.Println(digTrench(buildCoordinates(parseLines(readInput()))))
	points, perimeter := buildCoordinates(parseLines(readInput()))
	area := digTrench(points)
	//picks for inner boxes
	innerVertexes := area - perimeter/2 + 1
	//add perimeter and inner boxes for our area that counts the line
	fmt.Println(innerVertexes + perimeter)

	//part 2
	points2, perimeter2 := buildCoordinates(parseLinesPart2(readInput()))
	area2 := digTrench(points2)
	innerVertexes2 := area2 - perimeter2/2 + 1
	fmt.Println(innerVertexes2 + perimeter2)
}
func digTrench(points []Point) (area int){
	//Shoelace Formula area
	//A = ([x1 * y2 - y1 * x2] + [x2 * y3 - y2 * x3]...and so on) / 2
	for i := 0; i < len(points) - 1; i++ {
		area += points[i].X * points[i+1].Y - points[i].Y * points[i+1].X
	}
	area += points[len(points) -1 ].X * points[0].Y - points[len(points) -1 ].Y * points[0].X

	return area/2
}

func buildCoordinates(trenches []Trench) ([]Point, int){
	x, y := 0, 0
	points := make([]Point, 0)
	points = append(points, Point{x,y})
	perimeter :=0
	for _, trench := range trenches {
		switch trench.Direction {
		case "U":
			y -= trench.Length
		case "R":
			x += trench.Length
		case "D":
			y += trench.Length
		case "L":
			x -= trench.Length
		}
		perimeter+=trench.Length
		points = append(points, Point{x,y})
	}
	return points, perimeter
}

func parseLines(lines string) ( []Trench) {
	trenches := make([]Trench, 0)
	for _, line := range strings.Split(lines, "\n") {
		fields := strings.Fields(line)
		trenches = append(trenches, Trench{fields[0], stringToInt(fields[1]), fields[2][1:]})

	}
	return trenches
}
func parseLinesPart2(lines string) ( []Trench) {
	trenches := make([]Trench, 0)
	for _, line := range strings.Split(lines, "\n") {
		fields := strings.Fields(line)
		hex :=fields[2][2:len(fields[2]) - 1]
		length := stringToIntADVANCED(hex[:len(hex) - 1])
		var direction string
		switch hex[(len(hex) - 1):] {
		case "0":
			direction = "R"
		case "1":
			direction = "D"
		case "2":
			direction = "L"
		case "3":
			direction = "U"
		}
		trenches = append(trenches, Trench{direction, length, ""})

	}
	return trenches
}
type (
	Trench struct {
		Direction string
		Length int
		Color string
	}
	Point struct {
		X, Y int
	}
)

func stringToInt(this string) int {
	value, _ := strconv.Atoi(this)
	return value
}
func stringToIntADVANCED(this string) int {
	value, _ := strconv.ParseInt(this, 16, 64)
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
