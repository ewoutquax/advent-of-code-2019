package day16flawedfrequencytransmission

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

func init() {
	register.Day("16a", solvePart1)
	register.Day("16b", solvePart2)
}

func solvePart1(inputFile string) {
	input := utils.ReadFileAsLine(inputFile)

	fmt.Printf("Result of day-16 / part-1: %s\n", OutputAfterXPhases(input, 100))
}

func solvePart2(inputFile string) {
	input := utils.ReadFileAsLine(inputFile)

	fmt.Printf("Result of day-16 / part-2: %s\n", OutputAfterXPhasesByOffset(input))
}

func OutputAfterXPhases(input string, repeaters int) string {
	var intermediate string = input
	for repeat := repeaters; repeat > 0; repeat-- {
		intermediate = ApplyPhases(intermediate)
	}

	return string(intermediate[0:8])
}

// For some reason, a pattern becomes clear when looking at the intermediate results
// from the end:
// valueT1[idx] = valueT0[idx] + valueT1[idx+1]
//
// The values don't depend on any other value.
//
// So we just need to:
// - calculate each value in reverse, starting at the end, back to the position indicated by the first 7 numbers.
// - Keep the last 108 numbers
// - Do this 99 more times (100 times in total) on the remaining number
// - Return the first 8 numbers of the last iteration
func OutputAfterXPhasesByOffset(input string) (result string) {
	var offset int = utils.ConvStrToI(input[0:7])
	// Build the first iteration
	out := make([]int, 1)
	for idx := offset; idx < len(input)*10000; idx++ {
		tIdx := idx % len(input)
		out = append(out, utils.ConvStrToI(string(input[tIdx])))
	}

	// Let's get this over quickly
	for repeater := 100; repeater > 0; repeater-- {
		for idx := len(out) - 2; idx >= 0; idx-- {
			out[idx] = (out[idx] + out[idx+1]) % 10
		}
	}

	for idx := 1; idx < 9; idx++ {
		result += strconv.Itoa(out[idx])
	}

	return
}

func ApplyPhases(input string) (out string) {
	for repeater := 1; repeater <= len(input); repeater++ {
		out += ApplyPhase(input, repeater)
	}

	return
}

func ApplyPhase(input string, repeater int) string {
	var sum int
	for idx, char := range strings.Split(input, "") {
		sum += utils.ConvStrToI(char) * GetMultiplier(repeater, idx)
	}

	return strconv.Itoa(utils.Abs(sum) % 10)
}

func GetMultiplier(repeater, index int) int {
	adjustedIndex := ((index + 1) / repeater) % 4

	return []int{
		0,
		1,
		0,
		-1,
	}[adjustedIndex]
}
