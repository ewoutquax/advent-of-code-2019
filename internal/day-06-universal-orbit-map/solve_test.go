package day06universalorbitmap_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/internal/day-06-universal-orbit-map"
	"github.com/stretchr/testify/assert"
)

func TestChecksumOrbits(t *testing.T) {
	lines := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}

	planets := ParseInput(lines)

	assert.Equal(t, 42, ChecksumOrbits(planets))
}

func TestNrStepsFromPlanetToPlanet(t *testing.T) {
	lines := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}

	planets := ParseInput(lines)

	assert.Equal(t, 4, MinNrStepsFromPlanetToPlanet(planets, "YOU", "SAN"))
}
