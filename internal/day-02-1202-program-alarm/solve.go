package day021202programalarm

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/intcoder"
	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

func init() {
	register.Day("02a", solvePart1)
	register.Day("02b", solvePart2)
}

func solvePart1(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)
	chars := strings.Split(line, ",")

	var sourceCode []int
	for _, char := range chars {
		number := utils.ConvStrToI(char)
		sourceCode = append(sourceCode, number)
	}

	intCoder := intcoder.Compile(sourceCode)
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
			intCoder := intcoder.Compile(sourceCode)
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
