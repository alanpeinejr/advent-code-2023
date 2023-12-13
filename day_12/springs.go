package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//part 1
	springs, lengths := parseLines(readInput())
	fmt.Println(findTotalPossiblities(springs, lengths))

	//part 2
	unfoldedSprings := make([]string, len(springs))
	unfoldedLengths := make([][]int, len(lengths))
	for i, _ := range springs {
		unfoldedSprings[i] = unfoldSprings(springs[i])
		unfoldedLengths[i] = unfoldLengths(lengths[i])
	}
	fmt.Println(findTotalPossiblities(unfoldedSprings, unfoldedLengths))
}

func unfoldSprings(springs string) string {
	return springs + "?" + springs + "?" + springs + "?" + springs + "?" + springs
}
func unfoldLengths(lengths []int) []int {
	unfolded := make([]int, 0)
	unfolded= append(unfolded, lengths...)
	unfolded= append(unfolded, lengths...)
	unfolded= append(unfolded, lengths...)
	unfolded= append(unfolded, lengths...)
	unfolded= append(unfolded, lengths...)
	return unfolded
}

func findTotalPossiblities(springs []string, springLengths [][]int) int {
	total := 0
	for i, spring :=range springs {
		rowCache := createCache(spring, springLengths[i])
		total += iterateSprings(0, 0, spring, springLengths[i], rowCache)
	}
	return total
}
func createCache(springs string, lengths []int) (cache [][]int) {
	cache = make([][]int, 0)
	for range springs {
		groupCache := make([]int, len(lengths) + 1)
		for j, _ := range groupCache {
			groupCache[j] = -1
		}
		cache = append(cache, groupCache)
	}
	return cache
}

func iterateSprings(springIndex int, lengthIndex int, springs string, lengths []int, cache [][]int) int {
	//bases
	if springIndex >= len(springs) {
		if lengthIndex < len(lengths) {
			//deadend
			return 0
		}
		//at end of both, only 1
		return 1
	}

	//cache hit
	if cache[springIndex][lengthIndex] != -1 {
		return cache[springIndex][lengthIndex]
	}

	possibilities := 0
	if springs[springIndex] != '#' {
		possibilities += iterateSprings(springIndex + 1, lengthIndex, springs, lengths, cache)
	}

	if lengthIndex < len(lengths) {
		//springLength matches, skip ahead
		if springPotentialMatchesLength(springIndex, lengths[lengthIndex], springs){
			possibilities += iterateSprings(springIndex + lengths[lengthIndex] + 1, lengthIndex + 1, springs, lengths, cache)
		}
	}

	cache[springIndex][lengthIndex] = possibilities
	return possibilities
}

func springPotentialMatchesLength(springIndex int, length int, springs string) bool {
	springLength := 0
	for i := springIndex; i < len(springs); i++ {
		if springLength > length || springs[i] == '.' || springLength == length && springs[i] == '?'{
			//if we don't: reach the end, have too many #s, hit a '.', or are at correct length AND hit a '?'
			break
		}
		springLength += 1
	}
	return springLength == length
}

func parseLines(input string) ([]string, [][]int){
	springs := make([]string, 0)
	lengths := make([][]int, 0)
    for _, line := range strings.Split(input, "\n") {
		springsAndLengths := strings.Split(line, " ")
		springs = append(springs, springsAndLengths[0])
		lengthStrings := strings.Split(springsAndLengths[1], ",")
		length := make([]int, len(lengthStrings))
		for i, lengthString := range lengthStrings {
			length[i] = stringToInt(lengthString)
		}
		lengths = append(lengths, length)
	}
 return springs, lengths
}

func stringToInt(this string) int {
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
