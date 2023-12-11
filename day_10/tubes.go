package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	//part 1
	start, length, area := mapPipe(parseLines(readInput()))
	fmt.Println(start, length, area)
	fmt.Println(length/2)

	//part 2
	//Pick's
	//AREA = Inner + Border/2 - 1
	//Inner = Area + 1 - Border/2
	fmt.Println(area - (length/2) + 1 )
}
func mapPipe(start Pipe, field [][]rune) (looped Pipe, len int, area int){
	len = 1//because we add the first connection manually
	x, y := findFirstConnection(start, field)
	//Shoelace Formula area
	//A = ([x1 * y2 - y1 * x2] + [x2 * y3 - y2 * x3]...and so on) / 2
	area =  start.X * y - start.Y * x
	next := Pipe{x, y, field[y][x], &start, nil}
	start.b = &next
	for next.Symbol != 'S' {
		current := next
		next = nextPipe(current, field)
		len+=1
		area += current.X * next.Y - current.Y * next.X
	}
	//overwrite end just to do a full loop
	next.b = &start

	return start, len, area/2
}

func nextPipe(current Pipe, field [][]rune) Pipe {
	nextX, nextY := 0, 0
	switch current.Symbol {
	case '|':
		nextX = current.X
		//headed North?
		if current.a.Y > current.Y {
			nextY = current.Y-1
		}else {
			nextY = current.Y+1
		}
	case '-':
		nextY = current.Y
		//headed west?
		if current.a.X > current.X {
			nextX = current.X - 1
		}else {
			nextX = current.X + 1
		}
	case 'L':
		//headed west?
		if current.a.X > current.X {
			nextX = current.X
			nextY = current.Y -1
		}else {
			nextX = current.X + 1
			nextY = current.Y
		}
	case 'J':
		//headed south?
		if current.a.Y < current.Y {
			nextX = current.X - 1
			nextY = current.Y
		}else {
			nextX = current.X
			nextY = current.Y -1
		}
	case '7':
		if current.a.X < current.X {
			nextX = current.X
			nextY = current.Y + 1
		}else {
			nextY = current.Y
			nextX = current.X - 1
		}
	case 'F':
		if current.a.Y > current.Y {
			nextY = current.Y
			nextX = current.X + 1
		}else {
			nextY = current.Y + 1
			nextX = current.X
		}
	case '.':
		panic("Shouldn't have gotten here")
	}

	next := Pipe{nextX, nextY, field[nextY][nextX], &current, nil}
	current.b = &next
	return next

}
func findFirstConnection(start Pipe, field[][]rune) (x,y int) {
	switch  {
	case strings.ContainsAny(string(field[start.Y-1][start.X]),  "|7F")://go north first
		return start.X, start.Y-1
	case strings.ContainsAny(string(field[start.Y+1][start.X]),  "|LJ")://go south first
		return start.X, start.Y + 1
	case strings.ContainsAny(string(field[start.Y][start.X-1]),  "-LF")://go west first
		return start.X -1, start.Y
	case strings.ContainsAny(string(field[start.Y][start.X+1]),  "-J7")://go east first
		return start.X + 1, start.Y
	default:
		panic("no starts")
	}
}
func parseLines(lines string) ( start Pipe, field [][]rune) {
	field = make([][]rune, 0)
	for i, line := range strings.Split(lines, "\n") {
		row := make([]rune, 0)
		for j, char := range line {
			row = append(row, char)
			if char == 'S' {
				start = Pipe{j, i, char, nil, nil}
			}
		}
		field = append(field, row)
	}
	return start, field
}
type (
	Pipe struct {
		X, Y int
		Symbol rune
		a *Pipe
		b *Pipe

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
