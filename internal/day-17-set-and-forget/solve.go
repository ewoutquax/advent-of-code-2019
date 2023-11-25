package day17setandforget

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/intcoder"
	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

type Location struct {
	X int
	Y int
}

type Intersection struct {
	Location
}

func (i *Intersection) alignment() int {
	return i.X * i.Y
}

func init() {
	register.Day("17a", solvePart1)
	register.Day("17b", solvePart2)
}

func solvePart1(inputFile string) {
	var intCoder *intcoder.IntCoder = buildIntCoder(inputFile)
	fmt.Printf("Result of day-17 / part-1: %d\n", SumAlignmentParameters(getScaffolds(intCoder)))
}

func solvePart2(inputFile string) {
	var intCoder *intcoder.IntCoder = buildIntCoder(inputFile)
	drawScaffolds(intCoder)

	fmt.Printf("Result of day-17 / part-2: %d\n", 0)
}

func GenerateOptimizedParts(path string) (optimizedPaths []string) {
	var paths []string = strings.Split(path, ",")
	var totalLength int = len(paths)

	var lenPathA, lenPathB, lenPathC int

	for lenPathA = 1; lenPathA < totalLength-2; lenPathA++ {
		for lenPathB = 1; lenPathB < totalLength-lenPathA-1; lenPathB++ {
			lenPathC = totalLength - lenPathB - lenPathA
		}
	}
}

func SumAlignmentParameters(lines []string) (sum int) {
	for _, intersection := range FindIntersections(lines) {
		sum += intersection.alignment()
	}
	return
}

func FindIntersections(lines []string) (intersections []Intersection) {
	var locations = make(map[Location]bool, 0)

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				location := Location{X: x, Y: y}
				locations[location] = true

				if locationIsBelowIntersection(location, locations) {
					intersections = append(intersections, Intersection{Location{x, y - 1}})
				}
			}
		}
	}

	return
}

func locationIsBelowIntersection(location Location, locations map[Location]bool) bool {
	vectors := [][2]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{0, -2},
	}

	var debug bool = location.X == 2 && location.Y == 3

	if debug {
		fmt.Println("solve: locationIsBelowIntersection: start")
		fmt.Printf("current location: %v\n", location)
		fmt.Printf("locations: %v\n", locations)
	}

	for _, vector := range vectors {
		loc := Location{
			X: location.X + vector[0],
			Y: location.Y + vector[1],
		}

		if debug {
			fmt.Printf("solve: Testing location: %v\n", loc)
		}

		if _, exists := locations[loc]; !exists {
			return false
		}
	}

	return true
}

func drawScaffolds(intCoder *intcoder.IntCoder) {
	lines := getScaffolds(intCoder)

	for _, line := range lines {
		fmt.Println(line)
	}

	return
}

func getScaffolds(intCoder *intcoder.IntCoder) []string {
	var out string

	intCoder.Run()

	canContinue := true
	for canContinue {
		ascii := intCoder.Receive()

		if ascii == -1337 {
			canContinue = false
		} else {
			char := rune(ascii)
			out += fmt.Sprintf("%c", char)
		}
	}

	return strings.Split(out, "\n")
}

func buildIntCoder(inputFile string) *intcoder.IntCoder {
	input := utils.ReadFileAsLine(inputFile)

	var sourceCode []int
	for _, char := range strings.Split(input, ",") {
		sourceCode = append(sourceCode, utils.ConvStrToI(char))
	}
	return intcoder.Compile(sourceCode)
}
