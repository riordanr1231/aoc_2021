package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string
var Lines []string

type Octopus struct {
	Energy  int
	Flashed bool
}

type Octopuses struct {
	Members [][]Octopus
}

var flashes int
var octopuses Octopuses

func init() {
	InitInput()
	InitOctopuses()
}

func InitInput() {
	input = strings.TrimSuffix(input, "\n")
	Lines = strings.Split(input, "\n")
}

func InitOctopuses() {
	for _, l := range Lines {
		var opus []Octopus
		for _, b := range l {
			val, _ := strconv.Atoi(string(b))
			opus = append(opus, Octopus{val, false})

		}
		octopuses.Members = append(octopuses.Members, opus)
	}
}

func (o *Octopus) Reset() {
	o.Flashed = false
	if o.Energy > 9 {
		flashes++
		o.Energy = 0
	}
}

func (o *Octopus) IsEnergized() bool {
	if o.Energy > 9 && !o.Flashed {
		return true
	}

	return false
}

func (o *Octopus) Energize() {
	(*o).Energy += 1
}

func (op *Octopuses) Neighbors(i int, j int) []*Octopus {
	var nb []*Octopus

	// horizontal and vertical
	if j > 0 {
		nb = append(nb, &op.Members[i][j-1])
	}

	if j != len(op.Members[0])-1 {
		nb = append(nb, &op.Members[i][j+1])
	}

	if i != 0 {
		nb = append(nb, &op.Members[i-1][j])
	}

	if i != len(op.Members)-1 {
		nb = append(nb, &op.Members[i+1][j])
	}

	// diagonal
	if i > 0 && j > 0 {
		nb = append(nb, &op.Members[i-1][j-1])
	}

	if i != len(op.Members)-1 && j > 0 {
		nb = append(nb, &op.Members[i+1][j-1])
	}

	if i > 0 && j != len(op.Members[0])-1 {
		nb = append(nb, &op.Members[i-1][j+1])
	}

	if i != len(op.Members)-1 && j != len(op.Members[0])-1 {
		nb = append(nb, &op.Members[i+1][j+1])
	}

	return nb
}

func (op *Octopuses) Print() {
	for _, row := range op.Members {
		for _, col := range row {
			fmt.Print(col.Energy)
		}
		fmt.Println()
	}
}

func (op *Octopuses) Round() {
	for i := range op.Members {
		for j := range op.Members[i] {
			op.Members[i][j].Energize()
		}
	}

	complete := false
	for !complete {
		complete = true
		for i := range op.Members {
			for j := range op.Members[i] {
				if op.Members[i][j].IsEnergized() {
					complete = false
					op.Members[i][j].Flashed = true

					nb := op.Neighbors(i, j)
					for k := range nb {
						nb[k].Energize()
					}
				}
			}
		}
	}

	for i := range op.Members {
		for j := range op.Members[i] {
			op.Members[i][j].Reset()
		}
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		octopuses.Round()
	}

	fmt.Println(flashes)
}
