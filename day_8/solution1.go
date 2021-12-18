package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

//go:embed input
var input string

type Entry struct {
	Patterns []string
	Outputs  []string
}

var Entries []Entry

func init() {
	InitInput()
}

func InitInput() {
	re := regexp.MustCompile(`\n+$`)
	input = re.ReplaceAllString(input, "")
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		entry := Entry{}

		patternSection := true
		for _, s := range strings.Split(l, " ") {
			if s == "|" {
				patternSection = false
				continue
			}

			if patternSection {
				entry.Patterns = append(entry.Patterns, SortString(s))
			} else {
				entry.Outputs = append(entry.Outputs, SortString(s))
			}
		}
		Entries = append(Entries, entry)
	}
}

func SortString(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

func CountOutputDigits() int {
	count := 0
	for _, e := range Entries {
		for _, o := range e.Outputs {

			switch len(o) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return count
}

func main() {
	c := CountOutputDigits()
	fmt.Println(c)
}
