package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
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

func GetAdjacentPoints(i int, j int) [][]int {
	var points [][]int

	if j > 0 {
		points = append(points, []int{i, j - 1})
	}

	if j != len(Input[0])-1 {
		points = append(points, []int{i, j + 1})
	}

	if i != 0 {
		points = append(points, []int{i - 1, j})
	}

	if i != len(Input)-1 {
		points = append(points, []int{i + 1, j})
	}

	return points
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

func LowPoints() [][]int {
	var points [][]int

	for i, l := range Input {
		for j, number := range l {
			min := MinIntArray(GetAdjacentValues(i, j))
			if number < min {
				points = append(points, []int{i, j})
			}
		}
	}

	return points
}

func BasinSize(p [][]int) int {
	h := make(map[string]bool)

	for _, pp := range p {
		h[fmt.Sprintf("%d|%d", pp[0], pp[1])] = true
	}

	return len(h)
}

func FindBasin(p []int) [][]int {
	var res [][]int
	res = append(res, p)

	val := Input[p[0]][p[1]]

	for _, ap := range GetAdjacentPoints(p[0], p[1]) {
		if Input[ap[0]][ap[1]] == 9 {
			continue
		}

		if Input[ap[0]][ap[1]] > val {
			r := FindBasin(ap)
			res = append(res, r...)
		}

	}

	return res
}

func main() {
	var basinSizes []int
	for _, lp := range LowPoints() {
		fb := FindBasin(lp)
		s := BasinSize(fb)
		basinSizes = append(basinSizes, s)
	}

	sort.Ints(basinSizes)
	l := len(basinSizes)

	fmt.Println(basinSizes[l-1] *
		basinSizes[l-2] *
		basinSizes[l-3])
}
