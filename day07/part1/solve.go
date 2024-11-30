package main

import (
	"aoc2019/intcode"
	"aoc2019/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

func solve(s string) int {
	s = strings.TrimSpace(s)
	program := utils.Map(strings.Split(s, ","), utils.HandledAtoi)

	var permute func([]int, func([]int), int)
	permute = func(a []int, f func([]int), i int) {
		if i > len(a) {
			f(a)
			return
		}
		permute(a, f, i+1)
		for j := i + 1; j < len(a); j++ {
			a[i], a[j] = a[j], a[i]
			permute(a, f, i+1)
			a[i], a[j] = a[j], a[i]
		}
	}

	maxRet := 0
	permute([]int{0, 1, 2, 3, 4}, func(inputs []int) {
		ret := 0
		for _, input := range inputs {
			pc := intcode.NewIntcodeComputer(slices.Clone(program))
			pc.AddInputs(input, ret)
			ret = pc.Execute()
		}
		maxRet = max(maxRet, ret)
	}, 0)
	return maxRet
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
