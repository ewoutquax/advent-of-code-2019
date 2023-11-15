package day14spacestoichiometry

import (
	"fmt"
	"math"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

type Remaining map[string]int
type Requirements map[string]int

type Progress struct {
	CurrentRequirements Requirements
	RemainingMaterials  Remaining
	AmountOre           int
}

type Recepy struct {
	Product string // What product will this recepy give
	Gives   int    // How much of the product will be given?

	Requires Requirements
}

func init() {
	register.Day("14a", solvePart1)
	register.Day("14b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	recepies := ParseInput(lines)

	progress := BuildProgress(recepies)

	for len(progress.CurrentRequirements) > 0 {
		progress.BreakDown(recepies)
	}

	fmt.Printf("Result of day-14 / part-1: %d\n", progress.AmountOre)
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	recepies := ParseInput(lines)

	fmt.Printf("Result of day-14 / part-2: %d\n", FindMaxFuelForAmountOre(recepies, 1000000000000))
}

func FindMaxFuelForAmountOre(recepies []Recepy, amountOre int) int {
	var minBound, maxBound int = 0, amountOre

	for maxBound-minBound > 1 {
		currentBound := (minBound + maxBound) / 2

		progress := BuildProgress(recepies)
		progress.CurrentRequirements["FUEL"] = currentBound

		for len(progress.CurrentRequirements) > 0 {
			progress.BreakDown(recepies)
		}

		if progress.AmountOre > amountOre {
			maxBound = currentBound
		} else {
			minBound = currentBound
		}
	}

	return minBound
}

func (p *Progress) BreakDown(recepies []Recepy) {
	var out = make(Requirements, len(p.CurrentRequirements)*5)

	for product, amount := range p.CurrentRequirements {
		if p.RemainingMaterials[product] > amount {
			p.RemainingMaterials[product] -= amount
			amount = 0
		} else {
			amount -= p.RemainingMaterials[product]
			p.RemainingMaterials[product] = 0
		}

		recepy := findRecepyByProduct(product, recepies)
		repeats := int(math.Ceil(float64(amount) / float64(recepy.Gives)))

		for requiredProduct, requiredAmount := range recepy.Requires {
			if requiredProduct == "ORE" {
				p.AmountOre += requiredAmount * repeats
			} else {
				out[requiredProduct] += requiredAmount * repeats
			}
		}

		created := repeats * recepy.Gives
		p.RemainingMaterials[product] += created - amount
	}

	p.CurrentRequirements = out
	return
}

func findRecepyByProduct(product string, recepies []Recepy) Recepy {
	for _, recepy := range recepies {
		if recepy.Product == product {
			return recepy
		}
	}

	panic("Recepy not found by product")
}

func BuildProgress(recepies []Recepy) *Progress {
	var progress = Progress{
		CurrentRequirements: make(Requirements, len(recepies)),
		RemainingMaterials:  make(Remaining, len(recepies)),
		AmountOre:           0,
	}

	for _, recepy := range recepies {
		progress.RemainingMaterials[recepy.Product] = 0
	}

	progress.CurrentRequirements["FUEL"] = 1
	progress.RemainingMaterials["FUEL"] = 0
	progress.RemainingMaterials["ORE"] = 0

	return &progress
}

func ParseInput(lines []string) (out []Recepy) {
	for _, line := range lines {
		out = append(out, ParseLine(line))
	}

	return
}

func ParseLine(line string) Recepy {
	parts := strings.Split(line, " => ")
	rawRequirements := strings.Split(parts[0], ", ")

	var out = Recepy{
		Requires: make(map[string]int, len(rawRequirements)),
	}

	outParts := strings.Split(parts[1], " ")
	out.Gives = utils.ConvStrToI(outParts[0])
	out.Product = outParts[1]

	for _, rawRequirement := range rawRequirements {
		partRequirement := strings.Split(rawRequirement, " ")
		out.Requires[partRequirement[1]] = utils.ConvStrToI(partRequirement[0])
	}

	return out
}
