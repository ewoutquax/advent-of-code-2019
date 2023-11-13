package day10monitoringstation_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/internal/day-10-monitoring-station"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	u := ParseInput(testInput())
	assert.Equal(t, "day10monitoringstation.Universe", fmt.Sprintf("%s", reflect.TypeOf(u)))
	assert.Equal(t, 10, len(u.Astroid))
}

func TestVisibleAstroids(t *testing.T) {
	u := ParseInput(testInput())

	assert.Equal(t, 7, VisibleAstroids(&u, Location{X: 1, Y: 0}))
	assert.Equal(t, 7, VisibleAstroids(&u, Location{X: 4, Y: 0}))
	assert.Equal(t, 6, VisibleAstroids(&u, Location{X: 0, Y: 2}))
	assert.Equal(t, 7, VisibleAstroids(&u, Location{X: 1, Y: 2}))
	assert.Equal(t, 7, VisibleAstroids(&u, Location{X: 2, Y: 2}))
	assert.Equal(t, 7, VisibleAstroids(&u, Location{X: 3, Y: 2}))
	assert.Equal(t, 5, VisibleAstroids(&u, Location{X: 4, Y: 2}))
	assert.Equal(t, 7, VisibleAstroids(&u, Location{X: 4, Y: 3}))
	assert.Equal(t, 8, VisibleAstroids(&u, Location{X: 3, Y: 4}))
	assert.Equal(t, 7, VisibleAstroids(&u, Location{X: 4, Y: 4}))
}

func TestMaxCount(t *testing.T) {
	u := ParseInput(testInput())

	assert.Equal(t, 8, u.MaxCount())
}

func testInput() []string {
	return []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	}
}
