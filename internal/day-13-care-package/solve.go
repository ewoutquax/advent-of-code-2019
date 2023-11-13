package day13carepackage

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/intcoder"
	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

type Arcade struct {
	intCoder *intcoder.IntCoder

	BallX int // x-position of the ball
	PadX  int // x-position of the pad; use this to move the joystick

	Score    int // Current Score
	NrBlocks int // nr of blocks in the game; when no blocks are left, we're done
}

func (a *Arcade) InsertQuarter() {
	a.intCoder.Set(0, 2)
	a.intCoder.Run()
}

func (a *Arcade) PlayRound() {
	doContinue := true
	a.NrBlocks = 0

	for doContinue {
		x := a.intCoder.Receive()
		y := a.intCoder.Receive()
		v := a.intCoder.Receive()

		switch v {
		case 4:
			a.BallX = x
		case 3:
			a.PadX = x
		case 2:
			a.NrBlocks++
		default:
			if x == -1 && y == 0 {
				a.Score = v
			}
		}

		if v == -1337 {
			doContinue = false
		}
	}
}

func (a *Arcade) MoveJoystick() {
	if a.BallX < a.PadX {
		a.intCoder.Send(-1)
	} else if a.BallX > a.PadX {
		a.intCoder.Send(1)
	} else {
		a.intCoder.Send(0)
	}
}

func init() {
	register.Day("13a", solvePart1)
	register.Day("13b", solvePart2)
}

func solvePart1(inputFile string) {
	intCoder := ParseInput(inputFile)

	fmt.Printf("Result of day-13 / part-1: %d\n", CountBlocks(intCoder))
}

func solvePart2(inputFile string) {
	intCoder := ParseInput(inputFile)
	arcade := BuildArcadeWithIntCoder(intCoder)
	arcade.InsertQuarter()

	for arcade.intCoder.Status() != "halted" {
		arcade.MoveJoystick()
		arcade.PlayRound()
	}

	fmt.Printf("Result of day-13 / part-2: %d\n", arcade.Score)
}

func CountBlocks(intCoder *intcoder.IntCoder) (nrBlocks int) {
	intCoder.Run()

	canContinue := true
	for canContinue {
		intCoder.Receive() // read 'X'-coordinate
		intCoder.Receive() // read 'Y'-coordinate
		typeTile := intCoder.Receive()
		if typeTile == 2 {
			nrBlocks++
		}

		if typeTile == -1337 {
			canContinue = false
		}
	}

	return
}

func ParseInput(inputFile string) *intcoder.IntCoder {
	var sourceCode []int

	line := utils.ReadFileAsLine(inputFile)
	for _, char := range strings.Split(line, ",") {
		sourceCode = append(sourceCode, utils.ConvStrToI(char))
	}

	return intcoder.Compile(sourceCode)
}

func BuildArcadeWithIntCoder(intCoder *intcoder.IntCoder) *Arcade {
	return &Arcade{
		intCoder: intCoder,
	}
}
