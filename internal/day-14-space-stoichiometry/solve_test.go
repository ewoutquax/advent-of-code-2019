package day14spacestoichiometry_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/internal/day-14-space-stoichiometry"
	"github.com/stretchr/testify/assert"
)

func TestParseLineWith1Dependency(t *testing.T) {
	assert := assert.New(t)
	line := "10 ORE => 10 A"

	recepy := ParseLine(line)

	assert.Equal("day14spacestoichiometry.Recepy", fmt.Sprintf("%s", reflect.TypeOf(recepy)))
	assert.Equal(10, recepy.Gives)
	assert.Equal("A", recepy.Product)

	assert.Equal(1, len(recepy.Requires))
	assert.Equal(10, recepy.Requires["ORE"])
}

func TestParseLineWith2Dependencies(t *testing.T) {
	assert := assert.New(t)

	line := "7 A, 1 E => 1 FUEL"
	recepy := ParseLine(line)

	assert.Equal("day14spacestoichiometry.Recepy", fmt.Sprintf("%s", reflect.TypeOf(recepy)))
	assert.Equal(1, recepy.Gives)
	assert.Equal("FUEL", recepy.Product)

	assert.Equal(2, len(recepy.Requires))
	assert.Equal(7, recepy.Requires["A"])
	assert.Equal(1, recepy.Requires["E"])
}

func TestBuildProgress(t *testing.T) {
	assert := assert.New(t)
	recepies := ParseInput(testInput())

	progress := BuildProgress(recepies)

	assert.Equal("*day14spacestoichiometry.Progress", fmt.Sprintf("%s", reflect.TypeOf(progress)))
	assert.Equal(1, len(progress.CurrentRequirements))
	assert.Equal(1, progress.CurrentRequirements["FUEL"])
	assert.Equal(7, len(progress.RemainingMaterials))
	assert.Equal(0, progress.RemainingMaterials["ORE"])
	assert.Equal(0, progress.RemainingMaterials["A"])
	assert.Equal(0, progress.RemainingMaterials["B"])
	assert.Equal(0, progress.RemainingMaterials["C"])
	assert.Equal(0, progress.RemainingMaterials["D"])
	assert.Equal(0, progress.RemainingMaterials["E"])
	assert.Equal(0, progress.RemainingMaterials["FUEL"])
}

func TestBreakDown(t *testing.T) {
	assert := assert.New(t)

	lines := []string{
		"2 ORE => 1 FUEL",
	}
	recepies := ParseInput(lines)

	progress := BuildProgress(recepies)

	progress.BreakDown(recepies)

	assert.Equal("*day14spacestoichiometry.Progress", fmt.Sprintf("%s", reflect.TypeOf(progress)))
	assert.Equal(0, len(progress.CurrentRequirements))
	assert.Equal(2, progress.AmountOre)
}

func TestBreakDownList(t *testing.T) {
	var recepies []Recepy
	recepies = ParseInput(testInput())

	progress := BuildProgress(recepies)

	for len(progress.CurrentRequirements) > 0 {
		progress.BreakDown(recepies)
	}

	assert.Equal(t, 31, progress.AmountOre)
}

func TestBreakDownList2(t *testing.T) {
	var recepies []Recepy

	recepies = ParseInput(testInput2())

	progress := BuildProgress(recepies)

	var doContinue bool = true
	for doContinue {
		progress.BreakDown(recepies)
		doContinue = len(progress.CurrentRequirements) > 0
	}

	assert.Equal(t, 13312, progress.AmountOre)
}

func TestMaxFuelByOre(t *testing.T) {
	recepies := ParseInput(testInput2())

	maxFuel := FindMaxFuelForAmountOre(recepies, 1000000000000)

	assert.Equal(t, 82892753, maxFuel)
}

func testInput() []string {
	return []string{
		"10 ORE => 10 A",
		"1 ORE => 1 B",
		"7 A, 1 B => 1 C",
		"7 A, 1 C => 1 D",
		"7 A, 1 D => 1 E",
		"7 A, 1 E => 1 FUEL",
	}
}

func testInput2() []string {
	return []string{
		"157 ORE => 5 NZVS",
		"165 ORE => 6 DCFZ",
		"44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL",
		"12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ",
		"179 ORE => 7 PSHF",
		"177 ORE => 5 HKGWZ",
		"7 DCFZ, 7 PSHF => 2 XJWVT",
		"165 ORE => 2 GPVTF",
		"3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT",
	}
}
