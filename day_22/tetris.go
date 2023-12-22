package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main(){
	//part 1
	blocks := parseLines(readInput())
	slices.SortFunc(blocks, compareBlock)
	blocksFall(blocks)
	//for _, test := range blocks {
	//	fmt.Println(test.label)
	//	fmt.Print("supports: ")
	//	for support, _ := range test.supports {
	//		fmt.Print(support.label, " ")
	//	}
	//	fmt.Print("supportedBy:")
	//	for support, _ := range test.supportedBy {
	//		fmt.Print(support.label, " ")
	//	}
	//	fmt.Println("")
	//}
	fmt.Println(countDelatables(blocks))

	//part 2
	fmt.Println(sum(chainReactions(blocks)))
}

func sum(results []int) int {
	sum := 0
	for _, val := range results{
		sum += val
	}
	return sum
}

func countDelatables(blocks []*Block) int{
	//if a blocks supported.supportedBy is more than 1, it can be deleted
	deletable := map[*Block]bool{}
	for _, block := range blocks {
		if canDelete(block) {
			deletable[block] = true
		}
	}

	return len(deletable)
}

func chainReactions(blocks []*Block) []int{
	reactionResults := make([]int, len(blocks))
	for i, block := range blocks {
		reaction := map[*Block]bool{block:true}
		block.disintegrate(reaction)
		//first block to go doesnt count
		reactionResults[i] = len(reaction) - 1
	}
	return reactionResults
}
func (this *Block) disintegrate(reaction map[*Block]bool) {
	//for who I'm supporting, check if their supports still "exists"
	for support ,_ := range this.supports{
		redundanciesStillExist := false
		for redudancies,_ := range support.supportedBy {
			if _, exists := reaction[redudancies]; !exists{
				redundanciesStillExist = true
			}
		}
		//if not, add them to reaction, recursion
		if !redundanciesStillExist{
			reaction[support] = true
			support.disintegrate(reaction)
		}
	}
}

func canDelete(block * Block) bool {
	for supporting, _ := range block.supports {
		if len(supporting.supportedBy) ==  1{
			return false
		}
	}
	return true
}

func blocksFall(blocks []*Block) {
	stationaryBlocks := map[Point]*Block{}
	for _, block := range blocks {
		block.drop(stationaryBlocks)
	}
}

//because we assume each lower block has already fallen, fall til we hit a block, record on both sides
func (this *Block) drop(stationaryBlocks map[Point]*Block ) {
	for this.points[0].Z != 1 {
		stationary := false
		newPoints := make([]Point, len(this.points))
		for i, point := range this.points {
			newPoints[i] = Point{point.X, point.Y, point.Z - 1}
			if blocker, exists := stationaryBlocks[newPoints[i]]; exists {
				stationary = true
				blocker.supports[this] = true
				this.supportedBy[blocker] = true
			}
		}
		// record all points where we've collided, could be multi blocks
		if stationary {
			for _, point := range this.points {
				stationaryBlocks[point] = this
			}
			return
		}else {
			//set new points having fallen 1 space successfully
			this.points = newPoints
		}
	}
	//if we got here we hit the ground
	for _, point := range this.points {
		stationaryBlocks[point] = this
	}
}

func comparePoint(a,b Point) int {
	//is there any reason to check X/Y?
	return a.Z - b.Z
}
func compareBlock(a, b *Block) int {
	return a.points[0].Z - b.points[0].Z
}

func formBlock(a,b Point) []Point {
	points := make([]Point, 0)
	for x:= a.X; x <= b.X; x++ {
		for y := a.Y; y <= b.Y; y++ {
			for z := a.Z; z <= b.Z; z++{
				points = append(points, Point{x,y,z})
			}
		}
	}

	return points
}
func parseLines(lines string) []*Block{
	blocks := make([]*Block, 0)
	for i, line := range strings.Split(lines, "\n"){
		endpoints := strings.Split(line, "~")
		aFields := strings.Split(endpoints[0], ",")
		a := Point{stringToInt(aFields[0]), stringToInt(aFields[1]), stringToInt(aFields[2])}
		bFields := strings.Split(endpoints[1], ",")
		b := Point{stringToInt(bFields[0]), stringToInt(bFields[1]), stringToInt(bFields[2])}

		points := formBlock(a, b)
		slices.SortFunc(points, comparePoint)
		blocks = append(blocks, &Block{i,points, map[*Block]bool{}, map[*Block]bool{}})
	}
	return blocks
}

type (
	Point struct {
		X,Y,Z int
	}
	Block struct {
		label int
		points []Point
		supports map[*Block]bool
		supportedBy map[*Block]bool
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
