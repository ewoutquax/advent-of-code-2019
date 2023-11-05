package day01thetyrannyoftherocketequation_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/internal/day-01-the-tyranny-of-the-rocket-equation"
	"github.com/stretchr/testify/assert"
)

func TestCalculateFuel(t *testing.T) {
	assert.Equal(t, 2, CalculateFuel(12))
	assert.Equal(t, 2, CalculateFuel(14))
	assert.Equal(t, 654, CalculateFuel(1969))
	assert.Equal(t, 33583, CalculateFuel(100756))
}

func TestCalculateCumulativeFuel(t *testing.T) {
	assert.Equal(t, 2, CalculateCumulativeFuel(12))
	assert.Equal(t, 966, CalculateCumulativeFuel(1969))
	assert.Equal(t, 50346, CalculateCumulativeFuel(100756))
}
