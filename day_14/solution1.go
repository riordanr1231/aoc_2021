package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input
var input string
var Lines []string

var Template []string
var Rules map[string]string

func init() {
	InitInput()
}

func InitInput() {
	input = strings.TrimSuffix(input, "\n")
	Lines = strings.Split(input, "\n")
	Rules = make(map[string]string)

	for _, l := range Lines {
		re := regexp.MustCompile(`^\w+$`)
		match := re.FindStringSubmatch(l)
		if match != nil {
			Template = strings.Split(match[0], "")
		}

		re = regexp.MustCompile(`^(\w+)\s*->\s*(\w+)$`)
		match = re.FindStringSubmatch(l)
		if match != nil {
			Rules[match[1]] = match[2]
		}
	}
}

func ApplyInsertion() {
	var newTemplate []string
	for i := 0; i < len(Template)-1; i++ {
		newTemplate = append(newTemplate, Template[i])

		key := Template[i] + Template[i+1]
		if val, ok := Rules[key]; ok {
			newTemplate = append(newTemplate, val)
		}
	}

	newTemplate = append(newTemplate, Template[len(Template)-1])
	Template = newTemplate
}

func ApplyInsertions(count int, debug bool) {
	for i := 0; i < count; i++ {
		ApplyInsertion()
		if debug {
			fmt.Println(Template)
		}
	}
}

func CountChars() (int, int) {
	min := 2147483647
	max := 0
	m := make(map[string]int)

	for _, t := range Template {
		m[t]++
	}

	for _, v := range m {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return min, max
}

func main() {
	ApplyInsertions(10, false)
	min, max := CountChars()
	fmt.Println(max - min)
}
