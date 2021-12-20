package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strings"
)

var Delim = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var DelimClose = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var DelimValues = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

//go:embed input
var input string
var Lines []string

func init() {
	InitInput()
}

func InitInput() {
	input = strings.TrimSuffix(input, "\n")
	Lines = strings.Split(input, "\n")
}

func Invert(s string) string {
	var b strings.Builder

	b.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		if inverse, ok := DelimClose[string(s[i])]; ok {
			b.WriteString(inverse)
		} else {
			os.Exit(1)
		}

	}

	return b.String()
}

func BuildCompletionString(line string) (string, bool) {
	var openChars string

	for i, l := range line {

		c := string(l)
		if i == 0 {
			openChars += c
			continue
		}

		var ok bool
		var inverseChar string
		if inverseChar, ok = Delim[c]; !ok {
			openChars += c
			continue
		}

		lastChar := string(openChars[len(openChars)-1])
		if lastChar != inverseChar {

			return Invert(openChars), false
		}

		openChars = openChars[:len(openChars)-1]

	}

	return Invert(openChars), true
}

func Score(s string) int {
	var score int

	for _, ss := range s {
		if val, ok := DelimValues[string(ss)]; ok {
			score *= 5
			score += val
		} else {
			os.Exit(1)
		}
	}

	return score
}

func main() {
	var scores []int

	for _, l := range Lines {
		if s, valid := BuildCompletionString(l); valid {
			scores = append(scores, Score(s))
		}
	}

	mid := (len(scores) - 1) / 2
	sort.Ints(scores)
	fmt.Println(scores[mid])
}
