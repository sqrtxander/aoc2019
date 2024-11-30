package intcode

import (
	"log"
)

type IntcodeComputer struct {
	IP     int
	Memory []int
	Inputs []int
	Halted bool
}

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

func jnz(ip int, program []int, p1Mode int, p2Mode int) int {
	p1 := getValFromMode(program[ip+1], program, p1Mode)
	if p1 != 0 {
		return getValFromMode(program[ip+2], program, p2Mode) - ip
	}
	return 3
}

func jz(ip int, program []int, p1Mode int, p2Mode int) int {
	p1 := getValFromMode(program[ip+1], program, p1Mode)
	if p1 == 0 {
		return getValFromMode(program[ip+2], program, p2Mode) - ip
	}
	return 3
}

func less(ip int, program []int, p1Mode int, p2Mode int, p3Mode int) int {
	p1 := getValFromMode(program[ip+1], program, p1Mode)
	p2 := getValFromMode(program[ip+2], program, p2Mode)
	val := 0
	if p1 < p2 {
		val = 1
	}
	storeValFromMode(program[ip+3], val, program, p3Mode)
	return 4
}

func equal(ip int, program []int, p1Mode int, p2Mode int, p3Mode int) int {
	p1 := getValFromMode(program[ip+1], program, p1Mode)
	p2 := getValFromMode(program[ip+2], program, p2Mode)
	val := 0
	if p1 == p2 {
		val = 1
	}
	storeValFromMode(program[ip+3], val, program, p3Mode)
	return 4
}

func NewIntcodeComputer(memory []int) IntcodeComputer {
	return IntcodeComputer{
		IP:     0,
		Memory: memory,
		Inputs: []int{},
		Halted: false,
	}
}

func (pc *IntcodeComputer) AddInputs(inputs ...int) {
	pc.Inputs = append(pc.Inputs, inputs...)
}

func (pc *IntcodeComputer) SetNounVerb(noun int, verb int) {
	if len(pc.Memory) < 3 {
		log.Fatalf("Not enough room for noun and verb: length %d\n", len(pc.Memory))
	}
	pc.Memory[1] = noun
	pc.Memory[2] = verb
}

func (pc *IntcodeComputer) ExecuteUntilHalt() (outputs []int) {
	for !pc.Halted {
		ret := pc.Execute()
		outputs = append(outputs, ret)
	}
	return outputs[:len(outputs)-1]
}

func (pc *IntcodeComputer) Execute() int {
	var opcode int
	for opcode != 99 {
		code := pc.Memory[pc.IP]
		opcode = code % 100
		code /= 100
		p1Mode := code % 10
		code /= 10
		p2Mode := code % 10
		code /= 10
		p3Mode := code % 10
		switch opcode {
		case 1:
			pc.IP += add(pc.IP, pc.Memory, p1Mode, p2Mode, p3Mode)
		case 2:
			pc.IP += mult(pc.IP, pc.Memory, p1Mode, p2Mode, p3Mode)
		case 3:
			pc.IP += input(pc.IP, pc.Memory, pc.Inputs[0], p1Mode)
			pc.Inputs = pc.Inputs[1:]
		case 4:
			dip, outVal := output(pc.IP, pc.Memory, p1Mode)
			pc.IP += dip
			return outVal
		case 5:
			pc.IP += jnz(pc.IP, pc.Memory, p1Mode, p2Mode)
		case 6:
			pc.IP += jz(pc.IP, pc.Memory, p1Mode, p2Mode)
		case 7:
			pc.IP += less(pc.IP, pc.Memory, p1Mode, p2Mode, p3Mode)
		case 8:
			pc.IP += equal(pc.IP, pc.Memory, p1Mode, p2Mode, p3Mode)
		case 99:
			pc.Halted = true
			break
		default:
			log.Fatalf("Invalid opcode %d at position %d\n", pc.Memory[pc.IP], pc.IP)
		}
	}
	return -1
}
