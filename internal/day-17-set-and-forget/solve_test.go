package day17setandforget_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/internal/day-17-set-and-forget"
	"github.com/stretchr/testify/assert"
)

func TestFindIntersections(t *testing.T) {
	assert := assert.New(t)

	intersections := FindIntersections(testFixture())

	assert.Len(intersections, 4)

	expectedIntersections := []Intersection{
		{Location{X: 2, Y: 2}},
		{Location{X: 2, Y: 4}},
		{Location{X: 6, Y: 4}},
		{Location{X: 10, Y: 4}},
	}

	for _, expectedIntersection := range expectedIntersections {
		assert.Contains(intersections, expectedIntersection)
	}
}

func TestSumAlignmentParameters(t *testing.T) {
	sum := SumAlignmentParameters(testFixture())
	assert.Equal(t, 76, sum)
}

func TestOptimizePath(t *testing.T) {
	path := "R8,R8,R4,R4,R8,L6,L2,R4,R4,R8,R8,R8,L6,L2"

	expectedParts := [][]string{
		{"R8,R8"},
		{"R4,R4,R8"},
		{"L6,L2"},
	}

	optimizedParts := GenerateOptimizedParts(path)

	for _, part := range optimizedParts {
		assert.Contains(t, expectedParts, part)
	}
}

func testFixture() []string {
	return []string{
		"..#..........",
		"..#..........",
		"#######...###",
		"#.#...#...#.#",
		"#############",
		"..#...#...#..",
		"..#####...^..",
	}
}
