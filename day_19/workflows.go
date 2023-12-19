package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main(){
	//part 1
	parts, workflows := parseLines(readInput())
	part1Accepted, _ := sortTools(parts, workflows)
	fmt.Println(sum(part1Accepted))

	//part 2
	combinations := distinctCombinations("in", Part{1,1,1,1}, Part{4000,4000,4000,4000}, workflows)
	fmt.Println(combinations)
}

func distinctCombinations(in string, min, max Part, workflows map[string][]Rule) int {
	if in == "R" {
		return 0
	}
	if in == "A" {
		return (max.x - min.x + 1) * (max.m - min.m + 1) * (max.a - min.a + 1) * (max.s - min.s + 1)
	}
	totalCombinations := 0
	//for each rule, split min/max send successful min/max recursive
	//for fail, min/max recursive
	for _, rule := range workflows[in] {
		switch rule.comparison {
		case "<":
			newMin, newMax := min, max
			//if min <, follow destination
			if min.getValue(rule.concerns) < rule.value {
				newMax.setValue(rule.concerns, rule.value - 1)
				totalCombinations += distinctCombinations(rule.destination, newMin, newMax, workflows)
			}
			//adjust the min as if we failed, check rest of the workflows rules
			if max.getValue(rule.concerns) >= rule.value {
				min.setValue(rule.concerns, rule.value)
			}
		case ">":
			newMin, newMax := min, max
			//if max >, follow destination
			if max.getValue(rule.concerns) > rule.value {
				newMin.setValue(rule.concerns, rule.value + 1)
				totalCombinations += distinctCombinations(rule.destination, newMin, newMax, workflows)
			}
			//adjust the max as if we failed, check rest of the workflows rules
			if min.getValue(rule.concerns) <= rule.value {
				max.setValue(rule.concerns, rule.value)
			}
		case "":
			totalCombinations += distinctCombinations(rule.destination, min, max, workflows)
		}
	}
	return totalCombinations
}

//dont really feel like the reflection is worth it now in part 2
func (this Part) getValue(property string) int {
	switch property {
	case "x":
		return this.x
	case "m":
		return this.m
	case "a":
		return this.a
	case "s":
		return this.s
	}
	panic("Property does not exist")
}

func (this *Part) setValue(property string, value int) {
	switch property {
	case "x":
		this.x = value
	case "m":
		this.m = value
	case "a":
		this.a = value
	case "s":
		this.s = value
	}
}

func sum(parts []Part) int {
	sum := 0
	for _, part := range parts {
		sum += part.x + part.m + part.a + part.s
	}
	return sum
}
func sortTools(parts []Part, workflows map[string][]Rule) (A, R []Part) {
	A = make([]Part, 0)
	//dont know if I'll need it but keeping anyways
	R = make([]Part, 0)
	for _, part := range parts {
		rules := workflows["in"]
		for true {
			var destination string
			for _, rule := range rules {
				if rule.check(part) {
					destination = rule.destination
					break
				}
			}
			//yay relearning that break breaks switches
			if destination == "A" {
				A = append(A, part)
				break
			}else if destination == "R" {
				R = append(R, part)
				break
			}else {
				rules = workflows[destination]
			}
		}
	}
	return A, R
}

func createCheck(field, condition string, value int) func(Part)bool {
	var check func(Part)bool
	if condition == "<" {
		check = func(part Part)bool {
			//yay reflect in go? alternative is switch for x,m,a,s, then switch on condition
			return int(reflect.Indirect(reflect.ValueOf(part)).FieldByName(field).Int()) < value
		}
	}else {
		check = func(part Part)bool {
			return int(reflect.Indirect(reflect.ValueOf(part)).FieldByName(field).Int()) > value
		}
	}
	return check
}
func createRule(input string) Rule {
	//x>1382:R OR smx
	if strings.ContainsAny(input, ":") {
		conditionAndDestination := strings.Split(input, ":")
		field := conditionAndDestination[0][:1]
		condition := conditionAndDestination[0][1:2]
		value := stringToInt(conditionAndDestination[0][2:])
		return Rule{value, field, condition,createCheck(field, condition, value), conditionAndDestination[1]}
	}else {
		return Rule{ check:func(Part)bool {return true}, destination: input}
	}
}
func parseRules(input string) map[string][]Rule {
	//remove rights, replace lefts with whitespace for easier parsing
	noBrackets := strings.ReplaceAll(strings.ReplaceAll(input, "{", " "), "}", "")
	//zt s>1817:R,m>2308:A,x>1382:R,smx
	workflows := map[string][]Rule{}
	for _, row := range strings.Split(noBrackets, "\n") {
		nameAndRules := strings.Fields(row)
		ruleStrings := strings.Split(nameAndRules[1], ",")
		rules := make([]Rule, len(ruleStrings))
		for i, ruleString := range ruleStrings {
			rules[i] = createRule(ruleString)
		}

		workflows[nameAndRules[0]] = rules
	}
	return workflows
}

func parseParts(input string) []Part {
	noBrackets := strings.ReplaceAll(strings.ReplaceAll(input, "{", ""), "}", "")
	parts := make([]Part, 0)
	for _, row := range strings.Split(noBrackets, "\n") {
		valuesWithLabels := strings.Split(row, ",")
		newPart := Part{}
		//in order of x,m,a,s, always starts with x=
		newPart.x = stringToInt(valuesWithLabels[0][2:])
		newPart.m = stringToInt(valuesWithLabels[1][2:])
		newPart.a = stringToInt(valuesWithLabels[2][2:])
		newPart.s = stringToInt(valuesWithLabels[3][2:])
		parts = append(parts, newPart)
	}
	return parts
}
func parseLines(lines string) ([]Part, map[string][]Rule) {
	rulesAndParts := strings.Split(lines, "\n\n")
	return parseParts(rulesAndParts[1]), parseRules(rulesAndParts[0])
}

type (
	Part struct {
		x,m,a,s int
	}
	Rule struct {
		value int
		concerns string
		comparison string
		check func(Part)bool
		destination string
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
