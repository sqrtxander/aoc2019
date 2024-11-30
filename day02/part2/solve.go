package main

import (
	"aoc2019/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

func add(ip int, program []int) int {
	num1 := program[program[ip+1]]
	num2 := program[program[ip+2]]
	strIdx := program[ip+3]
	program[strIdx] = num1 + num2
	return 4
}

func mult(ip int, program []int) int {
	num1 := program[program[ip+1]]
	num2 := program[program[ip+2]]
	strIdx := program[ip+3]
	program[strIdx] = num1 * num2
	return 4
}

func execute(program []int, noun int, verb int) int {
	program[1] = noun
	program[2] = verb

	ip := 0
	for program[ip] != 99 {
		switch program[ip] {
		case 1:
			ip += add(ip, program)
		case 2:
			ip += mult(ip, program)
		default:
			log.Fatalf("Invalid opcode %d at position %d\n", program[ip], ip)
		}
	}

	return program[0]
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	program := utils.Map(strings.Split(s, ","), utils.HandledAtoi)

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if execute(slices.Clone(program), noun, verb) == 19690720 {
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
