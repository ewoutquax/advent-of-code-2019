package day04securecontainer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

func init() {
	register.Day("4a", solvePart1)
	register.Day("4b", solvePart2)
}

func solvePart1(inputFile string) {
	var from, to int

	line := utils.ReadFileAsLine(inputFile)
	parts := strings.Split(line, "-")
	from = utils.ConvStrToI(parts[0])
	to = utils.ConvStrToI(parts[1])

	fmt.Printf("Result of day-04 / part-1: %d\n", countValidPasswords(from, to))
}

func solvePart2(inputFile string) {
	var from, to int

	line := utils.ReadFileAsLine(inputFile)
	parts := strings.Split(line, "-")
	from = utils.ConvStrToI(parts[0])
	to = utils.ConvStrToI(parts[1])

	fmt.Printf("Result of day-04 / part-2: %d\n", countExtendedValidPasswords(from, to))
}

func countValidPasswords(from, to int) (count int) {
	for password := from; password <= to; password++ {
		if IsValidPassword(password) {
			count++
		}
	}

	return
}

func countExtendedValidPasswords(from, to int) (count int) {
	for password := from; password <= to; password++ {
		if IsExtendedValidPassword(password) {
			count++
		}
	}

	return
}

func IsValidPassword(password int) (valid bool) {
	valid = true

	var chars []string = strings.Split(strconv.Itoa(password), "")

	if valid {
		if !(chars[0] == chars[1] ||
			chars[1] == chars[2] ||
			chars[2] == chars[3] ||
			chars[3] == chars[4] ||
			chars[4] == chars[5]) {
			valid = false
		}
	}

	if valid {
		if !(chars[0] <= chars[1] &&
			chars[1] <= chars[2] &&
			chars[2] <= chars[3] &&
			chars[3] <= chars[4] &&
			chars[4] <= chars[5]) {
			valid = false
		}
	}

	return
}

func IsExtendedValidPassword(password int) (valid bool) {
	valid = IsValidPassword(password)

	if valid {
		var containsValidGroup bool = false
		var groups = make(map[string]int, 6)

		// Groups and counts used chars
		charPassword := strconv.Itoa(password)
		var chars []string = strings.Split(charPassword, "")

		for _, char := range chars {
			if _, exists := groups[char]; exists {
				groups[char]++
			} else {
				groups[char] = 1
			}
		}

		// For each char that appears exactly twice, check if these chars appear next to each other
		for char, count := range groups {
			if count == 2 {
				index1 := strings.Index(charPassword, char)
				index2 := strings.LastIndex(charPassword, char)

				if index2-index1 == 1 {
					containsValidGroup = true
				}
			}
		}

		valid = containsValidGroup
	}

	return
}
