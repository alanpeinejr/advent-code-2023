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
	print2(components)
	//freq(components)
	//fmt.Println(components["pzl"].without(components, []string{}))
	//fmt.Println(components["pzl"].without(components, []string{"hfx", "jqt", "cmg"}))
	//fmt.Println(components["pzl"].without(components, []string{"qnr", "xhk", "lhk"}))
	//fmt.Println(components["jpt"].without(components, []string{}))
	//fmt.Println(components["jpt"].without(components, []string{"rrd", "thz", "xrn"}))


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

func freq(components map[string]*Component)  {
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
	largest := []string{"vtrdjb", "vtrdjb", "vtrdjb"}
	for key, value := range freq {
		if value > freq[largest[0]]{
			largest[2] = largest[1]
			largest[1] = largest[0]
			largest[0] = key
		}else if value > freq[largest[1]] {
			largest[2] = largest[1]
			largest[1] = key
		}else if value > freq[largest[2]]{
			largest[2] = key
		}
		//fmt.Println(key, value)
	}
	fmt.Println(largest, freq[largest[0]], freq[largest[1]], freq[largest[2]])

	//return len(visited)

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
)


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
