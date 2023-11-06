package day06universalorbitmap

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

type Planet struct {
	name string

	parent *Planet
	moons  []*Planet
}

func (p *Planet) CountOrbits() (count int) {
	if len(p.moons) == 0 {
		return 0
	}

	for _, moon := range p.moons {
		count += 1 + moon.CountOrbits()
	}

	return
}

func init() {
	register.Day("6a", solvePart1)
	register.Day("6b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)

	planets := ParseInput(lines)

	fmt.Printf("Result of day-06 / part-1: %d\n", ChecksumOrbits(planets))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)

	planets := ParseInput(lines)

	fmt.Printf("Result of day-06 / part-2: %d\n", MinNrStepsFromPlanetToPlanet(planets, "YOU", "SAN"))
}

func MinNrStepsFromPlanetToPlanet(planets map[string]*Planet, nameFrom, nameTo string) (minDistance int) {
	var distances = make(map[string]int, len(planets))

	var current *Planet
	var from *Planet = planets[nameFrom]
	var to *Planet = planets[nameTo]

	current = from.parent
	for fromSteps := 0; current.parent != nil; fromSteps++ {
		distances[current.name] = fromSteps
		current = current.parent
	}

	current = to.parent
	for toSteps := 0; minDistance == 0; toSteps++ {
		if fromSteps, exists := distances[current.name]; exists {
			minDistance = fromSteps + toSteps
		}

		current = current.parent
	}

	return
}

func ChecksumOrbits(planets map[string]*Planet) (count int) {
	for _, planet := range planets {
		count += planet.CountOrbits()
	}

	return
}

func ParseInput(lines []string) map[string]*Planet {
	var planets = make(map[string]*Planet, len(lines)*2)

	// First, create all the planets
	for _, line := range lines {
		names := strings.Split(line, ")")

		for _, name := range names {
			if _, exists := planets[name]; !exists {
				planets[name] = &Planet{name: name}
			}
		}
	}

	// Then, read the file agian, but now build the links
	for _, line := range lines {
		names := strings.Split(line, ")")

		main := planets[names[0]]
		moon := planets[names[1]]
		moon.parent = main
		main.moons = append(main.moons, moon)
	}

	return planets
}
