package intcoder_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/pkg/intcoder"
	"github.com/stretchr/testify/assert"
)

func TestBuildStatement(t *testing.T) {
	assert := assert.New(t)
	intCoder := Compile([]int{1, 0, 0, 0, 99})

	statement := intCoder.BuildStatement()
	assert.Equal("intcoder.StatementAdd", fmt.Sprintf("%s", reflect.TypeOf(statement)))

	statementAdd := statement.(StatementAdd)
	assert.Equal(0, statementAdd.Left.Index)
	assert.Equal(1, statementAdd.Left.Value)
	assert.Equal(ModePosition, statementAdd.Target.Mode)
}

func TestSimpleProgram(t *testing.T) {
	var sourceCode, expectedResult []int
	var intCoder *IntCoder

	sourceCode = []int{1, 0, 0, 0, 99}
	expectedResult = []int{2, 0, 0, 0, 99}
	intCoder = Compile(sourceCode)
	intCoder.Run()

	assert.Equal(t, expectedResult, intCoder.Result())

	sourceCode = []int{2, 3, 0, 3, 99}
	expectedResult = []int{2, 3, 0, 6, 99}
	intCoder = Compile(sourceCode)
	intCoder.Run()

	assert.Equal(t, expectedResult, intCoder.Result())

	sourceCode = []int{2, 4, 4, 5, 99, 0}
	expectedResult = []int{2, 4, 4, 5, 99, 9801}
	intCoder = Compile(sourceCode)
	intCoder.Run()

	assert.Equal(t, expectedResult, intCoder.Result())

	sourceCode = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expectedResult = []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	intCoder = Compile(sourceCode)
	intCoder.Run()

	assert.Equal(t, expectedResult, intCoder.Result())
}

func TestOperationInputOutput(t *testing.T) {
	sourceCode := []int{3, 0, 4, 0, 99}
	intCoder := Compile(sourceCode)

	intCoder.Run()
	intCoder.Send(1337)
	assert.Equal(t, 1337, intCoder.Receive())
}

func TestProgramWithImmediateParams(t *testing.T) {
	var sourceCode, expectedResult []int
	var intCoder *IntCoder

	sourceCode = []int{1002, 4, 3, 4, 33}
	expectedResult = []int{1002, 4, 3, 4, 99}
	intCoder = Compile(sourceCode)
	intCoder.Run()
	assert.Equal(t, expectedResult, intCoder.Result())

	sourceCode = []int{1101, 100, -1, 4, 0}
	expectedResult = []int{1101, 100, -1, 4, 99}
	intCoder = Compile(sourceCode)
	intCoder.Run()
	assert.Equal(t, expectedResult, intCoder.Result())
}

func TestEqual(t *testing.T) {
	var intCoder *IntCoder
	sourceCode := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}

	intCoder = Compile(sourceCode)
	intCoder.Send(8)
	assert.Equal(t, 1, intCoder.Receive())

	intCoder = Compile(sourceCode)
	intCoder.Send(5)
	assert.Equal(t, 0, intCoder.Receive())

	intCoder = Compile(sourceCode)
	intCoder.Send(10)
	assert.Equal(t, 0, intCoder.Receive())
}

func TestLessThen(t *testing.T) {
	var intCoder *IntCoder
	sourceCode := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}

	intCoder = Compile(sourceCode)
	intCoder.Send(5)
	assert.Equal(t, 1, intCoder.Receive())

	intCoder = Compile(sourceCode)
	intCoder.Send(8)
	assert.Equal(t, 0, intCoder.Receive())

	intCoder = Compile(sourceCode)
	intCoder.Send(10)
	assert.Equal(t, 0, intCoder.Receive())
}

func TestJumps(t *testing.T) {
	var intCoder *IntCoder
	sourceCode := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}

	intCoder = Compile(sourceCode)
	intCoder.Send(5)
	assert.Equal(t, 999, intCoder.Receive())

	intCoder = Compile(sourceCode)
	intCoder.Send(8)
	assert.Equal(t, 1000, intCoder.Receive())

	intCoder = Compile(sourceCode)
	intCoder.Send(10)
	assert.Equal(t, 1001, intCoder.Receive())
}

func TestRelativeMode(t *testing.T) {
	sourceCode := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}

	intCoder := Compile(sourceCode)
	intCoder.Run()

	var result []int
	for next := intCoder.Receive(); next != -1337; next = intCoder.Receive() {
		result = append(result, next)
	}

	assert.Equal(t, result, sourceCode)
}
