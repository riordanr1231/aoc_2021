package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
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

func CommonCount(a string, b string) int {
	count := 0
	for _, c := range strings.Split(a, "") {
		if strings.Contains(b, c) {
			count++
		}
	}

	return count
}

func ReverseMap(m map[int]string) map[string]int {
	mm := make(map[string]int, len(m))
	for k, v := range m {
		mm[v] = k
	}

	return mm
}

func (e *Entry) OutputValue() int {
	m := make(map[int]string)

	for _, p := range e.Patterns {
		switch len(p) {
		case 2:
			m[1] = p
		case 3:
			m[7] = p
		case 4:
			m[4] = p
		case 7:
			m[8] = p
		}
	}

	for _, p := range e.Patterns {
		if len(p) == 5 {
			switch {
			case CommonCount(p, m[4]) == 2:
				m[2] = p
			case CommonCount(p, m[1]) == 2:
				m[3] = p
			case CommonCount(p, m[1]) == 1:
				m[5] = p
			default:
				os.Exit(1)
			}
		} else if len(p) == 6 {
			switch {
			case CommonCount(p, m[1]) == 1:
				m[6] = p
			case CommonCount(p, m[4]) == 3:
				m[0] = p
			case CommonCount(p, m[4]) == 4:
				m[9] = p
			default:
				os.Exit(1)
			}
		}
	}

	var res string
	mm := ReverseMap(m)
	for _, o := range e.Outputs {
		res += strconv.Itoa(mm[o])
	}

	ires, _ := strconv.Atoi(res)

	return ires
}

func main() {
	total := 0
	for _, e := range Entries {
		total += e.OutputValue()
	}

	fmt.Println(total)
}
