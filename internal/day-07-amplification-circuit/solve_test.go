package day07amplificationcircuit_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/internal/day-07-amplification-circuit"
	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	result := Permutations([]int{1, 2})

	expected := [][]int{
		{1, 2},
		{2, 1},
	}
	assert.Equal(t, expected, result)

	result = Permutations([]int{1, 2, 3})

	expected = [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	assert.Equal(t, expected, result)
}

func TestThrusterSignal(t *testing.T) {
	var sequence = []int{4, 3, 2, 1, 0}
	var sourceSource = []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}

	assert.Equal(t, 43210, ThrusterSignalBySequence(sequence, sourceSource))
}

func TestMaxThrusterSignal(t *testing.T) {
	var sourceSource = []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}

	assert.Equal(t, 43210, MaxThrusterSignal(sourceSource))
}

func TestLoopedThrusterSignal(t *testing.T) {
	var sequence = []int{9, 8, 7, 6, 5}
	var sourceSource = []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}

	assert.Equal(t, 139629729, LoopedThrusterSignalBySequence(sequence, sourceSource))
}

func TestMaxLoopedThrusterSignal(t *testing.T) {
	var sourceSource = []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}

	assert.Equal(t, 139629729, MaxLoopedThrusterSignalBySequence(sourceSource))
}
