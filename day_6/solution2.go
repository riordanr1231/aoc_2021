package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

const Days = 256

//go:embed input
var input string
var Input []int

var Map map[int]int

func init() {
	InitInput()
	InitMap()
}

func InitInput() {
	for _, tok := range strings.Split(input, ",") {
		number, _ := strconv.Atoi(tok)
		Input = append(Input, number)
	}
}

func InitMap() {
	Map = make(map[int]int)

	for _, i := range Input {
		Map[i] += 1
	}
}

func PassDay() {
	newMap := make(map[int]int)

	for k, v := range Map {
		newKey := k - 1
		if newKey < 0 {
			newKey = 6
			newMap[8] = v
		}

		newMap[newKey] += v
	}

	Map = newMap
}

func SumValues() int {
	sum := 0
	for _, v := range Map {
		sum += v
	}
	return sum
}

func main() {
	for i := 0; i < Days; i++ {
		PassDay()
	}

	fmt.Println(SumValues())
}
