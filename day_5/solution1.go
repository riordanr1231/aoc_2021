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
var Input [][][]int
var Matrix [][]int
var Lines []string

const BoardSize = 1000

func init() {
	InitLines()
	InitMatrix()
	InitInput()
}

func InitLines() {
	Lines = strings.Split(input, "\n")
}

func InitMatrix() {
	Matrix = make([][]int, BoardSize)
	for i := 0; i < BoardSize; i++ {
		Matrix[i] = make([]int, BoardSize)
	}
}

func InitInput() {
	Input = make([][][]int, BoardSize)
	for i := 0; i < BoardSize; i++ {
		Input[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			Input[i][j] = make([]int, 2)
		}
	}

	for i, l := range Lines {
		re, _ := regexp.Compile(`^(\d+),(\d+)\s*->\s*(\d+),(\d+)`)
		match := re.FindStringSubmatch(l)

		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		n3, _ := strconv.Atoi(match[3])
		n4, _ := strconv.Atoi(match[4])

		Input[i][0] = []int{n1, n2}
		Input[i][1] = []int{n3, n4}
	}
}

func Plot() {
	for _, in := range Input {
		if in[0][0] != in[1][0] && in[0][1] != in[1][1] {
			continue
		} else if in[0][0] != in[1][0] {
			// walk first
			min := Min(in[0][0], in[1][0])
			max := Max(in[0][0], in[1][0])

			for i := min; i <= max; i++ {
				Matrix[in[0][1]][i] += 1
			}
		} else {
			min := Min(in[0][1], in[1][1])
			max := Max(in[0][1], in[1][1])

			// walk second
			for i := min; i <= max; i++ {
				Matrix[i][in[0][0]] += 1
			}
		}
	}
}

func OverlapScore() int {
	score := 0
	for _, m := range Matrix {
		for _, mm := range m {
			if mm > 1 {
				score++
			}
		}
	}

	return score - 1
}

func Min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	Plot()
	fmt.Println(OverlapScore())
}
