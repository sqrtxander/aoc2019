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
	adj := false
	groupOf2 := false
	last2 := pass % 10
	pass /= 10
	last := pass % 10
	pass /= 10
	var curr int
	if last == last2 && pass%10 != last {
		groupOf2 = true
	}
	for {
		curr = pass % 10
		pass /= 10
		if !adj && curr != last && last == last2 {
			groupOf2 = true
		}
		adj = last == last2
		if curr > last || last > last2 {
			return false
		}
		if pass == 0 {
			break
		}
		last2 = last
		last = curr
	}
	if curr == last && last != last2 {
		return true
	}
	return groupOf2
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
