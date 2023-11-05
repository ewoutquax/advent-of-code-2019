package intcoder

type IntCoder struct {
	idxInstruction int
	sourceCode     map[int]int
}

func (i *IntCoder) Result() (sourceCode []int) {
	for idx := 0; idx < len(i.sourceCode); idx++ {
		sourceCode = append(sourceCode, i.sourceCode[idx])
	}

	return
}

func (i *IntCoder) Run() {
	var running bool = true

	for running {
		switch i.sourceCode[i.idxInstruction] {
		case 1:
			i.add()
			i.idxInstruction += 4
		case 2:
			i.multiply()
			i.idxInstruction += 4
		case 99:
			running = false
		}
	}
}

func (i *IntCoder) Set(index, instruction int) {
	i.sourceCode[index] = instruction
}

func (i *IntCoder) add() {
	var idxLeft int = i.sourceCode[i.idxInstruction+1]
	var idxRight int = i.sourceCode[i.idxInstruction+2]
	var idxTarget int = i.sourceCode[i.idxInstruction+3]

	i.sourceCode[idxTarget] = i.sourceCode[idxLeft] + i.sourceCode[idxRight]
}

func (i *IntCoder) multiply() {
	var idxLeft int = i.sourceCode[i.idxInstruction+1]
	var idxRight int = i.sourceCode[i.idxInstruction+2]
	var idxTarget int = i.sourceCode[i.idxInstruction+3]

	i.sourceCode[idxTarget] = i.sourceCode[idxLeft] * i.sourceCode[idxRight]
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
