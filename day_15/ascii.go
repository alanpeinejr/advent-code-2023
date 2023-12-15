package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//part 1
	fmt.Println(hashAll(parseLines(readInput())))
	//part 2
	boxes := createBoxes()
	for _, v := range parseFields(parseLines(readInput())) {
		boxes = v.follow(boxes)
	}
	fmt.Println(sumFocusPower(boxes))

}
func sumFocusPower(boxes []Box) int{
	total := 0
	for _, box := range boxes {
		total += focusPower(box)
	}
	return total
}
func focusPower(box Box) int {
	total := 0
	for i, v := range box.Instructions{
		total += (box.ID +1) * (i + 1) * v.Power
	}
	return total
}

func (this Instruction) follow(boxes []Box) []Box{
	relevantBox := hash(this.Label)
	if this.Symbol == '-'{
		for i := 0; i < len(boxes[relevantBox].Instructions); i++ {
			if this.Label == boxes[relevantBox].Instructions[i].Label {
				boxes[relevantBox].Instructions = append(boxes[relevantBox].Instructions[:i], boxes[relevantBox].Instructions[i+1:]...)
				break
			}
		}
	}else{
		replaced := false
		for i := 0; i < len(boxes[relevantBox].Instructions); i++ {
			if this.Label == boxes[relevantBox].Instructions[i].Label {
				boxes[relevantBox].Instructions[i] = this
				replaced =  true
				break
			}
		}
		if !replaced {
			boxes[relevantBox].Instructions = append(boxes[relevantBox].Instructions, this)
		}
	}
	return boxes
}

func hashAll(input []string) int {
	total := 0
	for _, value := range input {
		total+= hash(value)
	}
	return total
}

func createBoxes() []Box {
	lens := make([]Box, 256)
	for i:= 0; i < 256; i++ {
		lens[i] = Box{make([]Instruction, 0), i}
	}
	return lens
}

func hash(input string) int {
	current:=0
	for _, char := range input{
		current += int(char)
		current *= 17
		current = current % 256
	}
	return current
}

func parseFields(inputs []string) []Instruction{
	instructions := make([]Instruction, len(inputs))
	for i, input := range inputs {
		if strings.ContainsAny(input, "123456789"){
			instructions[i] = Instruction{input[0:len(input)-2], '=',stringToInt(input[len(input)-1:])}
		}else{
			instructions[i] = Instruction{input[0:len(input)-1], '-', 0}
		}
	}
	return instructions
}
func parseLines(input string) []string {
	return strings.Split(input, ",")
}

func stringToInt(this string) int {
	value, _ := strconv.Atoi(this)
	return int(value)
}
type (
	Instruction struct {
		Label string
		Symbol rune
		Power int
	}
	Box struct {
		Instructions []Instruction
		ID int
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
