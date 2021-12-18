package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string
var Input []int

var Map map[int]int

func init() {
	InitInput()
	InitMap()
}

func InitInput() {
	input = strings.TrimSuffix(input, "\n")

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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func AlignToPos(pos int) (cost int) {
	c := 0
	for k, v := range Map {
		c += Abs(k-pos) * v
	}

	return c
}

func main() {
	cost := 0
	for _, i := range Input {
		c := AlignToPos(i)
		if cost == 0 || cost > c {
			cost = c
		}
	}

	fmt.Println(cost)
}
