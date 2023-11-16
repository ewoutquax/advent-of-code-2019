package day15oxygensystem

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2019/pkg/intcoder"
	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
	"github.com/ewoutquax/advent-of-code-2019/pkg/utils"
)

var knownLocations = make(map[Location]int, 0)

type BaseIntCoder interface {
	Send(int)
	Receive() int
}

type Movement uint

const (
	North Movement = 1
	South Movement = 2
	West  Movement = 3
	East  Movement = 4
)

type Status uint

const (
	StatusHitWall       Status = 0
	StatusMoveSucceed   Status = 1
	StatusOxygenReached Status = 2

	RobotRunning Status = iota + 1
	RobotDone
)

type Location struct {
	X int
	Y int
}

func (l *Location) update(m Movement) {
	currentSteps, currentLocExists := knownLocations[*l]
	if !currentLocExists {
		currentSteps = 0
	}

	vXY := map[Movement][2]int{
		North: {0, 1},
		East:  {1, 0},
		South: {0, -1},
		West:  {-1, 0},
	}[m]

	l.X += vXY[0]
	l.Y += vXY[1]

	if _, newLocExists := knownLocations[*l]; !newLocExists {
		knownLocations[*l] = currentSteps + 1
	}
}

type Robot struct {
	intCoder BaseIntCoder
	Status   Status // Status of the robot; used to check if the robot has reached the tank

	CurrentLocation *Location // Location of the robot
	NextMovement    Movement  // Direction the robot has move to move NEXT step; depends on hitting a wall
}

func (r *Robot) NrSteps() int {
	return knownLocations[*r.CurrentLocation]
}

func init() {
	register.Day("15a", solvePart1)
	register.Day("15b", solvePart2)
}

func solvePart1(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)

	var sourceCode []int
	for _, statement := range strings.Split(line, ",") {
		code := utils.ConvStrToI(statement)
		sourceCode = append(sourceCode, code)
	}

	intCoder := intcoder.Compile(sourceCode)

	robot := BuildRobot(
		WithIntcoder(intCoder),
	)

	for robot.Status != RobotDone {
		robot.Move(robot.NextMovement)
	}

	fmt.Printf("Result of day-15 / part-1: %d\n", robot.NrSteps())
}

func solvePart2(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)

	var sourceCode []int
	for _, statement := range strings.Split(line, ",") {
		code := utils.ConvStrToI(statement)
		sourceCode = append(sourceCode, code)
	}

	intCoder := intcoder.Compile(sourceCode)

	robot := BuildRobot(
		WithIntcoder(intCoder),
	)

	for robot.Status != RobotDone {
		robot.Move(robot.NextMovement)
	}

	// from the tank, walk around the whole ship.
	// The max-number of steps from the tank is the solution
	knownLocations = make(map[Location]int, 0)
	robot.Status = RobotRunning
	for robot.Status != RobotDone {
		robot.Move(robot.NextMovement)
	}

	var maxSteps int = 0
	for _, steps := range knownLocations {
		if maxSteps < steps {
			maxSteps = steps
		}
	}

	fmt.Printf("Result of day-15 / part-2: %d\n", maxSteps)
}

func TurnClockwise(m Movement) (next Movement) {
	switch m {
	case North:
		next = East
	case East:
		next = South
	case South:
		next = West
	case West:
		next = North
	}

	return
}

func TurnCounterClockwise(m Movement) (next Movement) {
	next = m
	for ctr := 0; ctr < 3; ctr++ {
		next = TurnClockwise(next)
	}

	return
}

type RobotOptsFunc func(*Robot)

func (r *Robot) Move(m Movement) {
	r.intCoder.Send(int(m))

	switch Status(r.intCoder.Receive()) {
	case StatusMoveSucceed:
		r.CurrentLocation.update(m)
		r.NextMovement = TurnCounterClockwise(m)
	case StatusHitWall:
		r.NextMovement = TurnClockwise(m)
	case StatusOxygenReached:
		r.CurrentLocation.update(m)
		r.Status = RobotDone
	}
}

func BuildRobot(optsFunc ...RobotOptsFunc) *Robot {
	robot := defaultRobot()

	for _, optsFunc := range optsFunc {
		optsFunc(robot)
	}

	return robot
}

func defaultRobot() *Robot {
	return &Robot{
		Status: RobotRunning,
		CurrentLocation: &Location{
			X: 0,
			Y: 0,
		},
		NextMovement: North,
	}
}

func WithIntcoder(i BaseIntCoder) RobotOptsFunc {
	return func(r *Robot) {
		r.intCoder = i
	}
}
