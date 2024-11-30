package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	orbits := map[string]string{}

	for _, line := range lines {
		left, right, _ := strings.Cut(line, ")")
		orbits[right] = left
	}

	youMap := map[string]int{}
	you := orbits["YOU"]
	i := 0
	for you != "COM" {
		youMap[you] = i
		you = orbits[you]
		i++
	}
	youMap["COM"] = i

	sanMap := map[string]int{}
	san := orbits["SAN"]
	i = 0
	for san != "COM" {
		sanMap[san] = i
		san = orbits[san]
		i++
	}
	sanMap["COM"] = i

	minHops := math.MaxInt
	for planet := range youMap {
		if sanHops, ok := sanMap[planet]; ok {
			minHops = min(minHops, youMap[planet]+sanHops)
		}
	}
	return minHops
}

func main() {
	var inputPath string
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	} else {
		_, currentFilePath, _, _ := runtime.Caller(0)
		dir := filepath.Dir(currentFilePath)
		dir = filepath.Dir(dir)
		inputPath = filepath.Join(dir, "input.in")
	}
	contents, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Error reading file %s:\n%v\n", inputPath, err)
		return
	}
	fmt.Println(solve(string(contents)))
}
