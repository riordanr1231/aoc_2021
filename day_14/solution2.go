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

var Pairs = make(map[string]int)
var Rules = make(map[string]string)
var Score = make(map[string]int)

func init() {
	InitInput()
}

func InitInput() {
	input = strings.TrimSuffix(input, "\n")
	Lines = strings.Split(input, "\n")

	for _, l := range Lines {
		re := regexp.MustCompile(`^\w+$`)
		match := re.FindStringSubmatch(l)
		if match != nil {
			m := match[0]
			for i := 0; i < len(m)-1; i++ {
				Score[m[i:i+1]]++
				Pairs[m[i:i+2]]++
			}
			Score[m[len(m)-1:len(m)]]++
		}

		re = regexp.MustCompile(`^(\w+)\s*->\s*(\w+)$`)
		match = re.FindStringSubmatch(l)
		if match != nil {
			Rules[match[1]] = match[2]
		}
	}
}

func ApplyInsertion() {
	newPairs := make(map[string]int)
	for p, vv := range Pairs {

		if vv < 1 {
			continue
		}

		if v, ok := Rules[p]; ok {
			Score[v] += vv
			newPairs[p[0:1]+v] += vv
			newPairs[v+p[1:2]] += vv
			newPairs[p] -= vv
		}
	}

	for k, v := range newPairs {
		Pairs[k] += v
	}
}

func ApplyInsertions(count int) {
	for i := 0; i < count; i++ {
		ApplyInsertion()
	}
}

func CountChars() (int, int) {
	min := 9999999999999
	max := 0

	for _, v := range Score {
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
	ApplyInsertions(40)
	min, max := CountChars()
	fmt.Println(max - min)
}
