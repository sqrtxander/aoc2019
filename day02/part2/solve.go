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

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			pc := intcode.NewIntcodeComputer(slices.Clone(program))
			pc.SetNounVerb(noun, verb)
			pc.Execute()
			if pc.Memory[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return -1
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
