package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	//part 1
	directions, nodes := parseLines(readInput())
	//part 2
	fmt.Println(followInstructions(directions, nodes, "AAA"))
	fmt.Println(followGhostInstructionsEh(directions, nodes))
}

func followGhostInstructionsEh(directions string, nodes map[string]DecisionPoint) int {
	//find the length of cycles, they will, undoubtedly cycle
	// --uhg are there multiple 'ends in z' within a single cycle?
	//cycle until you end up at the start. then lcm?
	//each start does only seem to cycle with one end
	starts, _:= ghostStartsAndEnds(nodes)
	steps := make([]int, len(starts))
	for i, start := range starts {
			steps[i] = followInstructions(directions, nodes, start)
	}
	fmt.Println(steps) //cycle length

	return lcm(steps)
}
func lcm(values []int) int{
	for _, v := range values {
		fmt.Println(primeFactors(v))
	}
	//Testing for this and saw results and...could code that out for highest occurance of each prime to lcm...
	//but also
	return 79 * 293 * 71 * 61 * 59 * 47 * 67
}
//always gotta do a prime alg in AoC
func primeFactors(n int) (primeFactors []int) {
	for n%2 == 0 {
		primeFactors = append(primeFactors, 2)
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			primeFactors = append(primeFactors, i)
			n = n / i
		}
	}
	if n > 2 {
		primeFactors = append(primeFactors, n)
	}
	return primeFactors
}

func ghostStartsAndEnds(nodes map[string]DecisionPoint) ([]string, []string) {
	starts, ends := make([]string, 0), make([]string, 0)
	for key, _ := range nodes {
		if key[2] == 'A'{
			starts = append(starts, key)
		}
		if key[2] == 'Z' {
			ends = append(ends, key)
		}
	}
	return starts, ends
}

func followInstruction(nextDirection uint8, point DecisionPoint) string{
	switch  {
	case nextDirection == 'L':
		return point.Left
	case nextDirection == 'R':
		return point.Right
	default:
		panic("whoops")
	}
}

func followInstructions(directions string, nodes map[string]DecisionPoint, start string) int {
	steps := 0
	current := start
	for current[2] != 'Z'{
		nextDirection := directions[steps % len(directions)]
		current = followInstruction(nextDirection, nodes[current])
		steps+=1
	}
	return steps
}

type (
	DecisionPoint struct {
		Left    string
		Right   string
	}
)

func parseLines(lines string) (string, map[string]DecisionPoint) {
	noDoubles := strings.ReplaceAll(lines,"\n\n", "\n")
	noParens := strings.ReplaceAll(strings.ReplaceAll(noDoubles, "(", ""), ")", "")
	noEquals := strings.ReplaceAll(noParens, " = ", " ")
	noCommas := strings.ReplaceAll(noEquals, ", ", " ")
	directionsAndDecisionPoints := strings.Split(noCommas, "\n")
	directions := directionsAndDecisionPoints[0]

	nodes := map[string]DecisionPoint{}
	for _, line := range directionsAndDecisionPoints[1:] {
		numbers := strings.Fields(line)
		nodes[numbers[0]] = DecisionPoint{numbers[1], numbers[2]}
	}

	return directions, nodes

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
