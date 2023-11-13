package intcoder

import "strconv"

func (i *IntCoder) buildParameter(offset int) (p Parameter) {
	opcode := "0000" + strconv.Itoa(i.sourceCode[i.idxInstruction])

	p.Mode = convToMode(opcode[len(opcode)-2-offset])

	switch p.Mode {
	case ModePosition:
		p.Index = i.sourceCode[i.idxInstruction+offset]
	case ModeImmediate:
		p.Index = i.idxInstruction + offset
	case ModeRelative:
		p.Index = i.sourceCode[i.idxInstruction+offset] + i.relativeBase
	}

	p.Value = i.sourceCode[p.Index]

	return
}

func convToMode(b byte) ParameterMode {
	return map[byte]ParameterMode{
		'0': ModePosition,
		'1': ModeImmediate,
		'2': ModeRelative,
	}[b]
}
