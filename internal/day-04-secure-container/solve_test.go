package day04securecontainer_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/internal/day-04-secure-container"
	"github.com/stretchr/testify/assert"
)

func TestIsValidPassword(t *testing.T) {
	assert := assert.New(t)

	assert.True(IsValidPassword(111111))
	assert.True(IsValidPassword(122345))
	assert.False(IsValidPassword(223450))
	assert.False(IsValidPassword(123789))
}

func TestIsExtendedValidPassword(t *testing.T) {
	assert := assert.New(t)
	assert.True(IsExtendedValidPassword(112233))
	assert.False(IsExtendedValidPassword(123444))
	assert.True(IsExtendedValidPassword(111122))
}
