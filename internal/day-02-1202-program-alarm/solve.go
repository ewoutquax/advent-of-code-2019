package day021202programalarm

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

func init() {
	register.Day("2a", solvePart1)
	register.Day("2b", solvePart2)
}

func solvePart1(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)
	chars := strings.Split(line, ",")

	var sourceCode []int
	for _, char := range chars {
		number := utils.ConvStrToI(char)
		sourceCode = append(sourceCode, number)
	}

	intCoder := Compile(sourceCode)
	intCoder.Set(1, 12)
	intCoder.Set(2, 2)
	intCoder.Run()

	fmt.Printf("Result of day-02 / part-1: %d\n", intCoder.Result()[0])
}

func solvePart2(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)
	chars := strings.Split(line, ",")

	var sourceCode []int
	for _, char := range chars {
		number := utils.ConvStrToI(char)
		sourceCode = append(sourceCode, number)
	}

	fmt.Printf("Result of day-02 / part-2: %d\n", findTarget(19690720, sourceCode))
}

func findTarget(target int, sourceCode []int) int {
	for verb := 0; verb < 100; verb++ {
		for noun := 0; noun < 100; noun++ {
			intCoder := Compile(sourceCode)
			intCoder.Set(1, verb)
			intCoder.Set(2, noun)
			intCoder.Run()

			if intCoder.Result()[0] == 19690720 {
				return verb*100 + noun
			}
		}
	}

	panic("This should not happen")
}

type IntCoder struct {
	idxInstruction int
	sourceCode     map[int]int
}

func (i *IntCoder) Result() (sourceCode []int) {
	for idx := 0; idx < len(i.sourceCode); idx++ {
		sourceCode = append(sourceCode, i.sourceCode[idx])
	}

	return
}

func (i *IntCoder) Run() {
	var running bool = true

	for running {
		switch i.sourceCode[i.idxInstruction] {
		case 1:
			i.add()
			i.idxInstruction += 4
		case 2:
			i.multiply()
			i.idxInstruction += 4
		case 99:
			running = false
		}
	}
}

func (i *IntCoder) Set(index, instruction int) {
	i.sourceCode[index] = instruction
}

func (i *IntCoder) add() {
	var idxLeft int = i.sourceCode[i.idxInstruction+1]
	var idxRight int = i.sourceCode[i.idxInstruction+2]
	var idxTarget int = i.sourceCode[i.idxInstruction+3]

	i.sourceCode[idxTarget] = i.sourceCode[idxLeft] + i.sourceCode[idxRight]
}

func (i *IntCoder) multiply() {
	var idxLeft int = i.sourceCode[i.idxInstruction+1]
	var idxRight int = i.sourceCode[i.idxInstruction+2]
	var idxTarget int = i.sourceCode[i.idxInstruction+3]

	i.sourceCode[idxTarget] = i.sourceCode[idxLeft] * i.sourceCode[idxRight]
}

func Compile(sourceCode []int) *IntCoder {
	intCoder := &IntCoder{
		idxInstruction: 0,
		sourceCode:     make(map[int]int, len(sourceCode)),
	}

	for idx, instruction := range sourceCode {
		intCoder.sourceCode[idx] = instruction
	}

	return intCoder
}
