package intcoder

type Status string

const (
	StatusRunning       Status = "running"
	StatusAwaitingInput        = "awaiting input"
	StatusHalted               = "halted"
)

type IntCoder struct {
	status Status

	idxInstruction int
	relativeBase   int
	sourceCode     map[int]int

	inputs []int
	output []int
}

type ParameterMode uint

const (
	ModePosition ParameterMode = iota + 1
	ModeImmediate
	ModeRelative
)

type Parameter struct {
	Mode  ParameterMode
	Index int
	Value int
}

type CodeOperation string

const (
	OperationAdd                CodeOperation = "01"
	OperationMultiply                         = "02"
	OperationInput                            = "03"
	OperationOutput                           = "04"
	OperationJumpIfTrue                       = "05"
	OperationJumpIfFalse                      = "06"
	OperationLessThen                         = "07"
	OperationEqual                            = "08"
	OperationAdjustRelativeBase               = "08"
	OperationHalt                             = "99"
)

type Statement interface {
	exec(*IntCoder)
}

type StatementAdd struct {
	size int `default:"4"`

	Left   Parameter
	Right  Parameter
	Target Parameter
}

type StatementMultiply struct {
	size int `default:"4"`

	Left   Parameter
	Right  Parameter
	Target Parameter
}

type StatementInput struct {
	size int `default:"2"`

	Target Parameter
}

type StatementOutput struct {
	size int `default:"2"`

	Target Parameter
}

type StatementJumpIfTrue struct {
	size int `default:"3"`

	Comparer    Parameter
	JumpToIndex Parameter
}

type StatementJumpIfFalse struct {
	size int `default:"3"`

	Comparer    Parameter
	JumpToIndex Parameter
}

type StatementLessThen struct {
	size int `default:"2"`

	Left   Parameter
	Right  Parameter
	Target Parameter
}

type StatementEqual struct {
	size int `default:"2"`

	Left   Parameter
	Right  Parameter
	Target Parameter
}

type StatementAdjustRelativeBase struct {
	size int `default:"1"`

	Target Parameter
}

type StatementHalt struct {
	size int `default:"1"`
}
