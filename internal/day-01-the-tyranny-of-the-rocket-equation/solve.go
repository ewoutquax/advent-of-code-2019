package day01thetyrannyoftherocketequation

import (
	"fmt"
	"math"

	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

func init() {
	register.Day("01a", solvePart1)
	register.Day("01b", solvePart2)
}

func solvePart1(inputFile string) {
	fuel := sumNeededFuel(utils.ReadFileAsLines(inputFile))
	fmt.Printf("Result of day-01 / part-1: %d\n", fuel)
}

func solvePart2(inputFile string) {
	fuel := sumNeededCumulativeFuel(utils.ReadFileAsLines(inputFile))
	fmt.Printf("Result of day-01 / part-2: %d\n", fuel)
}

func sumNeededFuel(lines []string) (total int) {
	var weight int

	for _, line := range lines {
		weight = utils.ConvStrToI(line)
		total += CalculateFuel(weight)
	}

	return
}

func sumNeededCumulativeFuel(lines []string) (total int) {
	var weight int

	for _, line := range lines {
		weight = utils.ConvStrToI(line)
		total += CalculateCumulativeFuel(weight)
	}

	return
}

func CalculateFuel(weight int) (fuel int) {
	return int(math.Floor(float64(weight)/3)) - 2
}

func CalculateCumulativeFuel(weight int) int {
	fuel := int(math.Floor(float64(weight)/3)) - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + CalculateCumulativeFuel(fuel)
}
