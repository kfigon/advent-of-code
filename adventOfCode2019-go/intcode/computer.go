package intcode

import "math"

type Computer struct {
	instructions []int
	userInput    *inputHandler
	userOutput   int
	instructionCounter int
	relativeBase int
	extendedMemory map[int]int
}

func NewComputer(instructions []int) *Computer {
	return &Computer{
		instructions: instructions,
		userInput: newInputHandler(),
		extendedMemory: map[int]int{},
	}
}

const (
	OP_ADD       = 1
	OP_MULT      = 2
	OP_INPUT     = 3
	OP_OUTPUT    = 4
	OP_JMP_T = 5
	OP_JMP_F = 6
	OP_LT = 7
	OP_EQ = 8
	OP_SET_RELATIVE = 9
	OP_TERMINATE = 99

	IDX_TERMINATE = -1

	MODE_POSITION  = 0
	MODE_IMMEDIATE = 1
	MODE_RELATIVE = 2
)

func (c *Computer) Calc() []int {
	for c.moreComputations() {
		c.instructionCounter = c.handleCommand(c.instructionCounter)
	}
	return c.instructions
}

func (c *Computer) SingleInstruction() bool {
	c.instructionCounter = c.handleCommand(c.instructionCounter)
	return c.instructionCounter == IDX_TERMINATE
}

func (c *Computer) NextInput() bool {
	opcode := opcode(c.getValue(c.instructionCounter))
	return opcode.extractOpcode() == OP_INPUT
}

func (c *Computer) NextOuput() bool {
	opcode := opcode(c.getValue(c.instructionCounter))
	return opcode.extractOpcode() == OP_OUTPUT
}

func (c *Computer) getValue(idx int) int {
	if idx < len(c.instructions) {
		return c.instructions[idx]
	}
	return c.extendedMemory[idx]
}

func (c *Computer) setValue(idx, val int) {
	if idx < len(c.instructions) {
		c.instructions[idx] = val
	} else {
		c.extendedMemory[idx] = val
	}
}

func (c *Computer) CalcTilOutput() bool {

	for c.moreComputations() {
		c.instructionCounter = c.handleCommand(c.instructionCounter)
		
		if c.moreComputations() && opcode(c.getValue(c.instructionCounter)).extractOpcode() == OP_INPUT {
			break
		}
	}
	return c.instructionCounter == IDX_TERMINATE
}

func (c *Computer) moreComputations() bool {
	return c.instructionCounter < len(c.instructions) && c.instructionCounter != IDX_TERMINATE
}

func (c *Computer) handleCommand(idx int) int {
	opcode := opcode(c.getValue(idx))
	switch opcode.extractOpcode() {
	case OP_ADD:
		return c.handleAdd(idx)
	case OP_MULT:
		return c.handleMult(idx)
	case OP_INPUT:
		return c.handleInput(idx)
	case OP_OUTPUT:
		return c.handleOutput(idx)
	case OP_JMP_T:
		return c.handleJumpTrue(idx)
	case OP_JMP_F:
		return c.handleJumpFalse(idx)
	case OP_LT:
		return c.handleLessThan(idx)
	case OP_EQ:
		return c.handleEquals(idx)
	case OP_SET_RELATIVE:
		return c.handleSetRelative(idx)
	case OP_TERMINATE:
		return IDX_TERMINATE
	}
	// should not happen
	return IDX_TERMINATE
}

func (c *Computer) handleAdd(idx int) int {
	op := opcode(c.getValue(idx))
	param0, param1,param2 := c.getValue(idx+1), c.getValue(idx+2), c.getValue(idx+3)
	c.setValue(c.paramValueForOutput(op.mode(2),param2), c.paramValue(op.mode(0),param0) + c.paramValue(op.mode(1),param1))
	return idx + 4
}

func (c *Computer) handleInput(idx int) int {
	op := opcode(c.getValue(idx))
	param0 := c.getValue(idx+1)
	paramMode := op.mode(0)
	c.setValue(c.paramValueForOutput(paramMode,param0), c.userInput.next())
	return idx + 2
}

