package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BoardSize = 5
const SolutionSize = 10

type Board struct {
	winner bool
	lines  [SolutionSize][]string
}

var Boards []Board

//go:embed input
var input string

func main() {
	sections := strings.Split(input, "\n\n")
	drawn := sections[0]

	for _, s := range sections[1:] {
		b := Board{}
		b.Create(s)
		Boards = append(Boards, b)
	}

	// mark all boards, one draw at a time
	for _, d := range strings.Split(drawn, ",") {
		for i := range Boards {
			losers := CountLosers()
			Boards[i].Mark(d)
			if Boards[i].Is_Winner() && losers == 1 && CountLosers() == 0 {
				dd, _ := strconv.Atoi(d)
				fmt.Println(Boards[i].Totals() * dd)
				os.Exit(0)
			}
		}
	}
}

func (b *Board) Create(board string) *Board {
	// init matrix
	matrix := make([][]string, BoardSize)
	for i := 0; i < BoardSize; i++ {
		matrix[i] = make([]string, BoardSize)
	}

	// populate matrix from board string
	for i, s := range strings.Split(board, "\n") {
		for j, ss := range strings.Fields(s) {
			matrix[i][j] = ss
		}
	}

	// horizontal
	b.lines[0] = matrix[0]
	b.lines[1] = matrix[1]
	b.lines[2] = matrix[2]
	b.lines[3] = matrix[3]
	b.lines[4] = matrix[4]

	// vertical
	b.lines[5] = []string{matrix[0][0], matrix[1][0], matrix[2][0], matrix[3][0], matrix[4][0]}
	b.lines[6] = []string{matrix[0][1], matrix[1][1], matrix[2][1], matrix[3][1], matrix[4][1]}
	b.lines[7] = []string{matrix[0][2], matrix[1][2], matrix[2][2], matrix[3][2], matrix[4][2]}
	b.lines[8] = []string{matrix[0][3], matrix[1][3], matrix[2][3], matrix[3][3], matrix[4][3]}
	b.lines[9] = []string{matrix[0][4], matrix[1][4], matrix[2][4], matrix[3][4], matrix[4][4]}

	return b
}

func (b *Board) Mark(number string) *Board {
	for i := 0; i < SolutionSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if b.lines[i][j] == number {
				b.lines[i][j] = "*"
			}
		}
	}
	return b
}

func (b *Board) Is_Winner() bool {
	for i := 0; i < SolutionSize; i++ {
		if b.lines[i][0] == "*" &&
			b.lines[i][1] == "*" &&
			b.lines[i][2] == "*" &&
			b.lines[i][3] == "*" &&
			b.lines[i][4] == "*" {
			b.winner = true
			return true
		}
	}

	return false
}

func (b *Board) Totals() int {
	totals := 0

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if b.lines[i][j] != "*" {
				n, _ := strconv.Atoi(b.lines[i][j])
				totals += n
			}
		}
	}

	return totals
}

func (b *Board) Print() {
	for i := 0; i < BoardSize; i++ {
		fmt.Println(b.lines[i])
	}
	return
}

func CountLosers() int {
	count := 0
	for _, b := range Boards {
		if !b.winner {
			count++
		}
	}

	return count
}
