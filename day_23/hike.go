package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main(){
	//part 1
	trail, start, end := parseLines(readInput())
	longest := gohike(trail, start, end, true)
	fmt.Println(longest)
	//print(path, trail)


	//part 2
	longest = gohike(trail, start, end, false)
	fmt.Println(longest)
}

func gohike(trail map[Point]Trail, start, end Point, steepSlopes bool)(int) {
	longest := make(chan int, 1)
	longest <- 0
	//longest := Path{map[Point]bool{}, start}
	first := Path{map[Point]bool{start:true}, start}
	//queue := []Path{first}
	chanQueue := Queue{make(chan Path, 10000)}
	chanQueue.Insert(first)
	for  {
		go makeStep(trail, steepSlopes, chanQueue, end, longest)
	}
	test:= <- longest
	return test
}

func hike(trail map[Point]Trail, start, end Point, steepSlopes bool)(int, Path) {
	longest := Path{map[Point]bool{}, start}
	first := Path{map[Point]bool{start:true}, start}
	queue := []Path{first}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		//we reached the end the slowest so far
		if current.current == end && len(longest.visited) < len(current.visited){
			fmt.Println("Found end: ", len(current.visited) - 1, " Queue Length: ", len(queue))//dont count start
			longest = current
		}

		nextPossibilities := possibleDirections(current, trail, steepSlopes)
		//skips copying the whole path if we aren't at a branch
		if len(nextPossibilities) > 1 {
			for _, nextPoint := range nextPossibilities {
				//take the current path and add it, add to queue
				visited := current.copyVisited()
				visited[nextPoint] = true
				queue = append(queue, Path{visited, nextPoint})
			}
		} else if len(nextPossibilities) == 1 {
			current.current = nextPossibilities[0]
			current.visited[nextPossibilities[0]] = true
			queue = append(queue, current)
		}

		//fmt.Println(len(queue))

	}
	return len(longest.visited) - 1, longest
}

func makeStep(trail map[Point]Trail, steepSlopes bool, queue Queue, end Point, longest chan int)  {
	current, err := queue.Remove()
	if err != nil {
		return
	}
	//we reached the end the slowest so far
	if current.current == end {
		oldLongest := <- longest
		fmt.Println("Found end: ", len(current.visited) - 1, "was: ", oldLongest ," Queue Length: ", len(queue.q))//dont count start
		longest <- len(current.visited) - 1
	}

	nextPossibilities := possibleDirections(current, trail, steepSlopes)
	//skips copying the whole path if we aren't at a branch
	if len(nextPossibilities) > 1 {
		for _, nextPoint := range nextPossibilities {
			//take the current path and add it, add to queue
			visited := current.copyVisited()
			visited[nextPoint] = true
			queue.Insert(Path{visited, nextPoint})
		}
	} else if len(nextPossibilities) == 1 {
		current.current = nextPossibilities[0]
		current.visited[nextPossibilities[0]] = true
		queue.Insert(current)
	}
	return
}

type Queue struct {
	q        chan Path
}

func (q *Queue) Insert(item Path)  {
		q.q <- item
}

func (q *Queue) Remove() (Path, error) {
	if len(q.q) > 0 {
		item := <-q.q
		return item, nil
	}
	return Path{}, errors.New("Queue is empty")
}
//if this causes bottleneck we could do it only on branch points?
func (this Path) copyVisited() (copy map[Point]bool){
	copy = make(map[Point]bool, len(this.visited))
	for key, value := range this.visited {
		copy[key] = value
	}
	return
}
func possibleDirections(path Path, trail map[Point]Trail, steepSlopes bool) []Point{
	this := path.current
	possibilities := make([]Point, 0)
	//had this as a real switch but changed in part 2 to account for ignoring the slopes
	switch {
	case steepSlopes && trail[this].char == '^':
		possibilities = append(possibilities, this.add(North))
	case steepSlopes && trail[this].char == '>':
		possibilities = append(possibilities, this.add(East))
	case steepSlopes && trail[this].char == 'v':
		possibilities = append(possibilities, this.add(South))
	case steepSlopes && trail[this].char == '<':
		possibilities = append(possibilities, this.add(West))
	default:
		//check possibles
		for _, dir := range directions {
			next := this.add(dir)
			if trail[next].char != '#' {
					possibilities = append(possibilities, next)
			}
		}
	}
	//cant backtrack, doing here because i allowed backtracks specifically on slope force movement and wanted once check
	possibilitiesNoBacktracking := make([]Point, 0)
	for _, possibility := range possibilities {
		if _, exists := path.visited[possibility]; !exists{
			possibilitiesNoBacktracking = append(possibilitiesNoBacktracking, possibility)
		}
	}
	return possibilitiesNoBacktracking

}
//func print(path Path, trail map[Point]Trail) {
//	for y:= 0; y < 142; y ++ {//24 for test input
//		for x:=0; x < 142; x++ {//24 for test input
//			point := Point{x,y}
//			if _, exists := path.visited[point]; exists {
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
	Trail struct {
		Point
		char rune
	}
	Path struct {
		visited map[Point]bool
		current Point
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
