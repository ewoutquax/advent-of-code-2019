package intcoder

import "strconv"

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

func convToMode(b byte) ParameterMode {
	return map[byte]ParameterMode{
		'0': ModePosition,
		'1': ModeImmediate,
	}[b]
}
