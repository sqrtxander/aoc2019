package main

import (
	"aoc2019/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func isValid(pass int) bool {
	adjacent := false
	last := pass % 10
	pass /= 10
	for pass > 0 {
		curr := pass % 10
		pass /= 10
		if curr == last {
			adjacent = true
		}
		if curr > last {
			return false
		}
		last = curr
	}
	return adjacent
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	loStr, hiStr, _ := strings.Cut(s, "-")
	lo := utils.HandledAtoi(loStr)
	hi := utils.HandledAtoi(hiStr)
	count := 0
	for i := lo; i < hi; i++ {
		if isValid(i) {
			count++
		}
	}

	return count
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
