package intcoder

type Status uint

const (
	StatusRunning Status = iota + 1
	StatusAwaitingInput
	StatusHalted
)

type IntCoder struct {
	status Status

	idxInstruction int
	sourceCode     map[int]int

	inputs []int
	output int
}

type ParameterMode uint

const (
	ModePosition ParameterMode = iota + 1
	ModeImmediate
)

type Parameter struct {
	Mode  ParameterMode
	Index int
	Value int
}

type CodeOperation string

const (
	OperationAdd         CodeOperation = "01"
	OperationMultiply                  = "02"
	OperationInput                     = "03"
	OperationOutput                    = "04"
	OperationJumpIfTrue                = "05"
	OperationJumpIfFalse               = "06"
	OperationLessThen                  = "07"
	OperationEqual                     = "08"
	OperationHalt                      = "99"
)

type Statement interface {
	exec(*IntCoder)
}

type StatementAdd struct {
	size   int `default:"4"`
	OpCode CodeOperation

	Left   Parameter
	Right  Parameter
	Target Parameter
}

type StatementMultiply struct {
	size   int `default:"4"`
	OpCode CodeOperation

	Left   Parameter
	Right  Parameter
	Target Parameter
}

type StatementInput struct {
	size   int `default:"2"`
	OpCode CodeOperation

	Target Parameter
}

type StatementOutput struct {
	size   int `default:"2"`
	OpCode CodeOperation

	Target Parameter
}

type StatementJumpIfTrue struct {
	size   int `default:"3"`
	OpCode CodeOperation

	Comparer    Parameter
	JumpToIndex Parameter
}

type StatementJumpIfFalse struct {
	size   int `default:"3"`
	OpCode CodeOperation

	Comparer    Parameter
	JumpToIndex Parameter
}

type StatementLessThen struct {
	size   int `default:"2"`
	OpCode CodeOperation

	Left   Parameter
	Right  Parameter
	Target Parameter
}

type StatementEqual struct {
	size   int `default:"2"`
	OpCode CodeOperation

	Left   Parameter
	Right  Parameter
	Target Parameter
}

type StatementHalt struct {
	size   int `default:"1"`
	OpCode CodeOperation
}
