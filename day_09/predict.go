package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//part 1
	herstories := parseLines(readInput())
	//part 2
	fmt.Println(sum(predictHerFuture(herstories)))
	fmt.Println()
}
func predictHerFuture(herstories [][]int) []int {
	herFutures := make([]int, len(herstories))

	for i, v := range herstories {
		fmt.Println(v)
		herFutures[i] = predict(v)
	}

	return herFutures
}
//recursion day?
func predict(herstory []int) int {
	//if the differences are all 0, add a zero, return?
	if onlyContains(herstory, 0) {
		return 0
	}
	//find the differences, predict by adding return value to last element
	differences := make([]int, len(herstory) - 1)
	for i, _ := range differences {
		differences[i] = herstory[i+1] - herstory[i]
	}
	return herstory[len(herstory) - 1] + predict(differences)
}

func onlyContains(array []int, value int) bool {
	for _, v := range array {
		if v != value {
			return false
		}
	}
	return true
}
func parseLines(lines string) ([][]int) {
	herstories := make([][]int, 0)
	for _, line := range strings.Split(lines, "\n") {
		herstory := make([]int, 0)
		for _, field := range strings.Fields(line) {
			herstory = append(herstory, stringToint(field))
		}
		herstories = append(herstories, herstory)
	}

	return herstories

}

func sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum+= value
	}
	return sum
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
