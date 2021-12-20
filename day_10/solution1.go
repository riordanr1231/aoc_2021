package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

var Delim = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var DelimValues = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
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

func IsLineValid(line string) (string, bool) {
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
			return c, false
		}

		openChars = openChars[:len(openChars)-1]

	}

	return "", true
}

func main() {
	score := 0
	for _, l := range Lines {
		if c, valid := IsLineValid(l); !valid {
			if val, ok := DelimValues[c]; ok {
				score += val

			} else {
				os.Exit(1)
			}
		}
	}

	fmt.Println(score)
}
