package intcoder

import "strconv"

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

type ModeIndex uint

const (
	ModePosition ModeIndex = iota + 1
	ModeImmediate
)

type Parameter struct {
	Mode  ModeIndex
	Index int
	Value int
}

type CodeOperation string

const (
	OperationAdd      CodeOperation = "01"
	OperationMultiply               = "02"
	OperationInput                  = "03"
	OperationOutput                 = "04"
	OperationHalt                   = "99"
)

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

type StatementHalt struct {
	size int `default:"1"`
}

func (s StatementAdd) OpCode() CodeOperation      { return OperationAdd }
func (s StatementMultiply) OpCode() CodeOperation { return OperationMultiply }
func (s StatementInput) OpCode() CodeOperation    { return OperationInput }
func (s StatementOutput) OpCode() CodeOperation   { return OperationOutput }
func (s StatementHalt) OpCode() CodeOperation     { return OperationHalt }

func (s StatementAdd) exec(i *IntCoder) {
	i.sourceCode[s.Target.Index] = s.Left.Value + s.Right.Value
	i.idxInstruction += s.size
}
func (s StatementMultiply) exec(i *IntCoder) {
	i.sourceCode[s.Target.Index] = s.Left.Value * s.Right.Value
	i.idxInstruction += s.size
}
func (s StatementInput) exec(i *IntCoder) {
	if len(i.inputs) == 0 {
		i.status = StatusAwaitingInput
	} else {
		i.sourceCode[s.Target.Index] = i.inputs[0]
		i.inputs = i.inputs[1:]
		i.idxInstruction += s.size
	}
}
func (s StatementOutput) exec(i *IntCoder) {
	i.output = i.sourceCode[s.Target.Index]
	i.idxInstruction += s.size
}
func (s StatementHalt) exec(i *IntCoder) {
	i.status = StatusHalted
}

type Statement interface {
	OpCode() CodeOperation
	exec(*IntCoder)
}

func (i *IntCoder) Result() (sourceCode []int) {
	for idx := 0; idx < len(i.sourceCode); idx++ {
		sourceCode = append(sourceCode, i.sourceCode[idx])
	}

	return
}

func (i *IntCoder) Run() {
	i.status = StatusRunning

	for i.status == StatusRunning {
		statement := i.BuildStatement()
		statement.exec(i)
	}
}

func (i *IntCoder) Send(input int) {
	i.inputs = append(i.inputs, input)
	i.Run()
}

func (i *IntCoder) Receive() int {
	return i.output
}

func (i *IntCoder) Set(index, instruction int) {
	i.sourceCode[index] = instruction
}

func Compile(sourceCode []int) *IntCoder {
	intCoder := &IntCoder{
		idxInstruction: 0,
		sourceCode:     make(map[int]int, len(sourceCode)),
	}

	for idx, instruction := range sourceCode {
		intCoder.sourceCode[idx] = instruction
	}

	return intCoder
}

func (i *IntCoder) BuildStatement() Statement {
	buildFunc :=
		map[int]func(*IntCoder) Statement{
			1:  buildStatementAdd,
			2:  buildStatementMultiply,
			3:  buildStatementInput,
			4:  buildStatementOutput,
			99: buildStatementHalt,
		}[i.sourceCode[i.idxInstruction]%100]

	if buildFunc == nil {
		panic("unknown opcode")
	}

	return buildFunc(i)
}

func buildStatementAdd(i *IntCoder) Statement {
	return StatementAdd{
		size:   4,
		Left:   i.buildParameter(1),
		Right:  i.buildParameter(2),
		Target: i.buildParameter(3),
	}
}

func buildStatementMultiply(i *IntCoder) Statement {
	return StatementMultiply{
		size:   4,
		Left:   i.buildParameter(1),
		Right:  i.buildParameter(2),
		Target: i.buildParameter(3),
	}
}

func buildStatementInput(i *IntCoder) Statement {
	return StatementInput{
		size:   2,
		Target: i.buildParameter(1),
	}
}

func buildStatementOutput(i *IntCoder) Statement {
	return StatementOutput{
		size:   2,
		Target: i.buildParameter(1),
	}
}

func buildStatementHalt(i *IntCoder) Statement {
	return StatementHalt{
		size: 1,
	}
}

func (i *IntCoder) buildParameter(offset int) (p Parameter) {
	opcode := "0000" + strconv.Itoa(i.sourceCode[i.idxInstruction])

	p.Mode = convToMode(opcode[len(opcode)-2-offset])

	if p.Mode == ModePosition {
		p.Index = i.sourceCode[i.idxInstruction+offset]
	}
	if p.Mode == ModeImmediate {
		p.Index = i.idxInstruction + offset
	}

	p.Value = i.sourceCode[p.Index]

	return
}

func convToMode(b byte) ModeIndex {
	return map[byte]ModeIndex{
		'0': ModePosition,
		'1': ModeImmediate,
	}[b]
}
