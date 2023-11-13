package intcoder

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
	i.output = append(i.output, i.sourceCode[s.Target.Index])
	i.idxInstruction += s.size
}
func (s StatementJumpIfTrue) exec(i *IntCoder) {
	if s.Comparer.Value == 0 {
		i.idxInstruction += s.size
	} else {
		i.idxInstruction = s.JumpToIndex.Value
	}
}
func (s StatementJumpIfFalse) exec(i *IntCoder) {
	if s.Comparer.Value == 0 {
		i.idxInstruction = s.JumpToIndex.Value
	} else {
		i.idxInstruction += s.size
	}
}
func (s StatementLessThen) exec(i *IntCoder) {
	if s.Left.Value < s.Right.Value {
		i.sourceCode[s.Target.Index] = 1
	} else {
		i.sourceCode[s.Target.Index] = 0
	}
	i.idxInstruction += s.size
}
func (s StatementEqual) exec(i *IntCoder) {
	if s.Left.Value == s.Right.Value {
		i.sourceCode[s.Target.Index] = 1
	} else {
		i.sourceCode[s.Target.Index] = 0
	}
	i.idxInstruction += s.size
}
func (s StatementAdjustRelativeBase) exec(i *IntCoder) {
	i.relativeBase += s.Target.Value
	i.idxInstruction += s.size
}
func (s StatementHalt) exec(i *IntCoder) {
	i.status = StatusHalted
}
