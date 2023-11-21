package day03crossedwires

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

type Direction int8

const (
	Up Direction = iota + 1
	Left
	Right
	Down
)

type Path struct {
	direction Direction
	length    int
}

type Location struct {
	X       int
	Y       int
	nrSteps int
}

func (l *Location) toString() string {
	return fmt.Sprintf("%d,%d", l.X, l.Y)
}

func (l *Location) distance() int {
	return utils.Abs(l.X) + utils.Abs(l.Y)
}

type Wire struct {
	CurrentLocation Location `default:Location{0,0}`
	Locations       []Location
}

func (w *Wire) FollowInstruction(instruction string) {
	p := Path{}

	p.direction = map[byte]Direction{
		'U': Up,
		'L': Left,
		'R': Right,
		'D': Down,
	}[instruction[0]]
	p.length = utils.ConvStrToI(instruction[1:])

	for idx := p.length; idx > 0; idx-- {
		w.CurrentLocation.X += VectorX(p.direction)
		w.CurrentLocation.Y += VectorY(p.direction)
		w.CurrentLocation.nrSteps++

		w.Locations = append(w.Locations, w.CurrentLocation)
	}
}

func VectorX(d Direction) (vx int) {
	switch d {
	case Up, Down:
		vx = 0
	case Left:
		vx = -1
	case Right:
		vx = 1
	}

	return
}

func VectorY(d Direction) (vy int) {
	switch d {
	case Left, Right:
		vy = 0
	case Up:
		vy = 1
	case Down:
		vy = -1
	}

	return
}

func init() {
	register.Day("03a", solvePart1)
	register.Day("03b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)

	w0 := BuildWireByPath(lines[0])
	w1 := BuildWireByPath(lines[1])

	fmt.Printf("Result of day-03 / part-1: %d\n", FindClosestIntersection(w0, w1))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)

	w0 := BuildWireByPath(lines[0])
	w1 := BuildWireByPath(lines[1])

	fmt.Printf("Result of day-03 / part-2: %d\n", FindClosestIntersectionBySteps(w0, w1))
}

func BuildWireByPath(path string) (w Wire) {
	instructions := strings.Split(path, ",")
	for _, instruction := range instructions {
		w.FollowInstruction(instruction)
	}

	return w
}

func FindClosestIntersection(wire0, wire1 Wire) int {
	var intersections []Location
	var minDistance int

	// Store all position of wire-0
	w0Locations := make(map[string]Location, len(wire0.Locations))
	for _, location := range wire0.Locations {
		if _, exists := w0Locations[location.toString()]; !exists {
			w0Locations[location.toString()] = location
		}
	}

	// Travel through all position of wire-1, but store only the matches with wire-0
	for _, w1Location := range wire1.Locations {
		if _, exists := w0Locations[w1Location.toString()]; exists {
			intersections = append(intersections, w1Location)
		}
	}

	// Find the intersection with the closest manhattan distance
	minDistance = intersections[0].distance()
	for _, location := range intersections {
		if minDistance > location.distance() {
			minDistance = location.distance()
		}
	}

	return minDistance
}

func FindClosestIntersectionBySteps(wire0, wire1 Wire) (minNrSteps int) {
	minNrSteps = -1

	// Store all position of wire-0
	w0Locations := make(map[string]int, len(wire0.Locations))
	for _, location := range wire0.Locations {
		if _, exists := w0Locations[location.toString()]; exists {
			if w0Locations[location.toString()] > location.nrSteps {
				w0Locations[location.toString()] = location.nrSteps
			}
		} else {
			w0Locations[location.toString()] = location.nrSteps
		}
	}

	// Travel through all position of wire-1, but store only the matches with wire-0
	for _, w1Location := range wire1.Locations {
		if _, exists := w0Locations[w1Location.toString()]; exists {
			totalNrSteps := w0Locations[w1Location.toString()] + w1Location.nrSteps
			// fmt.Printf("Found intersection as %s, with %d and %d steps\n", w1Location.toString(), w0Locations[w1Location.toString()], w1Location.nrSteps)
			if minNrSteps == -1 || minNrSteps > totalNrSteps {
				minNrSteps = totalNrSteps
			}
		}
	}

	return
}
