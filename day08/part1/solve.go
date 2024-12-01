package main

import (
	"aoc2019/utils"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func solve(s string, w int, h int) int {
	s = strings.TrimSpace(s)
	ppl := w * h
	layers := make([][]int, 0, len(s)/ppl)

	for i, char := range s {
		if i%ppl == 0 {
			layers = append(layers, make([]int, 0, ppl))
		}
		layers[len(layers)-1] = append(layers[len(layers)-1], utils.HandledAtoi(string(char)))
	}
	minZeros := math.MaxInt
	minZeroIdx := 0
	for i, layer := range layers {
		count := 0
		for _, num := range layer {
			if num == 0 {
				count++
			}
		}
		if count < minZeros {
			minZeroIdx = i
			minZeros = count
		}
	}

	oneCount := 0
	twoCount := 0
	for _, num := range layers[minZeroIdx] {
		if num == 1 {
			oneCount++
		} else if num == 2 {
			twoCount++
		}
	}

	return oneCount * twoCount
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
	fmt.Println(solve(string(contents), 25, 6))
}
