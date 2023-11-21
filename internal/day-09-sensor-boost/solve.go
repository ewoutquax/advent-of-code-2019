package day09sensorboost

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/intcoder"
	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

func init() {
	register.Day("09a", solvePart1)
	register.Day("09b", solvePart2)
}

func solvePart1(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)
	var sourceCode []int
	for _, code := range strings.Split(line, ",") {
		sourceCode = append(sourceCode, utils.ConvStrToI(code))
	}
	intCoder := intcoder.Compile(sourceCode)
	intCoder.Send(1)

	fmt.Printf("Result of day-09 / part-1: %d\n", intCoder.Receive())
}

func solvePart2(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)
	var sourceCode []int
	for _, code := range strings.Split(line, ",") {
		sourceCode = append(sourceCode, utils.ConvStrToI(code))
	}
	intCoder := intcoder.Compile(sourceCode)
	intCoder.Send(2)

	fmt.Printf("Result of day-09 / part-2: %d\n", intCoder.Receive())
}
