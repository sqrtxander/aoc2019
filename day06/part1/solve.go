package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	orbits := map[string]string{}
	numOrbits := map[string]int{"COM": 1}

	for _, line := range lines {
		left, right, _ := strings.Cut(line, ")")
		orbits[right] = left
	}

	var countOrbits func(string) int
	countOrbits = func(planet string) int {
		if n, ok := numOrbits[planet]; ok {
			return n
		}
		numOrbits[planet] = 1 + countOrbits(orbits[planet])
		return numOrbits[planet]
	}

	total := 0
	for _, p := range orbits {
		total += countOrbits(p)
	}
	return total
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
