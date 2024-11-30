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

func getValFromMode(val int, program []int, mode int) int {
	switch mode {
	case 0:
		return program[val]
	case 1:
		return val
	default:
		log.Fatalf("Invalid parameter mode: %d\n", mode)
	}
	return -1
}

func storeValFromMode(idx int, val int, program []int, mode int) {
	switch mode {
	case 0:
		program[idx] = val
	default:
		log.Fatalf("Invalid parameter mode for storing: %d\n", mode)
	}
}

func add(ip int, program []int, p1Mode int, p2Mode int, p3Mode int) int {
	num1 := getValFromMode(program[ip+1], program, p1Mode)
	num2 := getValFromMode(program[ip+2], program, p2Mode)
	storeValFromMode(program[ip+3], num1+num2, program, p3Mode)
	return 4
}

func mult(ip int, program []int, p1Mode int, p2Mode int, p3Mode int) int {
	num1 := getValFromMode(program[ip+1], program, p1Mode)
	num2 := getValFromMode(program[ip+2], program, p2Mode)
	storeValFromMode(program[ip+3], num1*num2, program, p3Mode)
	return 4
}

func input(ip int, program []int, inp int, p1Mode int) int {
	storeValFromMode(program[ip+1], inp, program, p1Mode)
	return 2
}

func output(ip int, program []int, p1Mode int) (int, int) {
	out := getValFromMode(program[ip+1], program, p1Mode)
	return 2, out
}

func execute(program []int, inp int) int {
	ip := 0
	var outVal int
	var opcode int
	for opcode != 99 {
		code := program[ip]
		opcode = code % 100
		code /= 100
		p1Mode := code % 10
		code /= 10
		p2Mode := code % 10
		code /= 10
		p3Mode := code % 10
		switch opcode {
		case 1:
			ip += add(ip, program, p1Mode, p2Mode, p3Mode)
		case 2:
			ip += mult(ip, program, p1Mode, p2Mode, p3Mode)
		case 3:
			ip += input(ip, program, inp, p1Mode)
		case 4:
			var dip int
			dip, outVal = output(ip, program, p1Mode)
			ip += dip
		case 99:
			break
		default:
			log.Fatalf("Invalid opcode %d at position %d\n", program[ip], ip)
		}
	}

	return outVal
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	program := utils.Map(strings.Split(s, ","), utils.HandledAtoi)

	return execute(program, 1)
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
