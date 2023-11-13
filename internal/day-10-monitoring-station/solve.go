package day10monitoringstation

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

const MYSELF int = 1

type Location struct {
	X int
	Y int
}

func (l *Location) Vector(l1 *Location) string {
	var direction string
	var vector float32

	if l.X > l1.X {
		direction = "r"
	} else {
		direction = "l" // to indicate the line moves to the left
	}

	if l.Y == l1.Y {
		vector = 0
	} else if l1.X == l.X {
		switch {
		case l1.Y > l.Y:
			direction = "u"
		case l1.Y < l.Y:
			direction = "d"
		}
	} else {
		vector = float32(l1.Y-l.Y) / float32(l1.X-l.X)
	}

	return fmt.Sprintf("%s/%f", direction, vector)
}

type Universe struct {
	Astroid []*Location
}

func (u *Universe) MaxCount() (max int) {
	for _, astroid := range u.Astroid {
		count := VisibleAstroids(u, *astroid)
		if max < count {
			max = count
		}
	}

	return
}

func init() {
	register.Day("10a", solvePart1)
	// register.Day("10b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	u := ParseInput(lines)

	fmt.Printf("Result of day-10 / part-1: %d\n", u.MaxCount())
}

func solvePart2(inputFile string) {
	fmt.Printf("Result of day-10 / part-2: %d\n", 0)
}

func VisibleAstroids(universe *Universe, location Location) (count int) {
	var uniqueVectors = make(map[string]bool, len(universe.Astroid))

	for _, loc := range universe.Astroid {
		if !(loc.X == location.X && loc.Y == location.Y) {
			vector := location.Vector(loc)
			uniqueVectors[vector] = true
		}
	}

	return len(uniqueVectors)
}

func ParseInput(lines []string) (u Universe) {
	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				loc := Location{X: x, Y: y}
				u.Astroid = append(u.Astroid, &loc)
			}
		}
	}

	return
}
