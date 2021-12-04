package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	var prev int64
	count := -1
	lines := strings.Split(input, "\n")

	for ndx := range lines {
		if ndx < 2 {
			continue
		}

		l1, _ := strconv.ParseInt(lines[ndx-2], 10, 32)
		l2, _ := strconv.ParseInt(lines[ndx-1], 10, 32)
		l3, _ := strconv.ParseInt(lines[ndx], 10, 32)
		curr := l1 + l2 + l3

		if curr > prev {
			count += 1
		}

		prev = curr
	}

	fmt.Println(count)
}
