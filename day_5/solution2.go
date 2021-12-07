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

const InputSize = 500
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
	Input = make([][][]int, InputSize)
	for i := 0; i < InputSize; i++ {
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

func Print() {
	for _, m := range Matrix {
		fmt.Println(m)
	}
}

func Plot() {
	for _, in := range Input {
		if in[0][0] != in[1][0] && in[0][1] != in[1][1] {
			// plot diagonal
			inc_a := 1
			if in[0][0] > in[1][0] {
				inc_a = -1
			}

			inc_b := 1
			if in[0][1] > in[1][1] {
				inc_b = -1
			}

			i := in[0][0]
			j := in[0][1]
			for i != in[1][0] && j != in[1][1] {
				Matrix[j][i] += 1
				i += inc_a
				j += inc_b
			}

			Matrix[in[1][1]][in[1][0]] += 1

		} else if in[0][0] != in[1][0] {
			// walk first
			min := Min(in[0][0], in[1][0])
			max := Max(in[0][0], in[1][0])

			for i := min; i <= max; i++ {
				Matrix[in[0][1]][i] += 1
			}
		} else {
			// walk second
			min := Min(in[0][1], in[1][1])
			max := Max(in[0][1], in[1][1])

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

	return score
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
