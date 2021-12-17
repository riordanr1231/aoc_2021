package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

const Days = 80

type Fish struct {
	Clock int
}

type School struct {
	Members []Fish
}

var school School

//go:embed input
var input string
var Input []int

func init() {
	InitInput()
	InitSchool()
}

func InitInput() {
	for _, tok := range strings.Split(input, ",") {
		number, _ := strconv.Atoi(tok)
		Input = append(Input, number)
	}
}

func InitSchool() {
	for _, i := range Input {
		f := Fish{Clock: i}
		school.Members = append(school.Members, f)
	}
}

func (f *Fish) PassDay() (spawn bool) {
	(*f).Clock -= 1
	if f.Clock < 0 {
		(*f).Clock = 6
		return true
	}
	return false
}

func (s *School) Spawn() {
	f := Fish{Clock: 8}
	school.Members = append(school.Members, f)
}

func (s *School) PassDay() {
	spawnCount := 0

	for i := range s.Members {
		if s.Members[i].PassDay() {
			spawnCount += 1
		}
	}

	for i := 0; i < spawnCount; i++ {
		s.Spawn()
	}
}

func (s *School) Length() int {
	return len(s.Members)
}

func main() {
	for i := 0; i < Days; i++ {
		school.PassDay()
	}

	fmt.Println(school.Length())
}
