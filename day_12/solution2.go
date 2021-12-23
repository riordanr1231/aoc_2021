package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
)

//go:embed input
var input string
var Lines []string

type Cave struct {
	Name        string
	Connections []*Cave
}

type CavePath struct {
	Path []*Cave
}

type CaveSystem struct {
	Caves []*Cave
	Map   map[string]*Cave
}

var CS CaveSystem

func init() {
	InitInput()
	InitCaveSystem()
}

func InitInput() {
	input = strings.TrimSuffix(input, "\n")
	Lines = strings.Split(input, "\n")
}

func InitCaveSystem() {
	CS.Map = make(map[string]*Cave)

	for _, l := range Lines {
		caves := strings.Split(l, "-")
		CS.AddCaves(caves[0], caves[1])
	}
}

func (cp *CavePath) Print() {
	var arr []string
	for _, path := range cp.Path {
		arr = append(arr, path.Name)
	}

	fmt.Println(strings.Join(arr, ","))
}

func (cp *CavePath) Contains(cave *Cave) bool {

	for _, p := range cp.Path {
		if p.Name == cave.Name {
			return true
		}
	}

	return false
}

func (cp *CavePath) ContainsAnyTwice() bool {

	caveMap := make(map[string]int)
	for _, cave := range cp.Path {
		if cave.IsLargeCave() {
			continue
		}

		caveMap[cave.Name]++
	}

	for _, v := range caveMap {
		if v > 1 {
			return true
		}
	}

	return false
}

func (cp *CavePath) Add(cave *Cave) {
	cp.Path = append(cp.Path, cave)
}

func (c *Cave) IsSmallCave() bool {
	return c.Name != strings.ToUpper(c.Name)
}

func (c *Cave) IsLargeCave() bool {
	return !c.IsSmallCave()
}

func (c *Cave) AddConnection(conn *Cave) {
	if c == conn {
		panic("Cave cannot connect to itself")
	}

	for _, c := range c.Connections {
		if c == conn {
			return
		}
	}

	c.Connections = append(c.Connections, conn)
}

func (cs *CaveSystem) AddCaves(name1 string, name2 string) {
	var cave1, cave2 *Cave

	if val, ok := cs.Map[name1]; ok {
		cave1 = val
	} else {
		cave1 = &Cave{Name: name1}
		cs.Map[name1] = cave1
	}

	if val, ok := cs.Map[name2]; ok {
		cave2 = val
	} else {
		cave2 = &Cave{Name: name2}
		cs.Map[name2] = cave2
	}

	cave1.AddConnection(cave2)
	cave2.AddConnection(cave1)
}

func PrintCavePaths(cps []CavePath) {
	for _, cp := range cps {
		cp.Print()
	}
}

func FindPaths(curr *Cave, ccp CavePath) []CavePath {
	var cp CavePath
	var paths []CavePath

	copier.Copy(&cp.Path, &ccp.Path)
	cp.Add(curr)

	if curr.Name == "end" {
		return append(paths, cp)
	}

	for _, conn := range curr.Connections {
		if conn.Name == "start" ||
			(conn.Name != "end" && conn.IsSmallCave() && cp.Contains(conn) && cp.ContainsAnyTwice()) {
			continue
		}

		fp := FindPaths(conn, cp)
		paths = append(paths, fp...)
	}

	return paths
}

func main() {
	var path CavePath
	var start = CS.Map["start"]
	cps := FindPaths(start, path)
	fmt.Println(len(cps))
}
