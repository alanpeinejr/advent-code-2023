package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main(){
	//part 1
	engine := parseInput(readInput())
	fmt.Println(sum(neighbors(engine)))
	fmt.Println(sum(gears(engine)))
	//validGameIds := disqualify(Pull{12, 13, 14}, games)
	//fmt.Println(sum(validGameIds))
	//part 2

}

func sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum+= value
	}
	return sum
}
func gears(engine [][]rune) []int {
	//find its bounds
	//find and count the numbers
	gearRatios := make([]int, 0)
	y := 0
	for y < len(engine) {
		x := 0
		for x < len(engine[y]) {
			//find a gear
			if engine[y][x] == '*'{
				gearRatios = append(gearRatios, gearNeighborsRatio(engine, x, y))
			}
			x+=1
		}
		y+=1
	}

	return gearRatios
}
func neighbors(engine [][]rune) []int {
	numbersWithAdjacents := make([]int, 0)
	y := 0
	for y < len(engine) {
		x := 0
		for x < len(engine[y]) {
			//find digit
			number := getNumberFromArray(engine[y][x:])
			if len(number) > 0{
				//check full numbers neighbors
				//advance the x value by appropriate number length
				if neighborIsSymbol(engine, number, x, y) {
					value := stringToint(number)
					numbersWithAdjacents = append(numbersWithAdjacents, value)
				}
				x+=len(number)
				continue
			}
			x+=1
		}
		y+=1
	}
	return numbersWithAdjacents
}
func gearNeighborsRatio(engine [][]rune, x int, y int) int {
	ratio := 0
	minX,minY,maxX,maxY :=x-3,y-1,x+3,y+1//unsafe version
	if minX < 0 {
		minX = 0
	}
	if minY < 0 {
		minY = 0
	}
	if maxX >= len(engine[y]) {
		maxX = len(engine[y]) - 1
	}
	if maxY >= len(engine) {
		maxY = len(engine) - 1
	}
    //min,max's now safe, go check for numbers
	numbers := searchForNumbers(engine, x, y, minX, minY, maxX, maxY)
	//if exactly 2, gear ratio
	if len(numbers) == 2 {
		ratio = numbers[0] * numbers[1]
	}
	return ratio
}
func searchForNumbers(engine [][]rune, x int, y int, minX int, minY int, maxX int, maxY int) []int{
	numbers := make([]int, 0)
	yIter := minY
	for yIter <= maxY {
		xIter:= minX
		for xIter <= maxX {
			number := getNumberFromArray(engine[yIter][xIter:])
			if len(number) > 0{
				//we have a number, but its possible to not border our coordinate, ie 78.*
				//numbers START or its END must be within 1 space of our coord
				if distance(x, y, xIter+len(number) - 1, yIter) == 1  || distance(x, y, xIter, yIter) == 1{
					numbers = append(numbers, stringToint(number))
					xIter+=len(number)
					continue
				}
			}
			xIter+=1
		}
		yIter+=1
	}
	return numbers
}
//distance formula that floors so diags are 1 away
func distance(x1 int, y1 int, x2 int, y2 int) int {
	xSquared := (x2 -x1) * (x2-x1)
	ySquared := (y2 -y1) * (y2-y1)
	distance := math.Sqrt(float64(xSquared + ySquared))
	return int(math.Floor(distance))
}
func neighborIsSymbol(engine [][]rune, value string, x int, y int) bool{
	//find my bounds
	minX,minY,maxX,maxY :=x,y,x+len(value)-1,y//we know this is safe
	if minX > 0 {
		minX = minX-1
	}
	if maxX < len(engine[y]) -1 {
		maxX = maxX+1
	}
	if minY > 0 {
		minY = minY-1
	}
	if maxY < len(engine) -1{
		maxY = maxY+1
	}
	yIter := minY
	for yIter <= maxY {
		xIter:= minX
		for xIter <= maxX {
			if !isEmpty(engine[yIter][xIter]) && !isDigit(engine[yIter][xIter] ){
				return true
			}
			xIter+=1
		}
		yIter+=1
	}
	return false
}

func isEmpty(char rune) bool{
	return char == '.'
}
func isDigit(char rune) bool {
	return unicode.IsDigit(char)
}

func getNumberFromArray(numberRunes []rune) string {
	value := ""
	for _, char := range numberRunes {
		if unicode.IsDigit(char){
			value+=string(char)
		}else{
			//stops the loop at first non digit
			break
		}
	}
	return value
}
func parseInput(input string) [][]rune {
	rows := make([][]rune, 0)
	strings.Split(input, "\n")
	for _, stringRow := range strings.Split(input, "\n") {
		values := make([]rune, 0)
		for _, char := range stringRow {
			values = append(values, char)
		}
		rows = append(rows, values)
	}
	return rows
}



func stringToint(this string) int {
	value, _ := strconv.Atoi(this)
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