func (c *Computer) handleOutput(idx int) int {
	op := opcode(c.getValue(idx))
	param0 := c.getValue(idx+1)
	c.userOutput = c.paramValue(op.mode(0),param0)
	return idx + 2
}

func (c *Computer) handleMult(idx int) int {
	op := opcode(c.getValue(idx))
	param0, param1,param2 := c.getValue(idx+1), c.getValue(idx+2), c.getValue(idx+3)
	c.setValue(c.paramValueForOutput(op.mode(2),param2), c.paramValue(op.mode(0),param0) * c.paramValue(op.mode(1),param1))
	return idx + 4
}

func (c *Computer) handleJumpTrue(idx int) int {
	op := opcode(c.getValue(idx))
	param0, param1 := c.getValue(idx+1), c.getValue(idx+2)
	if c.paramValue(op.mode(0), param0) != 0 {
		return c.paramValue(op.mode(1), param1)
	}
	return idx + 3
}

func (c *Computer) handleJumpFalse(idx int) int {
	op := opcode(c.getValue(idx))
	param0, param1 := c.getValue(idx+1), c.getValue(idx+2)
	if c.paramValue(op.mode(0), param0) == 0 {
		return c.paramValue(op.mode(1), param1)
	}
	return idx + 3
}

func (c *Computer) handleLessThan(idx int) int {
	op := opcode(c.getValue(idx))
	param0, param1,param2 := c.getValue(idx+1), c.getValue(idx+2),c.getValue(idx+3)
	v0 := c.paramValue(op.mode(0), param0)
	v1 := c.paramValue(op.mode(1), param1)
	if v0 < v1 {
		c.setValue(c.paramValueForOutput(op.mode(2),param2), 1)
	} else {
		c.setValue(c.paramValueForOutput(op.mode(2),param2), 0)
	}
	return idx + 4
}

func (c *Computer) handleEquals(idx int) int {
	op := opcode(c.getValue(idx))
	param0, param1,param2 := c.getValue(idx+1), c.getValue(idx+2),c.getValue(idx+3)
	v0 := c.paramValue(op.mode(0), param0)
	v1 := c.paramValue(op.mode(1), param1)
	if v0 == v1 {
		c.setValue(c.paramValueForOutput(op.mode(2),param2), 1)
	} else {
		c.setValue(c.paramValueForOutput(op.mode(2),param2), 0)
	}
	return idx + 4
}

func (c *Computer) handleSetRelative(idx int) int {
	op := opcode(c.getValue(idx))
	param0 := c.getValue(idx+1)
	c.relativeBase += c.paramValue(op.mode(0), param0)
	return idx + 2
}

func (c *Computer) paramValueForOutput(mode, val int) int {
	if mode == MODE_RELATIVE {
		return val + c.relativeBase
	}
	return val
}

func (c *Computer) paramValue(paramMode int, value int) int {
	if paramMode == MODE_POSITION {
		return c.getValue(value)
	} else if paramMode == MODE_IMMEDIATE {
		return value
	} else if paramMode == MODE_RELATIVE {
		return c.getValue(value + c.relativeBase)
	}
	// should never happen
	return -1
}

func (c *Computer) SetUserInput(val int)  {
	c.userInput.add(val)
}

func (c *Computer) ClearUserInput() {
	c.userInput = newInputHandler()
}

func (c *Computer) GetOutput() int {
	return c.userOutput
}

type opcode int
func (op opcode) extractOpcode() int {
	return int(op) % 100
}

func (op opcode) mode(paramIdx int) int {
	paramFlags := int(op) / 100
	return (paramFlags / int(math.Pow10(paramIdx))) % 10
}

type inputHandler struct {
	inputs []int
	counter int
}

func newInputHandler() *inputHandler {
	return &inputHandler{
		inputs: make([]int, 0),
		counter: 0,
	}
}

func (in *inputHandler) next() int {
	if len(in.inputs) == 0 {
		return 0
	}
	out := in.inputs[in.counter]
	in.counter = (in.counter+1) % len(in.inputs)
	return out
}

func (in *inputHandler) add(val int) {
	in.inputs = append(in.inputs, val)
}