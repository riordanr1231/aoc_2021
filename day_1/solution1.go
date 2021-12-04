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
	count := 0
	lines := strings.Split(input, "\n")

	for ndx := range lines {
		if ndx == 0 {
			continue
		}

		l1, _ := strconv.ParseInt(lines[ndx-1], 10, 32)
		l2, _ := strconv.ParseInt(lines[ndx], 10, 32)

		if l1 < l2 {
			count += 1
		}
	}

	fmt.Println(count)
}
