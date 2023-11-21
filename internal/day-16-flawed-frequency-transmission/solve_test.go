package day16flawedfrequencytransmission_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/internal/day-16-flawed-frequency-transmission"
	"github.com/stretchr/testify/assert"
)

func TestMultiplier(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, GetMultiplier(1, 0))
	assert.Equal(0, GetMultiplier(1, 1))
	assert.Equal(-1, GetMultiplier(1, 2))
	assert.Equal(0, GetMultiplier(1, 3))
	assert.Equal(1, GetMultiplier(1, 4))
	assert.Equal(0, GetMultiplier(1, 5))
	assert.Equal(-1, GetMultiplier(1, 6))
	assert.Equal(0, GetMultiplier(1, 7))

	assert.Equal(0, GetMultiplier(2, 0))
	assert.Equal(1, GetMultiplier(2, 1))
	assert.Equal(1, GetMultiplier(2, 2))
	assert.Equal(0, GetMultiplier(2, 3))
	assert.Equal(0, GetMultiplier(2, 4))
	assert.Equal(-1, GetMultiplier(2, 5))
	assert.Equal(-1, GetMultiplier(2, 6))
	assert.Equal(0, GetMultiplier(2, 7))
}

func TestOutputAfterXPhasesByOffset(t *testing.T) {
	// assert.Equal(t, 84462026, OutputAfterXPhasesByOffset("4603456581043215999473424395046570134221182852363891114374810263887875638355730605895695123598637121"))
	assert.Equal(t, "84462026", OutputAfterXPhasesByOffset("03036732577212944063491565474664"))
}

func TestApplyPhase(t *testing.T) {
	assert.Equal(t, "4", ApplyPhase("12345678", 1))
	assert.Equal(t, "8", ApplyPhase("12345678", 2))
	assert.Equal(t, "2", ApplyPhase("12345678", 3))
	assert.Equal(t, "2", ApplyPhase("12345678", 4))
	assert.Equal(t, "6", ApplyPhase("12345678", 5))
	assert.Equal(t, "1", ApplyPhase("12345678", 6))
	assert.Equal(t, "5", ApplyPhase("12345678", 7))
	assert.Equal(t, "8", ApplyPhase("12345678", 8))
}

func TestApplyPhases(t *testing.T) {
	assert.Equal(t, "34040438", ApplyPhases("48226158"))
}

func TestOutputAfterXPhases(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("24176176", OutputAfterXPhases("80871224585914546619083218645595", 100))
	assert.Equal("73745418", OutputAfterXPhases("19617804207202209144916044189917", 100))
	assert.Equal("52432133", OutputAfterXPhases("69317163492948606335995924319873", 100))
}
