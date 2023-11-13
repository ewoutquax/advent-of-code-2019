package day13carepackage_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/internal/day-13-care-package"
	"github.com/stretchr/testify/assert"
)

func TestCountBlocks(t *testing.T) {
	intCoder := ParseInput("./input.txt")

	assert.Equal(t, 309, CountBlocks(intCoder))
}

func TestPlayRound(t *testing.T) {
	assert := assert.New(t)
	intCoder := ParseInput("./input.txt")

	arcade := BuildArcadeWithIntCoder(intCoder)
	arcade.InsertQuarter()

	arcade.PlayRound()
	assert.Equal(15, arcade.BallX)
	assert.Equal(17, arcade.PadX)
	assert.Equal(0, arcade.Score)
	assert.Equal(309, arcade.NrBlocks)

	arcade.MoveJoystick()
	arcade.PlayRound()
	assert.Equal(16, arcade.BallX)
	assert.Equal(16, arcade.PadX)
	assert.Equal(0, arcade.Score)
}
