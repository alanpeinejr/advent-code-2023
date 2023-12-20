package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//part 1
	fmt.Println(pushButton(1000, parseLines(readInput())))

	//part 2
	fmt.Println(pushButton(10000, parseLines(readInput())))}

func findTriggerModule(modules map[string]*Module) *Module{
	//whats the module that can send a high to our end?
	return modules["xm"]// only module in input is &xm -> rx
}

func pushButton(times int, modules map[string]*Module) int {
	//lazy way to increment signals w/o if
	lowHigh := map[bool]int{}
	//part 2 tracking
	trigger := findTriggerModule(modules)
	part2:= map[string]int{}

	for i := 0; i< times; i++ {
		//button counts
		lowHigh[false] +=1
		starterPulse := Pulse{modules["broadcaster"], modules["broadcaster"], false}
		queue := processPulse(starterPulse, modules)
		for len(queue) > 0 {
			current := queue[0]
			lowHigh[current.signal] +=1
			//for endpoint
			if current.destination == nil {
				queue = queue[1:]
				continue
			}
			//part 2 addition
			//would need to account for double cycling if one of these results is small/big enough
			if current.destination.label == trigger.label && current.signal {
				fmt.Println(current.source.label, i + 1)
				part2[current.source.label]=i+1
				if(len(part2) == len(trigger.upStream)){
					lcm := 1
					for _, v := range part2 {
						lcm*=v
					}
					return lcm
				}
			}
			//end part 2
			queue = append(queue[1:], processPulse(current, modules)...)
		}
	}
	fmt.Println(lowHigh[false], lowHigh[true])
	return lowHigh[false] * lowHigh[true]
}
func processPulse(pulse Pulse, modules map[string]*Module) []Pulse {
	if pulse.signal && pulse.destination.class == "%"{
		//flipflops ignore highs
		return []Pulse{}
	}

	unprocessedPulses := make([]Pulse, len(pulse.destination.downstream))
	processedPulseResult := pulse.destination.process(pulse.signal, pulse.destination, pulse.source.label)
	for i, label := range pulse.destination.downstream {
		unprocessedPulses[i] = Pulse{pulse.destination, modules[label], processedPulseResult}
	}
	return unprocessedPulses
}

func parseLines(lines string) (map[string]*Module) {
	modules := map[string]*Module{}
		for _, line := range strings.Split(lines, "\n") {
			switch line[:1] {
			case "%":
				label, destinations := parseLine(line[1:])
				modules[label] = createFlipflop(label, destinations)
			case "&":
				label, destinations := parseLine(line[1:])
				modules[label] = createConjunction(label, destinations)
			default:
				label, destinations := parseLine(line)
				modules[label] = createBroadcast(label, destinations)

			}
		}
		for this, conjuction := range modules {
			if conjuction.class == "&" {
				upstream := make([]string, 0)
				for sender, v1 := range modules {
					for _, label := range v1.downstream {
						if label == this {
							upstream = append(upstream, sender)
						}
					}
				}
				modules[this].upStream = upstream
				modules[this].upStreamMemory = make([]bool, len(upstream))
			}
		}

	return modules
}
func createBroadcast(label string, destinations []string) *Module{
	return &Module{label,
		"B",
		false,
		destinations,
		func(input bool, this *Module, source string) bool {return input},
		[]string{},
		[]bool{}}
}
func createFlipflop(label string, destinations []string) *Module{
	return &Module{label,
		"%",
		false,
		destinations,
		func(input bool, this *Module, source string) bool {
			if !input {
				this.state = !this.state
				return this.state
			}
			panic("Dont send it highs")
		},
		[]string{},
		[]bool{}}
}

func createConjunction(label string, destinations []string) *Module{
	return &Module{label,
		"&",
		true,
		destinations,
		func(input bool, this *Module, source string) bool {
			for i, label := range this.upStream {
				if source == label {
					this.upStreamMemory[i] = input
				}
			}
			for _, pulse := range this.upStreamMemory {
				//if any low, it returns high
				if !pulse {
					return true
				}
			}
			return false
		},
		[]string{},
		[]bool{}}
}
func parseLine(line string) (label string, downstream []string){
	labelAndDestinations := strings.Split(line, " -> ")
	return labelAndDestinations[0], strings.Split(labelAndDestinations[1], ", ")
}

type (
	Module struct {
		label string
		class string
		state bool
		downstream []string
		process func(bool, *Module, string)bool
		upStream []string
		upStreamMemory []bool
	}
	Pulse struct {
		source *Module
		destination *Module
		signal bool
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
