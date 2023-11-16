package day15oxygensystem_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2019/internal/day-15-oxygen-system"
	"github.com/stretchr/testify/assert"
)

func TestRobot(t *testing.T) {
	robot := BuildRobot()

	assert := assert.New(t)
	assert.Equal("*day15oxygensystem.Robot", fmt.Sprintf("%s", reflect.TypeOf(robot)))
	assert.Equal(RobotRunning, robot.Status)
}

func TestMoveRobotToOpenSpot(t *testing.T) {
	robot := BuildRobot(
		WithIntcoder(MockedIntCoderMoveToOpenSpot{}),
	)

	robot.Move(North)
	assert.Equal(t, &Location{X: 0, Y: 1}, robot.CurrentLocation)
	assert.Equal(t, 1, robot.NrSteps())
	assert.Equal(t, West, robot.NextMovement)
	robot.Move(East)
	assert.Equal(t, &Location{X: 1, Y: 1}, robot.CurrentLocation)
	assert.Equal(t, 2, robot.NrSteps())
	assert.Equal(t, North, robot.NextMovement)
	robot.Move(West)
	assert.Equal(t, &Location{X: 0, Y: 1}, robot.CurrentLocation)
	assert.Equal(t, 1, robot.NrSteps())
	assert.Equal(t, South, robot.NextMovement)
	assert.Equal(t, RobotRunning, robot.Status)
}

func TestMoveRobotIntoWall(t *testing.T) {
	robot := BuildRobot(
		WithIntcoder(MockedIntCoderMoveIntoWall{}),
	)

	robot.Move(North)
	assert.Equal(t, &Location{X: 0, Y: 0}, robot.CurrentLocation)
	assert.Equal(t, 0, robot.NrSteps())
	assert.Equal(t, East, robot.NextMovement)
	robot.Move(East)
	assert.Equal(t, &Location{X: 0, Y: 0}, robot.CurrentLocation)
	assert.Equal(t, 0, robot.NrSteps())
	assert.Equal(t, South, robot.NextMovement)
	assert.Equal(t, RobotRunning, robot.Status)
}

func TestMoveRobotToOxygenTank(t *testing.T) {
	robot := BuildRobot(
		WithIntcoder(MockedIntCoderMoveToOxygenTank{}),
	)

	robot.Move(North)
	assert.Equal(t, &Location{X: 0, Y: 1}, robot.CurrentLocation)
	assert.Equal(t, 1, robot.NrSteps())
	assert.Equal(t, North, robot.NextMovement)
	assert.Equal(t, RobotDone, robot.Status)
}

type MockedIntCoderMoveToOpenSpot struct{}

func (m MockedIntCoderMoveToOpenSpot) Send(_ int)   {}
func (m MockedIntCoderMoveToOpenSpot) Receive() int { return 1 }

type MockedIntCoderMoveIntoWall struct{}

func (m MockedIntCoderMoveIntoWall) Send(_ int)   {}
func (m MockedIntCoderMoveIntoWall) Receive() int { return 0 }

type MockedIntCoderMoveToOxygenTank struct{}

func (m MockedIntCoderMoveToOxygenTank) Send(_ int)   {}
func (m MockedIntCoderMoveToOxygenTank) Receive() int { return 2 }
