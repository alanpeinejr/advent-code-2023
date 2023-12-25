package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main(){
	//part 1
	//print(parseLines(readInput()))
	components := parseLines(readInput())
	//print2(components)
	//print3(readInput())
	//orderedOccurances := freq(components)
	//[dxtrrd jfzthz ppnxrn lfcnpm bxmxrn cfxvcb llzzck htfnpm jhvtfd hkjksp] 9861 8720 8517 8453 7953 8160 8095 7991 7751 7817
	//fmt.Println(components["pzl"].without(components, []string{}))
	//fmt.Println(components["pzl"].without(components, []string{"hfx", "jqt", "cmg"}))
	//fmt.Println(components["pzl"].without(components, []string{"qnr", "xhk", "lhk"}))
	//fmt.Println(components["jpt"].without(components, []string{}))
	//potentials := []string{"dxt", "rrd", "jfz", "ppn", "xrn", "lfc", "npm", "bxm", "cfx", "vcb", "thz", "htf", "jhv", "tfd", "hkj", "ksp"}
	//potentials := orderedOccurances[:50]
	//for i:= 0; i < len(potentials); i++{
	//	for j:= i + 1; j < len(potentials); j++ {
	//		for k := j + 1; k < len(potentials); k++ {
	//			length := components["jpt"].without(components, []string{potentials[i], potentials[j], potentials[k]})
	//			if length < 1571 {
	//				fmt.Printf("%s %s %s\n", potentials[i], potentials[j], potentials[k])
	//			}
	//			//fmt.Println(components["jpt"].without(components, []string{potentials[i], potentials[j], potentials[k]}))
	//
	//		}
	//	}
	//}
	fmt.Println(len(components))
	fmt.Println(components["dhj"].without(components, []string{"xqh", "khn", "mqb"}))//how big dhj side is excluding its own
	fmt.Println(components["dhj"].without(components, []string{"ssd", "nrs", "qlc"}))//how this side is exluding the other sides connectors



	fmt.Println()


	//part 2
	fmt.Println()

}

func print(components map[string]*Component) {
	for _, value := range components {
		fmt.Print(value.label,)
		for _, connections := range value.connections {
			fmt.Print(" ", connections.label)
		}
		fmt.Println()

	}
}

func print2(components map[string]*Component) {
	for _, value := range components {
		fmt.Printf("%s [label=\"%s\"] ;\n", value.label, value.label)
		for _, connections := range value.connections {
			fmt.Printf("%s -- %s ;\n", value.label, connections.label)
		}

	}
}
func print3(lines string) {
	uniques := map[string]bool{}
	for _, row := range strings.Split(lines, "\n") {
		labelAndConnections := strings.Split(row, ": ")
		label := labelAndConnections[0]
		uniques[label] = true
		connections := strings.Fields(labelAndConnections[1])
		for _, conLabel := range connections {
			uniques[conLabel] = true
			fmt.Printf("%s -- %s ;\n", label, conLabel)

		}
	}
	for key, _ := range  uniques {
		fmt.Printf("%s [label=\"%s\"] ;\n", key, key)
	}
}

func (this Component) without(components map[string]*Component, us []string) int {
	queue := []string{this.label}
	visited := map[string]bool{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited[current] = true
		for _, value := range components[current].connections {
			if _, exists := visited[value.label]; !exists && !slices.Contains(us, value.label) {
				queue = append(queue, value.label)
			}
		}
	}
	return len(visited)

}

func freq(components map[string]*Component) []string  {
	freq := map[string]int{}
	for key, _ := range components {
		queue := []string{key}
		visited := map[string]bool{}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			visited[current] = true
			for _, value := range components[current].connections {
				if _, exists := visited[value.label]; !exists {
					var edgeLabel string
					if current < value.label {
						edgeLabel = fmt.Sprintf("%s%s", current, value.label)
					}else {
						edgeLabel = fmt.Sprintf("%s%s", value.label, current)
					}
					freq[edgeLabel]+=1
					queue = append(queue, value.label)
				}
			}
		}
	}
	stringFreqs := []OrderableStringFreq{}
	for key, value := range freq {
		stringFreqs  = append(stringFreqs, OrderableStringFreq{key, value})
	}
	slices.SortFunc(stringFreqs, compare)

	splitFreqs := []string{}
	for _, value := range stringFreqs {
		a, b := value.label[:3], value.label[3:]
		if !slices.Contains(splitFreqs, a) {
			splitFreqs = append(splitFreqs, a)
		}
		if !slices.Contains(splitFreqs, b) {
			splitFreqs = append(splitFreqs, b)
		}
	}

	fmt.Println(splitFreqs[0])
	return splitFreqs

}

func parseLines(lines string) (map[string]*Component){
	components := map[string]*Component{}
	for _, row := range strings.Split(lines, "\n") {
		labelAndConnections := strings.Split(row, ": ")
		label := labelAndConnections[0]
		connections := strings.Fields(labelAndConnections[1])
		if _, exists := components[label]; !exists {
			components[label] = &Component{label, 0, make([]*Component, 0)}
		}
		for _, conLabel := range connections {
			if _, exists := components[conLabel]; !exists {
				components[conLabel] = &Component{conLabel, 0, make([]*Component, 0)}
			}
			components[label].connections = append(components[label].connections, components[conLabel])
			components[conLabel].connections = append(components[conLabel].connections, components[label])
		}
	}
	return components
}

type (
	Component struct {
		label string
		seen int
		connections []*Component
	}
	OrderableStringFreq struct {
		label string
		freq int
	}
)

func compare(a, b OrderableStringFreq) int {
	return b.freq - a.freq
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
