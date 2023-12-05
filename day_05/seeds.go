package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//part 1
	seeds, rangeMaps := parseInput(readInput())
	fmt.Println(findLowestLocation(seeds, rangeMaps))
	//part 2
	seedRange := make([]Seed, len(seeds)/2)
	for i, _ :=range seedRange {
		seedRange[i] = Seed{seeds[i*2], seeds[i*2+1]}
		values := make([]int, seedRange[i].Length)
		for j, _ := range values {
			values[j] = seedRange[i].Id + j
		}
		fmt.Println(findLowestLocation(values, rangeMaps))
	}

}

func intervalLocations(seeds []Seed, rangeMapArray [][]RangeMap) int {
	intervals := seeds[:]
	for _, rangeMaps := range rangeMapArray {
		//for every interval, for the rangeMap the contains start, split
		//update intervals
		//for every interval, rangeMap values
		newIntervals := make([]Seed, 0)
		for _, interval := range intervals {
			for _, rangeMap := range rangeMaps {
				containsBeginning, _ := rangeMap.Contains(interval.Id)
				//ooooooooh, the interval and go into a new range.
				containsEnd, _ := rangeMap.Contains(interval.Id + interval.Length)
				if containsBeginning || containsEnd{
					newIntervals = append(newIntervals , splitInterval( interval, rangeMap)...)
					continue
				}
			}
			newIntervals = append(newIntervals, interval)
		}
		intervals = newIntervals
		fmt.Println(len(intervals))
		// now change
		nextIteration := make([]Seed, len(intervals))
		for i, interval := range intervals {
			nextIteration[i] = Seed{findInRangeMap(intervals[i].Id, rangeMaps), interval.Length}
		}

		intervals = nextIteration
	}

	minLocation := intervals[0]
	for _, location := range intervals {
		if location.Id < minLocation.Id {
			minLocation = location
		}
	}
	return minLocation.Id
}
func splitInterval(seed Seed, rangeMap RangeMap) []Seed {
	if seed.Id >= rangeMap.SourceStart && seed.Id + seed.Length < rangeMap.SourceStart + rangeMap.Length{
		return []Seed{seed}
	}else {
		// if it overlaps end
		within:= (seed.Id + seed.Length ) - (rangeMap.SourceStart + rangeMap.Length)
		if within > 0 {
			// beginning overlap
			low := Seed{seed.Id, within}
			high := Seed{low.Id+within, seed.Length-within }
			return []Seed{low, high}
		}else {
			//ending overlap
			low := Seed{seed.Id, rangeMap.SourceStart - seed.Id }
			high := Seed{low.Id + low.Length, seed.Length - low.Length}
			return []Seed{low, high}
		}

	}
}

func findLowestLocation(seeds []int, rangeMaps [][]RangeMap) int {
	results := seeds[:]
	//for each seed, find its location
	for i, _ := range seeds {
		for _, rangeMap := range rangeMaps{
			results[i] = findInRangeMap(results[i], rangeMap)
		}
	}

	minLocation := results[0]
	for _, location := range results {
		if location < minLocation {
			minLocation = location
		}
	}
	return minLocation
}

func findInRangeMap(value int, rangeMap []RangeMap) int {
	for _, mapItem := range rangeMap {
		contains, destinationValue := mapItem.Contains(value)
		if contains {
			return destinationValue
		}
	}
	return value
}

func (this RangeMap) Contains(value int) (bool, int) {
	if value >= this.SourceStart && value < (this.SourceStart + this.Length) {
		return true, this.DestinationStart + (value - this.SourceStart)
	}
	return false, value
}

func parseInput(input string) ([]int, [][]RangeMap) {
	seeds := make([]int, 0)
	rangeMaps := make([][]RangeMap, 0)
	sections := strings.Split(input, "\n\n")

	//seeds
	seedStringArray, _ := strings.CutPrefix(sections[0], "seeds: ")
	for _, value := range strings.Split(seedStringArray, " ") {
		seeds = append(seeds, stringToint(value))
	}

	rangeMaps = append(rangeMaps, parseSection(sections[1], "seed-to-soil map:\n"))
	rangeMaps = append(rangeMaps, parseSection(sections[2], "soil-to-fertilizer map:\n"))
	rangeMaps = append(rangeMaps, parseSection(sections[3], "fertilizer-to-water map:\n"))
	rangeMaps = append(rangeMaps, parseSection(sections[4], "water-to-light map:\n"))
	rangeMaps = append(rangeMaps, parseSection(sections[5], "light-to-temperature map:\n"))
	rangeMaps = append(rangeMaps, parseSection(sections[6], "temperature-to-humidity map:\n"))
	rangeMaps = append(rangeMaps, parseSection(sections[7], "humidity-to-location map:\n"))
	//
	//seedsToSoil := parseSection(sections[1], "seed-to-soil map:\n")
	//soilToFertilizer := parseSection(sections[2], "soil-to-fertilizer map:\n")
	//fertilizerToWater := parseSection(sections[3], "fertilizer-to-water map:\n")
	//waterToLight := parseSection(sections[4], "water-to-light map:\n")
	//lightToTemp := parseSection(sections[5], "light-to-temperature map:\n")
	//tempToHumidity := parseSection(sections[6], "temperature-to-humidity map:\n")
	//humidityToLocation := parseSection(sections[7], "humidity-to-location map:\n")


	return seeds, rangeMaps
}
func parseSection(section string, headerText string) []RangeMap {
	values := make([]RangeMap, 0)
	sectionNoHeader, _ := strings.CutPrefix(section, headerText)
	sectionStringArray := strings.Split(sectionNoHeader, "\n")
	for _, value := range sectionStringArray {
		stringNumbers := strings.Split(value, " ")
		values = append(values, RangeMap{stringToint(stringNumbers[1]), stringToint(stringNumbers[0]), stringToint(stringNumbers[2])})
	}

	return values
}
type (
	RangeMap struct {
		SourceStart int
		DestinationStart int
		Length int
	}
	Seed struct {
		Id int
		Length int
	}
)
func stringToint(this string) int {
	value, _ := strconv.Atoi(this)
	return int(value)
}

func sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum+= value
	}
	return sum
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
