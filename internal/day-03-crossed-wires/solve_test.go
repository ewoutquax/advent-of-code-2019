package day03crossedwires_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/internal/day-03-crossed-wires"
	"github.com/stretchr/testify/assert"
)

func TestGetLocations(t *testing.T) {
	var w = Wire{}
	assert := assert.New(t)

	w.FollowInstruction("U6")
	assert.Equal(0, w.CurrentLocation.X)
	assert.Equal(6, w.CurrentLocation.Y)
	assert.Equal(6, len(w.Locations))

	w.FollowInstruction("R3")
	assert.Equal(3, w.CurrentLocation.X)
	assert.Equal(6, w.CurrentLocation.Y)
	assert.Equal(9, len(w.Locations))

	w.FollowInstruction("D2")
	assert.Equal(3, w.CurrentLocation.X)
	assert.Equal(4, w.CurrentLocation.Y)
	assert.Equal(11, len(w.Locations))

	w.FollowInstruction("L1")
	assert.Equal(2, w.CurrentLocation.X)
	assert.Equal(4, w.CurrentLocation.Y)
	assert.Equal(12, len(w.Locations))

	assert.Equal(0, w.Locations[0].X)
	assert.Equal(1, w.Locations[0].Y)

	assert.Equal(2, w.Locations[11].X)
	assert.Equal(4, w.Locations[11].Y)
}

func TestFindClosestIntersection(t *testing.T) {
	var wire0, wire1 Wire
	assert := assert.New(t)

	wire0 = BuildWireByPath("R8,U5,L5,D3")
	wire1 = BuildWireByPath("U7,R6,D4,L4")
	assert.Equal(30, FindClosestIntersectionBySteps(wire0, wire1))

	wire0 = BuildWireByPath("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	wire1 = BuildWireByPath("U62,R66,U55,R34,D71,R55,D58,R83")
	assert.Equal(610, FindClosestIntersectionBySteps(wire0, wire1))

	wire0 = BuildWireByPath("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	wire1 = BuildWireByPath("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	assert.Equal(410, FindClosestIntersectionBySteps(wire0, wire1))
}
