package day17

import (
	"github.com/sogard-dev/advent-of-code-2024/utils"
)

type machine struct {
	instructionPointer int
	memory             []int
	output             []int
	registers          []int
}

type instruction = func(m *machine)

func newMachine(m []int, r []int) machine {
	return machine{
		memory:             m,
		registers:          r,
		output:             []int{},
		instructionPointer: 0,
	}
}

func parse(input string) machine {
	nums := utils.GetAllNumbers(input)
	r := nums[0:3]
	m := nums[3:]
	return newMachine(m, r)
}

func getLiteralArg(m *machine) int {
	return m.memory[m.instructionPointer+1]
}

func getComboArg(m *machine) int {
	v := m.memory[m.instructionPointer+1]
	if v < 4 {
		return v
	}
	if v < 7 {
		return m.registers[v-4]
	}
	panic("unexpected")
}

func adv(m *machine) {
	m.registers[0] = m.registers[0] / (utils.IntPow(2, getComboArg(m)))
	m.instructionPointer += 2
}

func bdv(m *machine) {
	m.registers[1] = m.registers[0] / (utils.IntPow(2, getComboArg(m)))
	m.instructionPointer += 2
}

func cdv(m *machine) {
	m.registers[2] = m.registers[0] / (utils.IntPow(2, getComboArg(m)))
	m.instructionPointer += 2
}

func bxl(m *machine) {
	m.registers[1] = m.registers[1] ^ getLiteralArg(m)
	m.instructionPointer += 2
}

func bst(m *machine) {
	m.registers[1] = getComboArg(m) % 8
	m.instructionPointer += 2
}

func jnz(m *machine) {
	v := m.registers[0]
	if v == 0 {
		m.instructionPointer += 2
	} else {
		m.instructionPointer = getLiteralArg(m)
	}
}

func bxc(m *machine) {
	m.registers[1] = m.registers[1] ^ m.registers[2]
	m.instructionPointer += 2
}

func out(m *machine) {
	m.output = append(m.output, getComboArg(m)%8)
	m.instructionPointer += 2
}

func part1(input string) []int {
	m := parse(input)
	runUntilHalt(&m)
	return m.output
}

func part2(input string) int {
	orig := parse(input)
	if runProgram(108107566389757, orig.memory) {
		return 108107566389757
	}
	return -1
}

func runProgram(i int, orig []int) bool {
	fIdx := 0
	A := i

	for A != 0 {
		B := A % 8
		B = B ^ 3
		C := A / utils.IntPow(2, B)
		B = B ^ C
		B = B ^ 3

		A = A / 8

		if fIdx < len(orig) && B%8 == orig[fIdx] {
			fIdx++
		} else {
			return false
		}
	}

	return fIdx == len(orig)
}

func runUntilHalt(m *machine) {
	instructions := map[int]instruction{
		0: adv,
		1: bxl,
		2: bst,
		3: jnz,
		4: bxc,
		5: out,
		6: bdv,
		7: cdv,
	}

	for m.instructionPointer < len(m.memory) {
		instructions[m.memory[m.instructionPointer]](m)
	}
}
