package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string
var Input [][]int

func init() {
	InitInput()
}

func InitInput() {
	re := regexp.MustCompile(`\n+$`)
	input = re.ReplaceAllString(input, "")
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		var numbers []int
		for _, r := range l {
			h, _ := strconv.Atoi(string(r))
			numbers = append(numbers, h)
		}
		Input = append(Input, numbers)
	}
}

func MinIntArray(a []int) int {
	lowest := a[0]

	for _, aa := range a {
		if aa < lowest {
			lowest = aa
		}
	}

	return lowest
}

func GetAdjacentValues(i int, j int) []int {
	var adj []int

	if j > 0 {
		adj = append(adj, Input[i][j-1])
	}

	if j != len(Input[0])-1 {
		adj = append(adj, Input[i][j+1])
	}

	if i != 0 {
		adj = append(adj, Input[i-1][j])
	}

	if i != len(Input)-1 {
		adj = append(adj, Input[i+1][j])
	}

	return adj
}

func LowPoints() []int {
	var points []int

	for i, l := range Input {
		for j, number := range l {
			min := MinIntArray(GetAdjacentValues(i, j))
			if number < min {
				points = append(points, number)
			}
		}
	}

	return points
}

func Risk(a []int) int {
	risk := 0

	for _, aa := range a {
		risk += aa + 1
	}

	return risk
}

func main() {
	lp := LowPoints()
	r := Risk(lp)
	fmt.Println(r)
}
