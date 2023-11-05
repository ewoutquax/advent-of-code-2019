package intcoder

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
