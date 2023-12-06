package main

import (
	"fmt"
	"math"
)

func main(){
	//part 1
	fmt.Println(multiply(calculateRaces(readInput())))
	//part 2
	fmt.Println(quadraticRace(readInput()[0]))
	fmt.Println(quadraticRace(readInputPart2()))



}

type (
	Race struct {
		Time      int
		Record int
	}
)

func quadraticRace(race Race) int {
	// -b +- sqrt of b^2 -4ac
	// 2a
	// is a just 1?
	time := float64(race.Time)
	record := float64(race.Record)
	sqrt := math.Sqrt(time*time - (record * 4))
	fmt.Println(sqrt)
	//low needs to round up, high rounds down
	low := int(math.Ceil((time - sqrt) / 2))
	high := int(math.Floor((time + sqrt) / 2))
	return high - low + 1
}
func calculateRaces (races []Race) []int {
	wins := make([]int, len(races))
	for i, race := range races {
		wins[i] = findWinCount(race)
	}
	return wins
}
func findWinCount(race Race) int{
	wins := 0
	for i:= 0; i < race.Time; i++ {
		timeRemaining := race.Time - i
		if timeRemaining * i > race.Record {
			wins+=1
		}
	}
	return wins

}

func multiply(numbers []int) int {
	total := 1
	for _, value := range numbers {
		total*= value
	}
	return total
}

func readInput() []Race {
//Time:        47     84     74     67
//Distance:   207   1394   1209   1014
	//not really worth reading progamitcally is it?
	return []Race{{47, 207}, {84, 1394}, {74, 1209}, {67, 1014}}
}
func readInputPart2() Race{
	return Race{47847467, 207139412091014}
}
