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
var Lines []string

type Point struct {
	x int
	y int
}

type Sheet struct {
	Points []Point
}

var S Sheet

type Fold struct {
	Axis  string
	Level int
}

var Folds []Fold

func init() {
	InitInput()
	InitPoints()
	InitFolds()
}

func InitInput() {
	input = strings.TrimSuffix(input, "\n")
	Lines = strings.Split(input, "\n")
}

func InitPoints() {
	re := regexp.MustCompile(`^(\d+),(\d+)\s*$`)
	for _, l := range Lines {
		match := re.FindStringSubmatch(l)
		if match != nil {
			n1, _ := strconv.Atoi(match[1])
			n2, _ := strconv.Atoi(match[2])

			S.Points = append(S.Points, Point{n1, n2})
		}
	}
}

func InitFolds() {
	re := regexp.MustCompile(`^fold along (\w)=(\d+)$`)
	for _, l := range Lines {
		match := re.FindStringSubmatch(l)
		if match != nil {
			n, _ := strconv.Atoi(match[2])
			Folds = append(Folds, Fold{match[1], n})
		}
	}
}

func (s *Sheet) Fold(f Fold) {
	if f.Axis == "y" {
		for i := range S.Points {
			if S.Points[i].y > f.Level {
				shift := f.Level - (S.Points[i].y - f.Level)
				S.Points[i].y = shift
			}
		}
	}

	if f.Axis == "x" {
		for i := range S.Points {
			if S.Points[i].x > f.Level {
				shift := f.Level - (S.Points[i].x - f.Level)
				S.Points[i].x = shift
			}
		}
	}
}

func (s *Sheet) CountUnique() int {
	h := make(map[string]bool)

	for _, p := range S.Points {
		h[fmt.Sprintf("%d|%d", p.x, p.y)] = true
	}

	return len(h)
}

func main() {
	S.Fold(Folds[0])
	count := S.CountUnique()
	fmt.Println(count)
}
