package intcoder

func (i *IntCoder) Run() {
	i.status = StatusRunning

	for i.status == StatusRunning {
		statement := i.BuildStatement()
		statement.exec(i)
	}
}

func (i *IntCoder) Status() string {
	return string(i.status)
}

func (i *IntCoder) Set(index, instruction int) {
	i.sourceCode[index] = instruction
}

func (i *IntCoder) Send(input int) {
	i.inputs = append(i.inputs, input)
	i.Run()
}

func (i *IntCoder) Receive() (result int) {
	if len(i.output) == 0 {
		return -1337
	}

	result = i.output[0]
	i.output = i.output[1:]

	return
}

func (i *IntCoder) Result() (sourceCode []int) {
	for idx := 0; idx < len(i.sourceCode); idx++ {
		sourceCode = append(sourceCode, i.sourceCode[idx])
	}

	return
}
