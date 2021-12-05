package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	var pos, depth int64

	lines := strings.Split(input, "\n")
	for _, l := range lines {

		re, _ := regexp.Compile(`^(\w+)\s+(\d+)$`)
		match := re.FindStringSubmatch(l)
		if match == nil {
			os.Exit(1)
		}

		magnitude, _ := strconv.ParseInt(match[2], 10, 32)
		switch match[1] {
		case "forward":
			pos += magnitude
		case "up":
			depth -= magnitude
		case "down":
			depth += magnitude
		default:
			os.Exit(1)
		}
	}

	fmt.Println(pos, depth, pos*depth)
}
