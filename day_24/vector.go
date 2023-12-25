package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main(){
	//part 1
	vectors := parseLines(readInput())
	fmt.Println(countIntersections2D(vectors, 200000000000000, 400000000000000))


	//part 2
	x, y, z := theOdds(vectors)
	fmt.Println(x + y + z)

}

func theOdds(vectors []Vector) (int, int, int){
	xVs, yVs, zVs := potentialVelocities(vectors)
	fmt.Println(zVs)
	if len(xVs) == len(yVs) && len(yVs) == len(zVs) && len(zVs) == 1 {
		rXv, rYv, rZv := float64(xVs[0]), float64(yVs[0]), float64(zVs[0])
		a, b := vectors[0], vectors[1]
		mA := (a.Yv - rYv) / (a.Xv - rXv)
		mB := (b.Yv - rYv) / (b.Xv - rXv)
		cA := a.Y - (mA * a.X)
		cB := b.Y - (mB * b.X)
		xPos := (cB - cA) / (mA - mB)
		yPos := mA * xPos + cA
		time := (xPos - a.X) / (a.Xv - rXv)
		zPos := a.Z + (a.Zv - rZv) * time
		return int(xPos), int(yPos), int(zPos)
	}
	fmt.Println("no possibles")
	return -1, -1, -1
}

func potentialVelocities(vectors []Vector) ([]int, []int, []int){
	x, y, z := []int{}, []int{}, []int{}
	for i:= 0; i < len(vectors) - 1; i++ {
		for j:= i + 1; j < len(vectors); j++ {
			a, b := vectors[i], vectors[j]
			if a.Xv == b.Xv {
				maybe := matchingVelocity(int(b.X - a.X), int(a.Xv))
				if len(maybe) == 0 {
					x = maybe
				}else {
					x = intersect(x, maybe)
					fmt.Println(x)
				}
			}
			if a.Yv == b.Yv {
				maybe := matchingVelocity(int(b.Y - a.Y), int(a.Yv))
				if len(maybe) == 0 {
					y = maybe
				}else {
					y = intersect(y, maybe)
					fmt.Println(y)

				}
			}
			if a.Zv == b.Zv {
				maybe := matchingVelocity(int(b.Z - a.Z), int(a.Zv))
				if len(maybe) == 0 {
					z = maybe
				}else {
					z = intersect(z, maybe)
					fmt.Println(z)

				}
			}
		}
	}
	return x, y, z
}

func matchingVelocity(dV, pV int) []int {
	match := make([]int, 0)
	for v:= -1000; v < 1000; v++ {
		if v != pV && dV % (v - pV) == 0 {
			match = append(match, v)
		}
	}
	return match
}

func intersect(a, b []int) []int {
	intersections := make([]int, 0)
	for _, value := range a {
		if slices.Contains(b, value){
			intersections = append(intersections, value)
		}
	}
	return intersections
}
func countIntersections2D(vectors []Vector, rangeLow, rangeHigh float64) int {
	sum := 0
	for i := 0; i < len(vectors); i++ {
		for j := i + 1; j < len(vectors); j++ {
			intersects, at, iTime, jTime := intersects2D(vectors[i], vectors[j])
			if intersects && iTime > 0 && jTime > 0 {
				if at.X <= rangeHigh && at.X >= rangeLow && at.Y <= rangeHigh && at.Y >= rangeLow {
					sum += 1
				}
			}
		}
	}
	return sum
}
//u := (as.y*bd.x + bd.y*bs.x - bs.y*bd.x - bd.y*as.x ) / (ad.x*bd.y - ad.y*bd.x) as,bs = points
//p = as + ad * u
//p = bs + bd * v

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
