package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//part 1
	fmt.Println(sum(parseInput(readInput())))
	//part 2

}

func sum(numberLines []string) int {
	 sum := 0
	 for _, v := range numberLines {
		 number :=findNumber(v)
		 sum += number
	 }
	return sum
}
func findNumber(numberLine string) int{
	firstIndex, lastIndex := len(numberLine), 0
	firstValue, lastValue := "",""
	for key, value := range numberStringsValues {
		var first = strings.Index(numberLine, key)
		var last = strings.LastIndex(numberLine, key)
		// exists if > -1
		if first > -1 {
			if first <= firstIndex {
				firstIndex = first
				firstValue = value
			}
			if last >= lastIndex {
				lastIndex = last
				lastValue = value
			}
		}

	}
	return stringToint(firstValue + lastValue)
}
var numberStringsValues =map[string]string{
	"1":"1",
	"2":"2",
	"3":"3",
	"4":"4",
	"5":"5",
	"6":"6",
	"7":"7",
	"8":"8",
	"9":"9",
	"0":"0",
	//comment out below if part 1
	"one":"1",
	"two":"2",
	"three":"3",
	"four":"4",
	"five":"5",
	"six":"6",
	"seven":"7",
	"eight":"8",
	"nine":"9",
	"zero":"0",
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
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