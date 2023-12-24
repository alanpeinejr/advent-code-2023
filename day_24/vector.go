package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main(){
	//part 1
	vectors := parseLines(readInput())
	fmt.Println(countIntersections2D(vectors, 200000000000000, 400000000000000))


	//part 2
	fmt.Println()

}

func countIntersections2D(vectors []Vector, rangeLow, rangeHigh float64) int{
	sum := 0
	for i:= 0; i < len(vectors); i++ {
		for j:= i + 1; j < len(vectors); j++ {
			intersects, at, iTime, jTime := intersects2D(vectors[i], vectors[j])
			if intersects  && iTime > 0 && jTime > 0 {
				if at.X <= rangeHigh && at.X >= rangeLow  && at.Y <= rangeHigh && at.Y >= rangeLow {
					sum+=1
				}
			}
		}
	}
	return sum
}


//u := (as.y*bd.x + bd.y*bs.x - bs.y*bd.x - bd.y*as.x ) / (ad.x*bd.y - ad.y*bd.x) as,bs = points
//p = as + ad * u
//v := (as.x + ad.x * u - bs.x) / bd.x
//https://stackoverflow.com/questions/2931573/determining-if-two-rays-intersect
func intersects2D(a, b Vector) (bool, Point, float64, float64){
	aTime := (a.Y*b.Xv + b.Yv*b.X - b.Y*b.Xv - b.Yv*a.X) / (a.Xv*b.Yv - a.Yv*b.Xv)
	x,y := a.X + a.Xv*aTime, a.Y + a.Yv * aTime
	bTime := ( a.X + a.Xv*aTime - b.X) / b.Xv
	//x1,y1 := b.X + b.Xv*bTime, b.Y + b.Yv * bTime
	//if time is infinite they dont cross
	return !math.IsInf(aTime, 0), Point{X: x, Y: y}, aTime, bTime

}

func parseLines(lines string) ([]Vector){
	noFiller := strings.ReplaceAll(strings.ReplaceAll(lines, " @", ""), ",", "")
	vectors := make([]Vector, 0)
	for _, row := range strings.Split(noFiller, "\n") {
		values := strings.Fields(row)
		vectors = append(vectors, Vector{Point{stringToInt(values[0]), stringToInt(values[1]), stringToInt(values[2])}, stringToInt(values[3]), stringToInt(values[4]), stringToInt(values[5])})
	}
	return vectors
}

type (
	Point struct {
		X,Y,Z float64
	}
	Vector struct {
		Point
		Xv,Yv,Zv float64
	}
)
func stringToInt(this string) float64 {
	value, _ := strconv.Atoi(this)
	return float64(value)
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
