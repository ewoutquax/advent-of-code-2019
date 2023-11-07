package day08spaceimageformat_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/internal/day-08-space-image-format"
	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	input := "001002121212"

	assert.Equal(t, 9, Checksum(input, 3, 2))
}

func TestPrintInput(t *testing.T) {
	input := "0222112222120000"

	output := Print(input, 2, 2)

	assert.Equal(t, "01", output[0])
	assert.Equal(t, "10", output[1])
}
