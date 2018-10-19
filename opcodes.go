package gameboy

type Param interface{}

type Instruction func(...Param)

type Opcode byte

type Op struct {
	Instruction Instruction
	Params      []Param
	Cycles      []int
}

func NewOp(i Instruction, c []int, params ...Param) Op {
	return Op{
		Instruction: i,
		Params:      params,
		Cycles:      c,
	}
}
