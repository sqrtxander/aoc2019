package intcode

import (
	"log"
)

type IntcodeComputer struct {
	IP      int
	relBase int
	Memory  map[int]int
	Inputs  []int
	Halted  bool
}

func (pc *IntcodeComputer) getValFromMode(val int, mode int) int {
	switch mode {
	case 0:
		return pc.Memory[val]
	case 1:
		return val
	case 2:
		return pc.Memory[val+pc.relBase]
	default:
		log.Fatalf("Invalid parameter mode: %d\n", mode)
	}
	return -1
}

func (pc *IntcodeComputer) storeValFromMode(idx int, val int, mode int) {
	var setIdx int
	switch mode {
	case 0:
		setIdx = idx
	case 2:
		setIdx = idx + pc.relBase
	default:
		log.Fatalf("Invalid parameter mode for storing: %d\n", mode)
	}
	if setIdx < 0 {
		log.Fatalf("Invalid memory address for storing: %d\n", setIdx)
	}
	pc.Memory[setIdx] = val
}

func (pc *IntcodeComputer) add(p1Mode int, p2Mode int, p3Mode int) {
	num1 := pc.getValFromMode(pc.Memory[pc.IP+1], p1Mode)
	num2 := pc.getValFromMode(pc.Memory[pc.IP+2], p2Mode)
	pc.storeValFromMode(pc.Memory[pc.IP+3], num1+num2, p3Mode)
	pc.IP += 4
}

func (pc *IntcodeComputer) mult(p1Mode int, p2Mode int, p3Mode int) {
	num1 := pc.getValFromMode(pc.Memory[pc.IP+1], p1Mode)
	num2 := pc.getValFromMode(pc.Memory[pc.IP+2], p2Mode)
	pc.storeValFromMode(pc.Memory[pc.IP+3], num1*num2, p3Mode)
	pc.IP += 4
}

func (pc *IntcodeComputer) input(p1Mode int) {
	pc.storeValFromMode(pc.Memory[pc.IP+1], pc.Inputs[0], p1Mode)
	pc.Inputs = pc.Inputs[1:]
	pc.IP += 2
}

func (pc *IntcodeComputer) output(p1Mode int) int {
	out := pc.getValFromMode(pc.Memory[pc.IP+1], p1Mode)
	pc.IP += 2
	return out
}

func (pc *IntcodeComputer) jnz(p1Mode int, p2Mode int) {
	p1 := pc.getValFromMode(pc.Memory[pc.IP+1], p1Mode)
	if p1 != 0 {
		pc.IP = pc.getValFromMode(pc.Memory[pc.IP+2], p2Mode)
	} else {
		pc.IP += 3
	}
}

func (pc *IntcodeComputer) jz(p1Mode int, p2Mode int) {
	p1 := pc.getValFromMode(pc.Memory[pc.IP+1], p1Mode)
	if p1 == 0 {
		pc.IP = pc.getValFromMode(pc.Memory[pc.IP+2], p2Mode)
	} else {
		pc.IP += 3
	}
}

func (pc *IntcodeComputer) less(p1Mode int, p2Mode int, p3Mode int) {
	p1 := pc.getValFromMode(pc.Memory[pc.IP+1], p1Mode)
	p2 := pc.getValFromMode(pc.Memory[pc.IP+2], p2Mode)
	val := 0
	if p1 < p2 {
		val = 1
	}
	pc.storeValFromMode(pc.Memory[pc.IP+3], val, p3Mode)
	pc.IP += 4
}

func (pc *IntcodeComputer) equal(p1Mode int, p2Mode int, p3Mode int) {
	p1 := pc.getValFromMode(pc.Memory[pc.IP+1], p1Mode)
	p2 := pc.getValFromMode(pc.Memory[pc.IP+2], p2Mode)
	val := 0
	if p1 == p2 {
		val = 1
	}
	pc.storeValFromMode(pc.Memory[pc.IP+3], val, p3Mode)
	pc.IP += 4
}

func (pc *IntcodeComputer) adjRelBase(p1Mode int) {
	pc.relBase += pc.getValFromMode(pc.Memory[pc.IP+1], p1Mode)
    pc.IP += 2
}

func NewIntcodeComputer(memory []int) IntcodeComputer {
	memoryMap := map[int]int{}
	for i, num := range memory {
		memoryMap[i] = num
	}
	return IntcodeComputer{
		IP:      0,
		relBase: 0,
		Memory:  memoryMap,
		Inputs:  []int{},
		Halted:  false,
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
			pc.add(p1Mode, p2Mode, p3Mode)
		case 2:
			pc.mult(p1Mode, p2Mode, p3Mode)
		case 3:
			pc.input(p1Mode)
		case 4:
			return pc.output(p1Mode)
		case 5:
			pc.jnz(p1Mode, p2Mode)
		case 6:
			pc.jz(p1Mode, p2Mode)
		case 7:
			pc.less(p1Mode, p2Mode, p3Mode)
		case 8:
			pc.equal(p1Mode, p2Mode, p3Mode)
        case 9:
            pc.adjRelBase(p1Mode)
		case 99:
			pc.Halted = true
			break
		default:
			log.Fatalf("Invalid opcode %d at position %d\n", pc.Memory[pc.IP], pc.IP)
		}
	}
	return -1
}
