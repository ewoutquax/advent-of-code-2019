package day05sunnywithachangeofastroids

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/intcoder"
	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

func init() {
	register.Day("05a", solvePart1)
	register.Day("05b", solvePart2)
	// register.Day("?b", solvePart2)
}

func solvePart1(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)

	var sourceCode []int
	for _, number := range strings.Split(line, ",") {
		sourceCode = append(sourceCode, utils.ConvStrToI(number))
	}

	intCoder := intcoder.Compile(sourceCode)
	intCoder.Send(1)

	fmt.Printf("Result of day-05 / part-1: %d\n", intCoder.Receive())
}

func solvePart2(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)

	var sourceCode []int
	for _, number := range strings.Split(line, ",") {
		sourceCode = append(sourceCode, utils.ConvStrToI(number))
	}

	intCoder := intcoder.Compile(sourceCode)
	intCoder.Send(5)

	fmt.Printf("Result of day-05 / part-2: %d\n", intCoder.Receive())

}
