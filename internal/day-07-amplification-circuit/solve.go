package day07amplificationcircuit

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/intcoder"
	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

func init() {
	register.Day("7a", solvePart1)
	register.Day("7b", solvePart2)
}

func solvePart1(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)

	var sourceCode []int
	for _, code := range strings.Split(line, ",") {
		sourceCode = append(sourceCode, utils.ConvStrToI(code))
	}

	fmt.Printf("Result of day-07 / part-1: %d\n", MaxThrusterSignal(sourceCode))
}

func solvePart2(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)

	var sourceCode []int
	for _, code := range strings.Split(line, ",") {
		sourceCode = append(sourceCode, utils.ConvStrToI(code))
	}

	fmt.Printf("Result of day-07 / part-2: %d\n", MaxLoopedThrusterSignalBySequence(sourceCode))
}

// Build a list of strings made up of the received numbers, with these numbers in each possible order
func Permutations(numbers []int) (permutations [][]int) {
	if len(numbers) == 1 {
		return [][]int{numbers}
	}

	for idx := 0; idx < len(numbers); idx++ {
		var mine int = numbers[idx]
		var rest []int = make([]int, len(numbers))

		copy(rest, numbers)
		rest = append(rest[:idx], numbers[idx+1:]...)

		for _, subperm := range Permutations(rest) {
			permutations = append(permutations, append([]int{mine}, subperm...))
		}
	}

	return
}

func MaxThrusterSignal(sourceCode []int) (max int) {
	for _, sequence := range Permutations([]int{0, 1, 2, 3, 4}) {
		signal := ThrusterSignalBySequence(sequence, sourceCode)
		if max < signal {
			max = signal
		}
	}

	return
}

func ThrusterSignalBySequence(sequence, sourceCode []int) int {
	var amplifiers []*intcoder.IntCoder

	for _, startCode := range sequence {
		var amplifier = intcoder.Compile(sourceCode)
		amplifier.Send(startCode)
		amplifiers = append(amplifiers, amplifier)

	}

	amplifiers[0].Send(0)
	amplifiers[1].Send(amplifiers[0].Receive())
	amplifiers[2].Send(amplifiers[1].Receive())
	amplifiers[3].Send(amplifiers[2].Receive())
	amplifiers[4].Send(amplifiers[3].Receive())

	return amplifiers[4].Receive()
}

func MaxLoopedThrusterSignalBySequence(sourceCode []int) (max int) {
	for _, sequence := range Permutations([]int{5, 6, 7, 8, 9}) {
		signal := LoopedThrusterSignalBySequence(sequence, sourceCode)
		if max < signal {
			max = signal
		}
	}

	return
}

func LoopedThrusterSignalBySequence(sequence, sourceCode []int) int {
	var amplifiers []*intcoder.IntCoder

	for _, startCode := range sequence {
		var amplifier = intcoder.Compile(sourceCode)
		amplifier.Send(startCode)
		amplifiers = append(amplifiers, amplifier)

	}

	var signal int = 0
	for amplifiers[4].Status() != "halted" {
		amplifiers[0].Send(signal)
		amplifiers[1].Send(amplifiers[0].Receive())
		amplifiers[2].Send(amplifiers[1].Receive())
		amplifiers[3].Send(amplifiers[2].Receive())
		amplifiers[4].Send(amplifiers[3].Receive())
		signal = amplifiers[4].Receive()
	}

	return signal
}
