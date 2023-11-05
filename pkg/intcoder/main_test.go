package intcoder_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/pkg/intcoder"
	"github.com/stretchr/testify/assert"
)

func TestSimpleProgram(t *testing.T) {
	var sourceCode, expectedResult []int
	var intCode *IntCoder

	sourceCode = []int{1, 0, 0, 0, 99}
	expectedResult = []int{2, 0, 0, 0, 99}
	intCode = Compile(sourceCode)
	intCode.Run()

	assert.Equal(t, expectedResult, intCode.Result())

	sourceCode = []int{2, 3, 0, 3, 99}
	expectedResult = []int{2, 3, 0, 6, 99}
	intCode = Compile(sourceCode)
	intCode.Run()

	assert.Equal(t, expectedResult, intCode.Result())

	sourceCode = []int{2, 4, 4, 5, 99, 0}
	expectedResult = []int{2, 4, 4, 5, 99, 9801}
	intCode = Compile(sourceCode)
	intCode.Run()

	assert.Equal(t, expectedResult, intCode.Result())

	sourceCode = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expectedResult = []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	intCode = Compile(sourceCode)
	intCode.Run()

	assert.Equal(t, expectedResult, intCode.Result())
}
