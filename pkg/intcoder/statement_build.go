package intcoder

import (
	"fmt"
)

func (i *IntCoder) BuildStatement() Statement {
	buildFunc :=
		map[int]func(*IntCoder) Statement{
			1:  buildStatementAdd,
			2:  buildStatementMultiply,
			3:  buildStatementInput,
			4:  buildStatementOutput,
			5:  buildStatementJumpIfTrue,
			6:  buildStatementJumpIfFalse,
			7:  buildStatementLessThen,
			8:  buildStatementEqual,
			99: buildStatementHalt,
		}[i.sourceCode[i.idxInstruction]%100]

	if buildFunc == nil {
		fmt.Printf("BuildStatement: idxInstruction / OpCode: %d / %d\n", i.idxInstruction, i.sourceCode[i.idxInstruction])
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

func buildStatementJumpIfTrue(i *IntCoder) Statement {
	return StatementJumpIfTrue{
		size:        3,
		Comparer:    i.buildParameter(1),
		JumpToIndex: i.buildParameter(2),
	}
}

func buildStatementJumpIfFalse(i *IntCoder) Statement {
	return StatementJumpIfFalse{
		size:        3,
		Comparer:    i.buildParameter(1),
		JumpToIndex: i.buildParameter(2),
	}
}

func buildStatementLessThen(i *IntCoder) Statement {
	return StatementLessThen{
		size:   4,
		Left:   i.buildParameter(1),
		Right:  i.buildParameter(2),
		Target: i.buildParameter(3),
	}
}

func buildStatementEqual(i *IntCoder) Statement {
	return StatementEqual{
		size:   4,
		Left:   i.buildParameter(1),
		Right:  i.buildParameter(2),
		Target: i.buildParameter(3),
	}
}

func buildStatementHalt(i *IntCoder) Statement {
	return StatementHalt{
		size: 1,
	}
}
