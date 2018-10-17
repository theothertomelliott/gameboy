package gameboy

type Param interface{}

type Instruction func(...Param)

type Opcode byte

type Op struct {
	Instruction Instruction
	Params      []Param
	Cycles      int
}

func NewOp(i Instruction, c int, params ...Param) Op {
	return Op{
		Instruction: i,
		Params:      params,
		Cycles:      c,
	}
}

func unprefixedOpcodes(c *CPU) map[Opcode]Op {
	return map[Opcode]Op{
		// 0x0
		0x00: NewOp(c.NOP, 4),
		0x01: NewOp(c.LD, 12, c.BC, c.D16),
		0x02: NewOp(c.LD, 8, c.MemoryAt(c.BC), c.A),
		0x03: NewOp(c.INC, 8, c.BC),
		0x04: NewOp(c.INC, 4, c.B),
		0x05: NewOp(c.DEC, 4, c.B),
		0x06: NewOp(c.LD, 8, c.B, c.D8),
		0x07: NewOp(c.RCLA, 4),
		0x08: NewOp(c.LD, 20, c.MemoryAt(c.D16), c.SP),
		0x09: NewOp(c.ADDHL, 8, c.BC),
		0x0A: NewOp(c.LD, 8, c.A, c.BC),
		0x0B: NewOp(c.DEC, 8, c.BC),
		0x0C: NewOp(c.INC, 4, c.C),
		0x0D: NewOp(c.DEC, 4, c.C),
		0x0E: NewOp(c.LD, 8, c.C, c.D8),
		0x0F: NewOp(c.RRCA, 4),
	}
}
