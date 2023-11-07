package day08spaceimageformat

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

func init() {
	register.Day("8a", solvePart1)
	register.Day("8b", solvePart2)
}

func solvePart1(inputFile string) {
	input := utils.ReadFileAsLine(inputFile)

	fmt.Printf("Result of day-08 / part-1: %d\n", Checksum(input, 25, 6))
}

func solvePart2(inputFile string) {
	input := utils.ReadFileAsLine(inputFile)
	lines := Print(input, 25, 6)

	fmt.Printf("Result of day-08 / part-2:\n")
	for _, line := range lines {
		temp := line
		temp = strings.Replace(temp, "1", "#", -1)
		temp = strings.Replace(temp, "0", ".", -1)
		fmt.Printf("\t%s\n", temp)
	}

	fmt.Printf("\n")
}

func Checksum(input string, width, height int) (checksum int) {
	var minCount0 int = -1

	for _, layer := range inputToLayers(input, width, height) {
		count0 := strings.Count(layer, "0")
		if minCount0 == -1 || minCount0 > count0 {
			minCount0 = count0
			checksum = strings.Count(layer, "1") * strings.Count(layer, "2")
		}
	}

	return
}

func Print(input string, width, height int) (lines []string) {
	var layers []string = inputToLayers(input, width, height)

	for y := 0; y < height; y++ {
		var line string

		for x := 0; x < width; x++ {
			var chars string

			idxChar := y*width + x
			for _, layer := range layers {
				chars += string(layer[idxChar])
			}

			line += string(strings.Replace(chars, "2", "", -1)[0])
		}

		lines = append(lines, line)
	}

	return
}

func inputToLayers(input string, width, height int) (layers []string) {
	for offset := 0; offset < len(input); offset += width * height {
		layer := string(input[offset : offset+width*height])
		layers = append(layers, layer)
	}

	return
}
