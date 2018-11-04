package gameboy

import "fmt"

func cbprefixedOpcodes(c *CPU, code Opcode) Op {
	switch code {
	case 0x0:
		return NewOp("RLC B", c.RLC, []int{
			8,
		},
			c.B,
		)
	case 0x1:
		return NewOp("RLC C", c.RLC, []int{
			8,
		},
			c.C,
		)
	case 0x10:
		return NewOp("RL B", c.RL, []int{
			8,
		},
			c.B,
		)
	case 0x11:
		return NewOp("RL C", c.RL, []int{
			8,
		},
			c.C,
		)
	case 0x12:
		return NewOp("RL D", c.RL, []int{
			8,
		},
			c.D,
		)
	case 0x13:
		return NewOp("RL E", c.RL, []int{
			8,
		},
			c.E,
		)
	case 0x14:
		return NewOp("RL H", c.RL, []int{
			8,
		},
			c.H,
		)
	case 0x15:
		return NewOp("RL L", c.RL, []int{
			8,
		},
			c.L,
		)
	case 0x16:
		return NewOp("RL (HL)", c.RL, []int{
			16,
		},
			c.MemoryAt(c.HL),
		)
	case 0x17:
		return NewOp("RL A", c.RL, []int{
			8,
		},
			c.A,
		)
	case 0x18:
		return NewOp("RR B", c.RR, []int{
			8,
		},
			c.B,
		)
	case 0x19:
		return NewOp("RR C", c.RR, []int{
			8,
		},
			c.C,
		)
	case 0x1a:
		return NewOp("RR D", c.RR, []int{
			8,
		},
			c.D,
		)
	case 0x1b:
		return NewOp("RR E", c.RR, []int{
			8,
		},
			c.E,
		)
	case 0x1c:
		return NewOp("RR H", c.RR, []int{
			8,
		},
			c.H,
		)
	case 0x1d:
		return NewOp("RR L", c.RR, []int{
			8,
		},
			c.L,
		)
	case 0x1e:
		return NewOp("RR (HL)", c.RR, []int{
			16,
		},
			c.MemoryAt(c.HL),
		)
	case 0x1f:
		return NewOp("RR A", c.RR, []int{
			8,
		},
			c.A,
		)
	case 0x2:
		return NewOp("RLC D", c.RLC, []int{
			8,
		},
			c.D,
		)
	case 0x20:
		return NewOp("SLA B", c.SLA, []int{
			8,
		},
			c.B,
		)
	case 0x21:
		return NewOp("SLA C", c.SLA, []int{
			8,
		},
			c.C,
		)
	case 0x22:
		return NewOp("SLA D", c.SLA, []int{
			8,
		},
			c.D,
		)
	case 0x23:
		return NewOp("SLA E", c.SLA, []int{
			8,
		},
			c.E,
		)
	case 0x24:
		return NewOp("SLA H", c.SLA, []int{
			8,
		},
			c.H,
		)
	case 0x25:
		return NewOp("SLA L", c.SLA, []int{
			8,
		},
			c.L,
		)
	case 0x26:
		return NewOp("SLA (HL)", c.SLA, []int{
			16,
		},
			c.MemoryAt(c.HL),
		)
	case 0x27:
		return NewOp("SLA A", c.SLA, []int{
			8,
		},
			c.A,
		)
	case 0x28:
		return NewOp("SRA B", c.SRA, []int{
			8,
		},
			c.B,
		)
	case 0x29:
		return NewOp("SRA C", c.SRA, []int{
			8,
		},
			c.C,
		)
	case 0x2a:
		return NewOp("SRA D", c.SRA, []int{
			8,
		},
			c.D,
		)
	case 0x2b:
		return NewOp("SRA E", c.SRA, []int{
			8,
		},
			c.E,
		)
	case 0x2c:
		return NewOp("SRA H", c.SRA, []int{
			8,
		},
			c.H,
		)
	case 0x2d:
		return NewOp("SRA L", c.SRA, []int{
			8,
		},
			c.L,
		)
	case 0x2e:
		return NewOp("SRA (HL)", c.SRA, []int{
			16,
		},
			c.MemoryAt(c.HL),
		)
	case 0x2f:
		return NewOp("SRA A", c.SRA, []int{
			8,
		},
			c.A,
		)
	case 0x3:
		return NewOp("RLC E", c.RLC, []int{
			8,
		},
			c.E,
		)
	case 0x30:
		return NewOp("SWAP B", c.SWAP, []int{
			8,
		},
			c.B,
		)
	case 0x31:
		return NewOp("SWAP C", c.SWAP, []int{
			8,
		},
			c.C,
		)
	case 0x32:
		return NewOp("SWAP D", c.SWAP, []int{
			8,
		},
			c.D,
		)
	case 0x33:
		return NewOp("SWAP E", c.SWAP, []int{
			8,
		},
			c.E,
		)
	case 0x34:
		return NewOp("SWAP H", c.SWAP, []int{
			8,
		},
			c.H,
		)
	case 0x35:
		return NewOp("SWAP L", c.SWAP, []int{
			8,
		},
			c.L,
		)
	case 0x36:
		return NewOp("SWAP (HL)", c.SWAP, []int{
			16,
		},
			c.MemoryAt(c.HL),
		)
	case 0x37:
		return NewOp("SWAP A", c.SWAP, []int{
			8,
		},
			c.A,
		)
	case 0x38:
		return NewOp("SRL B", c.SRL, []int{
			8,
		},
			c.B,
		)
	case 0x39:
		return NewOp("SRL C", c.SRL, []int{
			8,
		},
			c.C,
		)
	case 0x3a:
		return NewOp("SRL D", c.SRL, []int{
			8,
		},
			c.D,
		)
	case 0x3b:
		return NewOp("SRL E", c.SRL, []int{
			8,
		},
			c.E,
		)
	case 0x3c:
		return NewOp("SRL H", c.SRL, []int{
			8,
		},
			c.H,
		)
	case 0x3d:
		return NewOp("SRL L", c.SRL, []int{
			8,
		},
			c.L,
		)
	case 0x3e:
		return NewOp("SRL (HL)", c.SRL, []int{
			16,
		},
			c.MemoryAt(c.HL),
		)
	case 0x3f:
		return NewOp("SRL A", c.SRL, []int{
			8,
		},
			c.A,
		)
	case 0x4:
		return NewOp("RLC H", c.RLC, []int{
			8,
		},
			c.H,
		)
	case 0x40:
		return NewOp("BIT 0,B", c.BIT, []int{
			8,
		},
			0,
			c.B,
		)
	case 0x41:
		return NewOp("BIT 0,C", c.BIT, []int{
			8,
		},
			0,
			c.C,
		)
	case 0x42:
		return NewOp("BIT 0,D", c.BIT, []int{
			8,
		},
			0,
			c.D,
		)
	case 0x43:
		return NewOp("BIT 0,E", c.BIT, []int{
			8,
		},
			0,
			c.E,
		)
	case 0x44:
		return NewOp("BIT 0,H", c.BIT, []int{
			8,
		},
			0,
			c.H,
		)
	case 0x45:
		return NewOp("BIT 0,L", c.BIT, []int{
			8,
		},
			0,
			c.L,
		)
	case 0x46:
		return NewOp("BIT 0,(HL)", c.BIT, []int{
			16,
		},
			0,
			c.MemoryAt(c.HL),
		)
	case 0x47:
		return NewOp("BIT 0,A", c.BIT, []int{
			8,
		},
			0,
			c.A,
		)
	case 0x48:
		return NewOp("BIT 1,B", c.BIT, []int{
			8,
		},
			1,
			c.B,
		)
	case 0x49:
		return NewOp("BIT 1,C", c.BIT, []int{
			8,
		},
			1,
			c.C,
		)
	case 0x4a:
		return NewOp("BIT 1,D", c.BIT, []int{
			8,
		},
			1,
			c.D,
		)
	case 0x4b:
		return NewOp("BIT 1,E", c.BIT, []int{
			8,
		},
			1,
			c.E,
		)
	case 0x4c:
		return NewOp("BIT 1,H", c.BIT, []int{
			8,
		},
			1,
			c.H,
		)
	case 0x4d:
		return NewOp("BIT 1,L", c.BIT, []int{
			8,
		},
			1,
			c.L,
		)
	case 0x4e:
		return NewOp("BIT 1,(HL)", c.BIT, []int{
			16,
		},
			1,
			c.MemoryAt(c.HL),
		)
	case 0x4f:
		return NewOp("BIT 1,A", c.BIT, []int{
			8,
		},
			1,
			c.A,
		)
	case 0x5:
		return NewOp("RLC L", c.RLC, []int{
			8,
		},
			c.L,
		)
	case 0x50:
		return NewOp("BIT 2,B", c.BIT, []int{
			8,
		},
			2,
			c.B,
		)
	case 0x51:
		return NewOp("BIT 2,C", c.BIT, []int{
			8,
		},
			2,
			c.C,
		)
	case 0x52:
		return NewOp("BIT 2,D", c.BIT, []int{
			8,
		},
			2,
			c.D,
		)
	case 0x53:
		return NewOp("BIT 2,E", c.BIT, []int{
			8,
		},
			2,
			c.E,
		)
	case 0x54:
		return NewOp("BIT 2,H", c.BIT, []int{
			8,
		},
			2,
			c.H,
		)
	case 0x55:
		return NewOp("BIT 2,L", c.BIT, []int{
			8,
		},
			2,
			c.L,
		)
	case 0x56:
		return NewOp("BIT 2,(HL)", c.BIT, []int{
			16,
		},
			2,
			c.MemoryAt(c.HL),
		)
	case 0x57:
		return NewOp("BIT 2,A", c.BIT, []int{
			8,
		},
			2,
			c.A,
		)
	case 0x58:
		return NewOp("BIT 3,B", c.BIT, []int{
			8,
		},
			3,
			c.B,
		)
	case 0x59:
		return NewOp("BIT 3,C", c.BIT, []int{
			8,
		},
			3,
			c.C,
		)
	case 0x5a:
		return NewOp("BIT 3,D", c.BIT, []int{
			8,
		},
			3,
			c.D,
		)
	case 0x5b:
		return NewOp("BIT 3,E", c.BIT, []int{
			8,
		},
			3,
			c.E,
		)
	case 0x5c:
		return NewOp("BIT 3,H", c.BIT, []int{
			8,
		},
			3,
			c.H,
		)
	case 0x5d:
		return NewOp("BIT 3,L", c.BIT, []int{
			8,
		},
			3,
			c.L,
		)
	case 0x5e:
		return NewOp("BIT 3,(HL)", c.BIT, []int{
			16,
		},
			3,
			c.MemoryAt(c.HL),
		)
	case 0x5f:
		return NewOp("BIT 3,A", c.BIT, []int{
			8,
		},
			3,
			c.A,
		)
	case 0x6:
		return NewOp("RLC (HL)", c.RLC, []int{
			16,
		},
			c.MemoryAt(c.HL),
		)
	case 0x60:
		return NewOp("BIT 4,B", c.BIT, []int{
			8,
		},
			4,
			c.B,
		)
	case 0x61:
		return NewOp("BIT 4,C", c.BIT, []int{
			8,
		},
			4,
			c.C,
		)
	case 0x62:
		return NewOp("BIT 4,D", c.BIT, []int{
			8,
		},
			4,
			c.D,
		)
	case 0x63:
		return NewOp("BIT 4,E", c.BIT, []int{
			8,
		},
			4,
			c.E,
		)
	case 0x64:
		return NewOp("BIT 4,H", c.BIT, []int{
			8,
		},
			4,
			c.H,
		)
	case 0x65:
		return NewOp("BIT 4,L", c.BIT, []int{
			8,
		},
			4,
			c.L,
		)
	case 0x66:
		return NewOp("BIT 4,(HL)", c.BIT, []int{
			16,
		},
			4,
			c.MemoryAt(c.HL),
		)
	case 0x67:
		return NewOp("BIT 4,A", c.BIT, []int{
			8,
		},
			4,
			c.A,
		)
	case 0x68:
		return NewOp("BIT 5,B", c.BIT, []int{
			8,
		},
			5,
			c.B,
		)
	case 0x69:
		return NewOp("BIT 5,C", c.BIT, []int{
			8,
		},
			5,
			c.C,
		)
	case 0x6a:
		return NewOp("BIT 5,D", c.BIT, []int{
			8,
		},
			5,
			c.D,
		)
	case 0x6b:
		return NewOp("BIT 5,E", c.BIT, []int{
			8,
		},
			5,
			c.E,
		)
	case 0x6c:
		return NewOp("BIT 5,H", c.BIT, []int{
			8,
		},
			5,
			c.H,
		)
	case 0x6d:
		return NewOp("BIT 5,L", c.BIT, []int{
			8,
		},
			5,
			c.L,
		)
	case 0x6e:
		return NewOp("BIT 5,(HL)", c.BIT, []int{
			16,
		},
			5,
			c.MemoryAt(c.HL),
		)
	case 0x6f:
		return NewOp("BIT 5,A", c.BIT, []int{
			8,
		},
			5,
			c.A,
		)
	case 0x7:
		return NewOp("RLC A", c.RLC, []int{
			8,
		},
			c.A,
		)
	case 0x70:
		return NewOp("BIT 6,B", c.BIT, []int{
			8,
		},
			6,
			c.B,
		)
	case 0x71:
		return NewOp("BIT 6,C", c.BIT, []int{
			8,
		},
			6,
			c.C,
		)
	case 0x72:
		return NewOp("BIT 6,D", c.BIT, []int{
			8,
		},
			6,
			c.D,
		)
	case 0x73:
		return NewOp("BIT 6,E", c.BIT, []int{
			8,
		},
			6,
			c.E,
		)
	case 0x74:
		return NewOp("BIT 6,H", c.BIT, []int{
			8,
		},
			6,
			c.H,
		)
	case 0x75:
		return NewOp("BIT 6,L", c.BIT, []int{
			8,
		},
			6,
			c.L,
		)
	case 0x76:
		return NewOp("BIT 6,(HL)", c.BIT, []int{
			16,
		},
			6,
			c.MemoryAt(c.HL),
		)
	case 0x77:
		return NewOp("BIT 6,A", c.BIT, []int{
			8,
		},
			6,
			c.A,
		)
	case 0x78:
		return NewOp("BIT 7,B", c.BIT, []int{
			8,
		},
			7,
			c.B,
		)
	case 0x79:
		return NewOp("BIT 7,C", c.BIT, []int{
			8,
		},
			7,
			c.C,
		)
	case 0x7a:
		return NewOp("BIT 7,D", c.BIT, []int{
			8,
		},
			7,
			c.D,
		)
	case 0x7b:
		return NewOp("BIT 7,E", c.BIT, []int{
			8,
		},
			7,
			c.E,
		)
	case 0x7c:
		return NewOp("BIT 7,H", c.BIT, []int{
			8,
		},
			7,
			c.H,
		)
	case 0x7d:
		return NewOp("BIT 7,L", c.BIT, []int{
			8,
		},
			7,
			c.L,
		)
	case 0x7e:
		return NewOp("BIT 7,(HL)", c.BIT, []int{
			16,
		},
			7,
			c.MemoryAt(c.HL),
		)
	case 0x7f:
		return NewOp("BIT 7,A", c.BIT, []int{
			8,
		},
			7,
			c.A,
		)
	case 0x8:
		return NewOp("RRC B", c.RRC, []int{
			8,
		},
			c.B,
		)
	case 0x80:
		return NewOp("RES 0,B", c.RES, []int{
			8,
		},
			0,
			c.B,
		)
	case 0x81:
		return NewOp("RES 0,C", c.RES, []int{
			8,
		},
			0,
			c.C,
		)
	case 0x82:
		return NewOp("RES 0,D", c.RES, []int{
			8,
		},
			0,
			c.D,
		)
	case 0x83:
		return NewOp("RES 0,E", c.RES, []int{
			8,
		},
			0,
			c.E,
		)
	case 0x84:
		return NewOp("RES 0,H", c.RES, []int{
			8,
		},
			0,
			c.H,
		)
	case 0x85:
		return NewOp("RES 0,L", c.RES, []int{
			8,
		},
			0,
			c.L,
		)
	case 0x86:
		return NewOp("RES 0,(HL)", c.RES, []int{
			16,
		},
			0,
			c.MemoryAt(c.HL),
		)
	case 0x87:
		return NewOp("RES 0,A", c.RES, []int{
			8,
		},
			0,
			c.A,
		)
	case 0x88:
		return NewOp("RES 1,B", c.RES, []int{
			8,
		},
			1,
			c.B,
		)
	case 0x89:
		return NewOp("RES 1,C", c.RES, []int{
			8,
		},
			1,
			c.C,
		)
	case 0x8a:
		return NewOp("RES 1,D", c.RES, []int{
			8,
		},
			1,
			c.D,
		)
	case 0x8b:
		return NewOp("RES 1,E", c.RES, []int{
			8,
		},
			1,
			c.E,
		)
	case 0x8c:
		return NewOp("RES 1,H", c.RES, []int{
			8,
		},
			1,
			c.H,
		)
	case 0x8d:
		return NewOp("RES 1,L", c.RES, []int{
			8,
		},
			1,
			c.L,
		)
	case 0x8e:
		return NewOp("RES 1,(HL)", c.RES, []int{
			16,
		},
			1,
			c.MemoryAt(c.HL),
		)
	case 0x8f:
		return NewOp("RES 1,A", c.RES, []int{
			8,
		},
			1,
			c.A,
		)
	case 0x9:
		return NewOp("RRC C", c.RRC, []int{
			8,
		},
			c.C,
		)
	case 0x90:
		return NewOp("RES 2,B", c.RES, []int{
			8,
		},
			2,
			c.B,
		)
	case 0x91:
		return NewOp("RES 2,C", c.RES, []int{
			8,
		},
			2,
			c.C,
		)
	case 0x92:
		return NewOp("RES 2,D", c.RES, []int{
			8,
		},
			2,
			c.D,
		)
	case 0x93:
		return NewOp("RES 2,E", c.RES, []int{
			8,
		},
			2,
			c.E,
		)
	case 0x94:
		return NewOp("RES 2,H", c.RES, []int{
			8,
		},
			2,
			c.H,
		)
	case 0x95:
		return NewOp("RES 2,L", c.RES, []int{
			8,
		},
			2,
			c.L,
		)
	case 0x96:
		return NewOp("RES 2,(HL)", c.RES, []int{
			16,
		},
			2,
			c.MemoryAt(c.HL),
		)
	case 0x97:
		return NewOp("RES 2,A", c.RES, []int{
			8,
		},
			2,
			c.A,
		)
	case 0x98:
		return NewOp("RES 3,B", c.RES, []int{
			8,
		},
			3,
			c.B,
		)
	case 0x99:
		return NewOp("RES 3,C", c.RES, []int{
			8,
		},
			3,
			c.C,
		)
	case 0x9a:
		return NewOp("RES 3,D", c.RES, []int{
			8,
		},
			3,
			c.D,
		)
	case 0x9b:
		return NewOp("RES 3,E", c.RES, []int{
			8,
		},
			3,
			c.E,
		)
	case 0x9c:
		return NewOp("RES 3,H", c.RES, []int{
			8,
		},
			3,
			c.H,
		)
	case 0x9d:
		return NewOp("RES 3,L", c.RES, []int{
			8,
		},
			3,
			c.L,
		)
	case 0x9e:
		return NewOp("RES 3,(HL)", c.RES, []int{
			16,
		},
			3,
			c.MemoryAt(c.HL),
		)
	case 0x9f:
		return NewOp("RES 3,A", c.RES, []int{
			8,
		},
			3,
			c.A,
		)
	case 0xa:
		return NewOp("RRC D", c.RRC, []int{
			8,
		},
			c.D,
		)
	case 0xa0:
		return NewOp("RES 4,B", c.RES, []int{
			8,
		},
			4,
			c.B,
		)
	case 0xa1:
		return NewOp("RES 4,C", c.RES, []int{
			8,
		},
			4,
			c.C,
		)
	case 0xa2:
		return NewOp("RES 4,D", c.RES, []int{
			8,
		},
			4,
			c.D,
		)
	case 0xa3:
		return NewOp("RES 4,E", c.RES, []int{
			8,
		},
			4,
			c.E,
		)
	case 0xa4:
		return NewOp("RES 4,H", c.RES, []int{
			8,
		},
			4,
			c.H,
		)
	case 0xa5:
		return NewOp("RES 4,L", c.RES, []int{
			8,
		},
			4,
			c.L,
		)
	case 0xa6:
		return NewOp("RES 4,(HL)", c.RES, []int{
			16,
		},
			4,
			c.MemoryAt(c.HL),
		)
	case 0xa7:
		return NewOp("RES 4,A", c.RES, []int{
			8,
		},
			4,
			c.A,
		)
	case 0xa8:
		return NewOp("RES 5,B", c.RES, []int{
			8,
		},
			5,
			c.B,
		)
	case 0xa9:
		return NewOp("RES 5,C", c.RES, []int{
			8,
		},
			5,
			c.C,
		)
	case 0xaa:
		return NewOp("RES 5,D", c.RES, []int{
			8,
		},
			5,
			c.D,
		)
	case 0xab:
		return NewOp("RES 5,E", c.RES, []int{
			8,
		},
			5,
			c.E,
		)
	case 0xac:
		return NewOp("RES 5,H", c.RES, []int{
			8,
		},
			5,
			c.H,
		)
	case 0xad:
		return NewOp("RES 5,L", c.RES, []int{
			8,
		},
			5,
			c.L,
		)
	case 0xae:
		return NewOp("RES 5,(HL)", c.RES, []int{
			16,
		},
			5,
			c.MemoryAt(c.HL),
		)
	case 0xaf:
		return NewOp("RES 5,A", c.RES, []int{
			8,
		},
			5,
			c.A,
		)
	case 0xb:
		return NewOp("RRC E", c.RRC, []int{
			8,
		},
			c.E,
		)
	case 0xb0:
		return NewOp("RES 6,B", c.RES, []int{
			8,
		},
			6,
			c.B,
		)
	case 0xb1:
		return NewOp("RES 6,C", c.RES, []int{
			8,
		},
			6,
			c.C,
		)
	case 0xb2:
		return NewOp("RES 6,D", c.RES, []int{
			8,
		},
			6,
			c.D,
		)
	case 0xb3:
		return NewOp("RES 6,E", c.RES, []int{
			8,
		},
			6,
			c.E,
		)
	case 0xb4:
		return NewOp("RES 6,H", c.RES, []int{
			8,
		},
			6,
			c.H,
		)
	case 0xb5:
		return NewOp("RES 6,L", c.RES, []int{
			8,
		},
			6,
			c.L,
		)
	case 0xb6:
		return NewOp("RES 6,(HL)", c.RES, []int{
			16,
		},
			6,
			c.MemoryAt(c.HL),
		)
	case 0xb7:
		return NewOp("RES 6,A", c.RES, []int{
			8,
		},
			6,
			c.A,
		)
	case 0xb8:
		return NewOp("RES 7,B", c.RES, []int{
			8,
		},
			7,
			c.B,
		)
	case 0xb9:
		return NewOp("RES 7,C", c.RES, []int{
			8,
		},
			7,
			c.C,
		)
	case 0xba:
		return NewOp("RES 7,D", c.RES, []int{
			8,
		},
			7,
			c.D,
		)
	case 0xbb:
		return NewOp("RES 7,E", c.RES, []int{
			8,
		},
			7,
			c.E,
		)
	case 0xbc:
		return NewOp("RES 7,H", c.RES, []int{
			8,
		},
			7,
			c.H,
		)
	case 0xbd:
		return NewOp("RES 7,L", c.RES, []int{
			8,
		},
			7,
			c.L,
		)
	case 0xbe:
		return NewOp("RES 7,(HL)", c.RES, []int{
			16,
		},
			7,
			c.MemoryAt(c.HL),
		)
	case 0xbf:
		return NewOp("RES 7,A", c.RES, []int{
			8,
		},
			7,
			c.A,
		)
	case 0xc:
		return NewOp("RRC H", c.RRC, []int{
			8,
		},
			c.H,
		)
	case 0xc0:
		return NewOp("SET 0,B", c.SET, []int{
			8,
		},
			0,
			c.B,
		)
	case 0xc1:
		return NewOp("SET 0,C", c.SET, []int{
			8,
		},
			0,
			c.C,
		)
	case 0xc2:
		return NewOp("SET 0,D", c.SET, []int{
			8,
		},
			0,
			c.D,
		)
	case 0xc3:
		return NewOp("SET 0,E", c.SET, []int{
			8,
		},
			0,
			c.E,
		)
	case 0xc4:
		return NewOp("SET 0,H", c.SET, []int{
			8,
		},
			0,
			c.H,
		)
	case 0xc5:
		return NewOp("SET 0,L", c.SET, []int{
			8,
		},
			0,
			c.L,
		)
	case 0xc6:
		return NewOp("SET 0,(HL)", c.SET, []int{
			16,
		},
			0,
			c.MemoryAt(c.HL),
		)
	case 0xc7:
		return NewOp("SET 0,A", c.SET, []int{
			8,
		},
			0,
			c.A,
		)
	case 0xc8:
		return NewOp("SET 1,B", c.SET, []int{
			8,
		},
			1,
			c.B,
		)
	case 0xc9:
		return NewOp("SET 1,C", c.SET, []int{
			8,
		},
			1,
			c.C,
		)
	case 0xca:
		return NewOp("SET 1,D", c.SET, []int{
			8,
		},
			1,
			c.D,
		)
	case 0xcb:
		return NewOp("SET 1,E", c.SET, []int{
			8,
		},
			1,
			c.E,
		)
	case 0xcc:
		return NewOp("SET 1,H", c.SET, []int{
			8,
		},
			1,
			c.H,
		)
	case 0xcd:
		return NewOp("SET 1,L", c.SET, []int{
			8,
		},
			1,
			c.L,
		)
	case 0xce:
		return NewOp("SET 1,(HL)", c.SET, []int{
			16,
		},
			1,
			c.MemoryAt(c.HL),
		)
	case 0xcf:
		return NewOp("SET 1,A", c.SET, []int{
			8,
		},
			1,
			c.A,
		)
	case 0xd:
		return NewOp("RRC L", c.RRC, []int{
			8,
		},
			c.L,
		)
	case 0xd0:
		return NewOp("SET 2,B", c.SET, []int{
			8,
		},
			2,
			c.B,
		)
	case 0xd1:
		return NewOp("SET 2,C", c.SET, []int{
			8,
		},
			2,
			c.C,
		)
	case 0xd2:
		return NewOp("SET 2,D", c.SET, []int{
			8,
		},
			2,
			c.D,
		)
	case 0xd3:
		return NewOp("SET 2,E", c.SET, []int{
			8,
		},
			2,
			c.E,
		)
	case 0xd4:
		return NewOp("SET 2,H", c.SET, []int{
			8,
		},
			2,
			c.H,
		)
	case 0xd5:
		return NewOp("SET 2,L", c.SET, []int{
			8,
		},
			2,
			c.L,
		)
	case 0xd6:
		return NewOp("SET 2,(HL)", c.SET, []int{
			16,
		},
			2,
			c.MemoryAt(c.HL),
		)
	case 0xd7:
		return NewOp("SET 2,A", c.SET, []int{
			8,
		},
			2,
			c.A,
		)
	case 0xd8:
		return NewOp("SET 3,B", c.SET, []int{
			8,
		},
			3,
			c.B,
		)
	case 0xd9:
		return NewOp("SET 3,C", c.SET, []int{
			8,
		},
			3,
			c.C,
		)
	case 0xda:
		return NewOp("SET 3,D", c.SET, []int{
			8,
		},
			3,
			c.D,
		)
	case 0xdb:
		return NewOp("SET 3,E", c.SET, []int{
			8,
		},
			3,
			c.E,
		)
	case 0xdc:
		return NewOp("SET 3,H", c.SET, []int{
			8,
		},
			3,
			c.H,
		)
	case 0xdd:
		return NewOp("SET 3,L", c.SET, []int{
			8,
		},
			3,
			c.L,
		)
	case 0xde:
		return NewOp("SET 3,(HL)", c.SET, []int{
			16,
		},
			3,
			c.MemoryAt(c.HL),
		)
	case 0xdf:
		return NewOp("SET 3,A", c.SET, []int{
			8,
		},
			3,
			c.A,
		)
	case 0xe:
		return NewOp("RRC (HL)", c.RRC, []int{
			16,
		},
			c.MemoryAt(c.HL),
		)
	case 0xe0:
		return NewOp("SET 4,B", c.SET, []int{
			8,
		},
			4,
			c.B,
		)
	case 0xe1:
		return NewOp("SET 4,C", c.SET, []int{
			8,
		},
			4,
			c.C,
		)
	case 0xe2:
		return NewOp("SET 4,D", c.SET, []int{
			8,
		},
			4,
			c.D,
		)
	case 0xe3:
		return NewOp("SET 4,E", c.SET, []int{
			8,
		},
			4,
			c.E,
		)
	case 0xe4:
		return NewOp("SET 4,H", c.SET, []int{
			8,
		},
			4,
			c.H,
		)
	case 0xe5:
		return NewOp("SET 4,L", c.SET, []int{
			8,
		},
			4,
			c.L,
		)
	case 0xe6:
		return NewOp("SET 4,(HL)", c.SET, []int{
			16,
		},
			4,
			c.MemoryAt(c.HL),
		)
	case 0xe7:
		return NewOp("SET 4,A", c.SET, []int{
			8,
		},
			4,
			c.A,
		)
	case 0xe8:
		return NewOp("SET 5,B", c.SET, []int{
			8,
		},
			5,
			c.B,
		)
	case 0xe9:
		return NewOp("SET 5,C", c.SET, []int{
			8,
		},
			5,
			c.C,
		)
	case 0xea:
		return NewOp("SET 5,D", c.SET, []int{
			8,
		},
			5,
			c.D,
		)
	case 0xeb:
		return NewOp("SET 5,E", c.SET, []int{
			8,
		},
			5,
			c.E,
		)
	case 0xec:
		return NewOp("SET 5,H", c.SET, []int{
			8,
		},
			5,
			c.H,
		)
	case 0xed:
		return NewOp("SET 5,L", c.SET, []int{
			8,
		},
			5,
			c.L,
		)
	case 0xee:
		return NewOp("SET 5,(HL)", c.SET, []int{
			16,
		},
			5,
			c.MemoryAt(c.HL),
		)
	case 0xef:
		return NewOp("SET 5,A", c.SET, []int{
			8,
		},
			5,
			c.A,
		)
	case 0xf:
		return NewOp("RRC A", c.RRC, []int{
			8,
		},
			c.A,
		)
	case 0xf0:
		return NewOp("SET 6,B", c.SET, []int{
			8,
		},
			6,
			c.B,
		)
	case 0xf1:
		return NewOp("SET 6,C", c.SET, []int{
			8,
		},
			6,
			c.C,
		)
	case 0xf2:
		return NewOp("SET 6,D", c.SET, []int{
			8,
		},
			6,
			c.D,
		)
	case 0xf3:
		return NewOp("SET 6,E", c.SET, []int{
			8,
		},
			6,
			c.E,
		)
	case 0xf4:
		return NewOp("SET 6,H", c.SET, []int{
			8,
		},
			6,
			c.H,
		)
	case 0xf5:
		return NewOp("SET 6,L", c.SET, []int{
			8,
		},
			6,
			c.L,
		)
	case 0xf6:
		return NewOp("SET 6,(HL)", c.SET, []int{
			16,
		},
			6,
			c.MemoryAt(c.HL),
		)
	case 0xf7:
		return NewOp("SET 6,A", c.SET, []int{
			8,
		},
			6,
			c.A,
		)
	case 0xf8:
		return NewOp("SET 7,B", c.SET, []int{
			8,
		},
			7,
			c.B,
		)
	case 0xf9:
		return NewOp("SET 7,C", c.SET, []int{
			8,
		},
			7,
			c.C,
		)
	case 0xfa:
		return NewOp("SET 7,D", c.SET, []int{
			8,
		},
			7,
			c.D,
		)
	case 0xfb:
		return NewOp("SET 7,E", c.SET, []int{
			8,
		},
			7,
			c.E,
		)
	case 0xfc:
		return NewOp("SET 7,H", c.SET, []int{
			8,
		},
			7,
			c.H,
		)
	case 0xfd:
		return NewOp("SET 7,L", c.SET, []int{
			8,
		},
			7,
			c.L,
		)
	case 0xfe:
		return NewOp("SET 7,(HL)", c.SET, []int{
			16,
		},
			7,
			c.MemoryAt(c.HL),
		)
	case 0xff:
		return NewOp("SET 7,A", c.SET, []int{
			8,
		},
			7,
			c.A,
		)
	default:
		panic(fmt.Sprintf("unknown opcode: 0x%X", code))
	}
}
func unprefixedOpcodes(c *CPU, code Opcode) Op {
	switch code {
	case 0x0:
		return NewOp("NOP", c.NOP, []int{
			4,
		},
		)
	case 0x1:
		return NewOp("LD BC,d16", c.LD, []int{
			12,
		},
			c.BC,
			c.D16(),
		)
	case 0x10:
		return NewOp("STOP 0", c.STOP, []int{
			4,
		},
			0,
		)
	case 0x11:
		return NewOp("LD DE,d16", c.LD, []int{
			12,
		},
			c.DE,
			c.D16(),
		)
	case 0x12:
		return NewOp("LD (DE),A", c.LD, []int{
			8,
		},
			c.MemoryAt(c.DE),
			c.A,
		)
	case 0x13:
		return NewOp("INC DE", c.INC, []int{
			8,
		},
			c.DE,
		)
	case 0x14:
		return NewOp("INC D", c.INC, []int{
			4,
		},
			c.D,
		)
	case 0x15:
		return NewOp("DEC D", c.DEC, []int{
			4,
		},
			c.D,
		)
	case 0x16:
		return NewOp("LD D,d8", c.LD, []int{
			8,
		},
			c.D,
			c.D8(),
		)
	case 0x17:
		return NewOp("RLA", c.RLA, []int{
			4,
		},
		)
	case 0x18:
		return NewOp("JR r8", c.JR, []int{
			12,
		},
			c.R8(),
		)
	case 0x19:
		return NewOp("ADD HL,DE", c.ADD, []int{
			8,
		},
			c.HL,
			c.DE,
		)
	case 0x1a:
		return NewOp("LD A,(DE)", c.LD, []int{
			8,
		},
			c.A,
			c.MemoryAt(c.DE),
		)
	case 0x1b:
		return NewOp("DEC DE", c.DEC, []int{
			8,
		},
			c.DE,
		)
	case 0x1c:
		return NewOp("INC E", c.INC, []int{
			4,
		},
			c.E,
		)
	case 0x1d:
		return NewOp("DEC E", c.DEC, []int{
			4,
		},
			c.E,
		)
	case 0x1e:
		return NewOp("LD E,d8", c.LD, []int{
			8,
		},
			c.E,
			c.D8(),
		)
	case 0x1f:
		return NewOp("RRA", c.RRA, []int{
			4,
		},
		)
	case 0x2:
		return NewOp("LD (BC),A", c.LD, []int{
			8,
		},
			c.MemoryAt(c.BC),
			c.A,
		)
	case 0x20:
		return NewOp("JR NZ,r8", c.JRC, []int{
			12,

			8,
		},
			CaseNZ,
			c.R8(),
		)
	case 0x21:
		return NewOp("LD HL,d16", c.LD, []int{
			12,
		},
			c.HL,
			c.D16(),
		)
	case 0x22:
		return NewOp("LDI (HL),A", c.LDI, []int{
			8,
		},
			c.MemoryAt(c.HL),
			c.A,
		)
	case 0x23:
		return NewOp("INC HL", c.INC, []int{
			8,
		},
			c.HL,
		)
	case 0x24:
		return NewOp("INC H", c.INC, []int{
			4,
		},
			c.H,
		)
	case 0x25:
		return NewOp("DEC H", c.DEC, []int{
			4,
		},
			c.H,
		)
	case 0x26:
		return NewOp("LD H,d8", c.LD, []int{
			8,
		},
			c.H,
			c.D8(),
		)
	case 0x27:
		return NewOp("DAA", c.DAA, []int{
			4,
		},
		)
	case 0x28:
		return NewOp("JR Z,r8", c.JRC, []int{
			12,

			8,
		},
			CaseZ,
			c.R8(),
		)
	case 0x29:
		return NewOp("ADD HL,HL", c.ADD, []int{
			8,
		},
			c.HL,
			c.HL,
		)
	case 0x2a:
		return NewOp("LDI A,(HL)", c.LDI, []int{
			8,
		},
			c.A,
			c.MemoryAt(c.HL),
		)
	case 0x2b:
		return NewOp("DEC HL", c.DEC, []int{
			8,
		},
			c.HL,
		)
	case 0x2c:
		return NewOp("INC L", c.INC, []int{
			4,
		},
			c.L,
		)
	case 0x2d:
		return NewOp("DEC L", c.DEC, []int{
			4,
		},
			c.L,
		)
	case 0x2e:
		return NewOp("LD L,d8", c.LD, []int{
			8,
		},
			c.L,
			c.D8(),
		)
	case 0x2f:
		return NewOp("CPL", c.CPL, []int{
			4,
		},
		)
	case 0x3:
		return NewOp("INC BC", c.INC, []int{
			8,
		},
			c.BC,
		)
	case 0x30:
		return NewOp("JR NC,r8", c.JRC, []int{
			12,

			8,
		},
			CaseNC,
			c.R8(),
		)
	case 0x31:
		return NewOp("LD SP,d16", c.LD, []int{
			12,
		},
			c.SP,
			c.D16(),
		)
	case 0x32:
		return NewOp("LDD (HL),A", c.LDD, []int{
			8,
		},
			c.MemoryAt(c.HL),
			c.A,
		)
	case 0x33:
		return NewOp("INC SP", c.INC, []int{
			8,
		},
			c.SP,
		)
	case 0x34:
		return NewOp("INC (HL)", c.INC, []int{
			12,
		},
			c.MemoryAt(c.HL),
		)
	case 0x35:
		return NewOp("DEC (HL)", c.DEC, []int{
			12,
		},
			c.MemoryAt(c.HL),
		)
	case 0x36:
		return NewOp("LD (HL),d8", c.LD, []int{
			12,
		},
			c.MemoryAt(c.HL),
			c.D8(),
		)
	case 0x37:
		return NewOp("SCF", c.SCF, []int{
			4,
		},
		)
	case 0x38:
		return NewOp("JR C,r8", c.JRC, []int{
			12,

			8,
		},
			CaseC,
			c.R8(),
		)
	case 0x39:
		return NewOp("ADD HL,SP", c.ADD, []int{
			8,
		},
			c.HL,
			c.SP,
		)
	case 0x3a:
		return NewOp("LDD A,(HL)", c.LDD, []int{
			8,
		},
			c.A,
			c.MemoryAt(c.HL),
		)
	case 0x3b:
		return NewOp("DEC SP", c.DEC, []int{
			8,
		},
			c.SP,
		)
	case 0x3c:
		return NewOp("INC A", c.INC, []int{
			4,
		},
			c.A,
		)
	case 0x3d:
		return NewOp("DEC A", c.DEC, []int{
			4,
		},
			c.A,
		)
	case 0x3e:
		return NewOp("LD A,d8", c.LD, []int{
			8,
		},
			c.A,
			c.D8(),
		)
	case 0x3f:
		return NewOp("CCF", c.CCF, []int{
			4,
		},
		)
	case 0x4:
		return NewOp("INC B", c.INC, []int{
			4,
		},
			c.B,
		)
	case 0x40:
		return NewOp("LD B,B", c.LD, []int{
			4,
		},
			c.B,
			c.B,
		)
	case 0x41:
		return NewOp("LD B,C", c.LD, []int{
			4,
		},
			c.B,
			c.C,
		)
	case 0x42:
		return NewOp("LD B,D", c.LD, []int{
			4,
		},
			c.B,
			c.D,
		)
	case 0x43:
		return NewOp("LD B,E", c.LD, []int{
			4,
		},
			c.B,
			c.E,
		)
	case 0x44:
		return NewOp("LD B,H", c.LD, []int{
			4,
		},
			c.B,
			c.H,
		)
	case 0x45:
		return NewOp("LD B,L", c.LD, []int{
			4,
		},
			c.B,
			c.L,
		)
	case 0x46:
		return NewOp("LD B,(HL)", c.LD, []int{
			8,
		},
			c.B,
			c.MemoryAt(c.HL),
		)
	case 0x47:
		return NewOp("LD B,A", c.LD, []int{
			4,
		},
			c.B,
			c.A,
		)
	case 0x48:
		return NewOp("LD C,B", c.LD, []int{
			4,
		},
			c.C,
			c.B,
		)
	case 0x49:
		return NewOp("LD C,C", c.LD, []int{
			4,
		},
			c.C,
			c.C,
		)
	case 0x4a:
		return NewOp("LD C,D", c.LD, []int{
			4,
		},
			c.C,
			c.D,
		)
	case 0x4b:
		return NewOp("LD C,E", c.LD, []int{
			4,
		},
			c.C,
			c.E,
		)
	case 0x4c:
		return NewOp("LD C,H", c.LD, []int{
			4,
		},
			c.C,
			c.H,
		)
	case 0x4d:
		return NewOp("LD C,L", c.LD, []int{
			4,
		},
			c.C,
			c.L,
		)
	case 0x4e:
		return NewOp("LD C,(HL)", c.LD, []int{
			8,
		},
			c.C,
			c.MemoryAt(c.HL),
		)
	case 0x4f:
		return NewOp("LD C,A", c.LD, []int{
			4,
		},
			c.C,
			c.A,
		)
	case 0x5:
		return NewOp("DEC B", c.DEC, []int{
			4,
		},
			c.B,
		)
	case 0x50:
		return NewOp("LD D,B", c.LD, []int{
			4,
		},
			c.D,
			c.B,
		)
	case 0x51:
		return NewOp("LD D,C", c.LD, []int{
			4,
		},
			c.D,
			c.C,
		)
	case 0x52:
		return NewOp("LD D,D", c.LD, []int{
			4,
		},
			c.D,
			c.D,
		)
	case 0x53:
		return NewOp("LD D,E", c.LD, []int{
			4,
		},
			c.D,
			c.E,
		)
	case 0x54:
		return NewOp("LD D,H", c.LD, []int{
			4,
		},
			c.D,
			c.H,
		)
	case 0x55:
		return NewOp("LD D,L", c.LD, []int{
			4,
		},
			c.D,
			c.L,
		)
	case 0x56:
		return NewOp("LD D,(HL)", c.LD, []int{
			8,
		},
			c.D,
			c.MemoryAt(c.HL),
		)
	case 0x57:
		return NewOp("LD D,A", c.LD, []int{
			4,
		},
			c.D,
			c.A,
		)
	case 0x58:
		return NewOp("LD E,B", c.LD, []int{
			4,
		},
			c.E,
			c.B,
		)
	case 0x59:
		return NewOp("LD E,C", c.LD, []int{
			4,
		},
			c.E,
			c.C,
		)
	case 0x5a:
		return NewOp("LD E,D", c.LD, []int{
			4,
		},
			c.E,
			c.D,
		)
	case 0x5b:
		return NewOp("LD E,E", c.LD, []int{
			4,
		},
			c.E,
			c.E,
		)
	case 0x5c:
		return NewOp("LD E,H", c.LD, []int{
			4,
		},
			c.E,
			c.H,
		)
	case 0x5d:
		return NewOp("LD E,L", c.LD, []int{
			4,
		},
			c.E,
			c.L,
		)
	case 0x5e:
		return NewOp("LD E,(HL)", c.LD, []int{
			8,
		},
			c.E,
			c.MemoryAt(c.HL),
		)
	case 0x5f:
		return NewOp("LD E,A", c.LD, []int{
			4,
		},
			c.E,
			c.A,
		)
	case 0x6:
		return NewOp("LD B,d8", c.LD, []int{
			8,
		},
			c.B,
			c.D8(),
		)
	case 0x60:
		return NewOp("LD H,B", c.LD, []int{
			4,
		},
			c.H,
			c.B,
		)
	case 0x61:
		return NewOp("LD H,C", c.LD, []int{
			4,
		},
			c.H,
			c.C,
		)
	case 0x62:
		return NewOp("LD H,D", c.LD, []int{
			4,
		},
			c.H,
			c.D,
		)
	case 0x63:
		return NewOp("LD H,E", c.LD, []int{
			4,
		},
			c.H,
			c.E,
		)
	case 0x64:
		return NewOp("LD H,H", c.LD, []int{
			4,
		},
			c.H,
			c.H,
		)
	case 0x65:
		return NewOp("LD H,L", c.LD, []int{
			4,
		},
			c.H,
			c.L,
		)
	case 0x66:
		return NewOp("LD H,(HL)", c.LD, []int{
			8,
		},
			c.H,
			c.MemoryAt(c.HL),
		)
	case 0x67:
		return NewOp("LD H,A", c.LD, []int{
			4,
		},
			c.H,
			c.A,
		)
	case 0x68:
		return NewOp("LD L,B", c.LD, []int{
			4,
		},
			c.L,
			c.B,
		)
	case 0x69:
		return NewOp("LD L,C", c.LD, []int{
			4,
		},
			c.L,
			c.C,
		)
	case 0x6a:
		return NewOp("LD L,D", c.LD, []int{
			4,
		},
			c.L,
			c.D,
		)
	case 0x6b:
		return NewOp("LD L,E", c.LD, []int{
			4,
		},
			c.L,
			c.E,
		)
	case 0x6c:
		return NewOp("LD L,H", c.LD, []int{
			4,
		},
			c.L,
			c.H,
		)
	case 0x6d:
		return NewOp("LD L,L", c.LD, []int{
			4,
		},
			c.L,
			c.L,
		)
	case 0x6e:
		return NewOp("LD L,(HL)", c.LD, []int{
			8,
		},
			c.L,
			c.MemoryAt(c.HL),
		)
	case 0x6f:
		return NewOp("LD L,A", c.LD, []int{
			4,
		},
			c.L,
			c.A,
		)
	case 0x7:
		return NewOp("RLCA", c.RLCA, []int{
			4,
		},
		)
	case 0x70:
		return NewOp("LD (HL),B", c.LD, []int{
			8,
		},
			c.MemoryAt(c.HL),
			c.B,
		)
	case 0x71:
		return NewOp("LD (HL),C", c.LD, []int{
			8,
		},
			c.MemoryAt(c.HL),
			c.C,
		)
	case 0x72:
		return NewOp("LD (HL),D", c.LD, []int{
			8,
		},
			c.MemoryAt(c.HL),
			c.D,
		)
	case 0x73:
		return NewOp("LD (HL),E", c.LD, []int{
			8,
		},
			c.MemoryAt(c.HL),
			c.E,
		)
	case 0x74:
		return NewOp("LD (HL),H", c.LD, []int{
			8,
		},
			c.MemoryAt(c.HL),
			c.H,
		)
	case 0x75:
		return NewOp("LD (HL),L", c.LD, []int{
			8,
		},
			c.MemoryAt(c.HL),
			c.L,
		)
	case 0x76:
		return NewOp("HALT", c.HALT, []int{
			4,
		},
		)
	case 0x77:
		return NewOp("LD (HL),A", c.LD, []int{
			8,
		},
			c.MemoryAt(c.HL),
			c.A,
		)
	case 0x78:
		return NewOp("LD A,B", c.LD, []int{
			4,
		},
			c.A,
			c.B,
		)
	case 0x79:
		return NewOp("LD A,C", c.LD, []int{
			4,
		},
			c.A,
			c.C,
		)
	case 0x7a:
		return NewOp("LD A,D", c.LD, []int{
			4,
		},
			c.A,
			c.D,
		)
	case 0x7b:
		return NewOp("LD A,E", c.LD, []int{
			4,
		},
			c.A,
			c.E,
		)
	case 0x7c:
		return NewOp("LD A,H", c.LD, []int{
			4,
		},
			c.A,
			c.H,
		)
	case 0x7d:
		return NewOp("LD A,L", c.LD, []int{
			4,
		},
			c.A,
			c.L,
		)
	case 0x7e:
		return NewOp("LD A,(HL)", c.LD, []int{
			8,
		},
			c.A,
			c.MemoryAt(c.HL),
		)
	case 0x7f:
		return NewOp("LD A,A", c.LD, []int{
			4,
		},
			c.A,
			c.A,
		)
	case 0x8:
		return NewOp("LD (a16),SP", c.LD, []int{
			20,
		},
			c.MemoryAt(c.A16()),
			c.SP,
		)
	case 0x80:
		return NewOp("ADD A,B", c.ADD, []int{
			4,
		},
			c.A,
			c.B,
		)
	case 0x81:
		return NewOp("ADD A,C", c.ADD, []int{
			4,
		},
			c.A,
			c.C,
		)
	case 0x82:
		return NewOp("ADD A,D", c.ADD, []int{
			4,
		},
			c.A,
			c.D,
		)
	case 0x83:
		return NewOp("ADD A,E", c.ADD, []int{
			4,
		},
			c.A,
			c.E,
		)
	case 0x84:
		return NewOp("ADD A,H", c.ADD, []int{
			4,
		},
			c.A,
			c.H,
		)
	case 0x85:
		return NewOp("ADD A,L", c.ADD, []int{
			4,
		},
			c.A,
			c.L,
		)
	case 0x86:
		return NewOp("ADD A,(HL)", c.ADD, []int{
			8,
		},
			c.A,
			c.MemoryAt(c.HL),
		)
	case 0x87:
		return NewOp("ADD A,A", c.ADD, []int{
			4,
		},
			c.A,
			c.A,
		)
	case 0x88:
		return NewOp("ADC A,B", c.ADC, []int{
			4,
		},
			c.A,
			c.B,
		)
	case 0x89:
		return NewOp("ADC A,C", c.ADC, []int{
			4,
		},
			c.A,
			c.C,
		)
	case 0x8a:
		return NewOp("ADC A,D", c.ADC, []int{
			4,
		},
			c.A,
			c.D,
		)
	case 0x8b:
		return NewOp("ADC A,E", c.ADC, []int{
			4,
		},
			c.A,
			c.E,
		)
	case 0x8c:
		return NewOp("ADC A,H", c.ADC, []int{
			4,
		},
			c.A,
			c.H,
		)
	case 0x8d:
		return NewOp("ADC A,L", c.ADC, []int{
			4,
		},
			c.A,
			c.L,
		)
	case 0x8e:
		return NewOp("ADC A,(HL)", c.ADC, []int{
			8,
		},
			c.A,
			c.MemoryAt(c.HL),
		)
	case 0x8f:
		return NewOp("ADC A,A", c.ADC, []int{
			4,
		},
			c.A,
			c.A,
		)
	case 0x9:
		return NewOp("ADD HL,BC", c.ADD, []int{
			8,
		},
			c.HL,
			c.BC,
		)
	case 0x90:
		return NewOp("SUB B", c.SUB, []int{
			4,
		},
			c.B,
		)
	case 0x91:
		return NewOp("SUB C", c.SUB, []int{
			4,
		},
			c.C,
		)
	case 0x92:
		return NewOp("SUB D", c.SUB, []int{
			4,
		},
			c.D,
		)
	case 0x93:
		return NewOp("SUB E", c.SUB, []int{
			4,
		},
			c.E,
		)
	case 0x94:
		return NewOp("SUB H", c.SUB, []int{
			4,
		},
			c.H,
		)
	case 0x95:
		return NewOp("SUB L", c.SUB, []int{
			4,
		},
			c.L,
		)
	case 0x96:
		return NewOp("SUB (HL)", c.SUB, []int{
			8,
		},
			c.MemoryAt(c.HL),
		)
	case 0x97:
		return NewOp("SUB A", c.SUB, []int{
			4,
		},
			c.A,
		)
	case 0x98:
		return NewOp("SBC A,B", c.SBC, []int{
			4,
		},
			c.A,
			c.B,
		)
	case 0x99:
		return NewOp("SBC A,C", c.SBC, []int{
			4,
		},
			c.A,
			c.C,
		)
	case 0x9a:
		return NewOp("SBC A,D", c.SBC, []int{
			4,
		},
			c.A,
			c.D,
		)
	case 0x9b:
		return NewOp("SBC A,E", c.SBC, []int{
			4,
		},
			c.A,
			c.E,
		)
	case 0x9c:
		return NewOp("SBC A,H", c.SBC, []int{
			4,
		},
			c.A,
			c.H,
		)
	case 0x9d:
		return NewOp("SBC A,L", c.SBC, []int{
			4,
		},
			c.A,
			c.L,
		)
	case 0x9e:
		return NewOp("SBC A,(HL)", c.SBC, []int{
			8,
		},
			c.A,
			c.MemoryAt(c.HL),
		)
	case 0x9f:
		return NewOp("SBC A,A", c.SBC, []int{
			4,
		},
			c.A,
			c.A,
		)
	case 0xa:
		return NewOp("LD A,(BC)", c.LD, []int{
			8,
		},
			c.A,
			c.MemoryAt(c.BC),
		)
	case 0xa0:
		return NewOp("AND B", c.AND, []int{
			4,
		},
			c.B,
		)
	case 0xa1:
		return NewOp("AND C", c.AND, []int{
			4,
		},
			c.C,
		)
	case 0xa2:
		return NewOp("AND D", c.AND, []int{
			4,
		},
			c.D,
		)
	case 0xa3:
		return NewOp("AND E", c.AND, []int{
			4,
		},
			c.E,
		)
	case 0xa4:
		return NewOp("AND H", c.AND, []int{
			4,
		},
			c.H,
		)
	case 0xa5:
		return NewOp("AND L", c.AND, []int{
			4,
		},
			c.L,
		)
	case 0xa6:
		return NewOp("AND (HL)", c.AND, []int{
			8,
		},
			c.MemoryAt(c.HL),
		)
	case 0xa7:
		return NewOp("AND A", c.AND, []int{
			4,
		},
			c.A,
		)
	case 0xa8:
		return NewOp("XOR B", c.XOR, []int{
			4,
		},
			c.B,
		)
	case 0xa9:
		return NewOp("XOR C", c.XOR, []int{
			4,
		},
			c.C,
		)
	case 0xaa:
		return NewOp("XOR D", c.XOR, []int{
			4,
		},
			c.D,
		)
	case 0xab:
		return NewOp("XOR E", c.XOR, []int{
			4,
		},
			c.E,
		)
	case 0xac:
		return NewOp("XOR H", c.XOR, []int{
			4,
		},
			c.H,
		)
	case 0xad:
		return NewOp("XOR L", c.XOR, []int{
			4,
		},
			c.L,
		)
	case 0xae:
		return NewOp("XOR (HL)", c.XOR, []int{
			8,
		},
			c.MemoryAt(c.HL),
		)
	case 0xaf:
		return NewOp("XOR A", c.XOR, []int{
			4,
		},
			c.A,
		)
	case 0xb:
		return NewOp("DEC BC", c.DEC, []int{
			8,
		},
			c.BC,
		)
	case 0xb0:
		return NewOp("OR B", c.OR, []int{
			4,
		},
			c.B,
		)
	case 0xb1:
		return NewOp("OR C", c.OR, []int{
			4,
		},
			c.C,
		)
	case 0xb2:
		return NewOp("OR D", c.OR, []int{
			4,
		},
			c.D,
		)
	case 0xb3:
		return NewOp("OR E", c.OR, []int{
			4,
		},
			c.E,
		)
	case 0xb4:
		return NewOp("OR H", c.OR, []int{
			4,
		},
			c.H,
		)
	case 0xb5:
		return NewOp("OR L", c.OR, []int{
			4,
		},
			c.L,
		)
	case 0xb6:
		return NewOp("OR (HL)", c.OR, []int{
			8,
		},
			c.MemoryAt(c.HL),
		)
	case 0xb7:
		return NewOp("OR A", c.OR, []int{
			4,
		},
			c.A,
		)
	case 0xb8:
		return NewOp("CP B", c.CP, []int{
			4,
		},
			c.B,
		)
	case 0xb9:
		return NewOp("CP C", c.CP, []int{
			4,
		},
			c.C,
		)
	case 0xba:
		return NewOp("CP D", c.CP, []int{
			4,
		},
			c.D,
		)
	case 0xbb:
		return NewOp("CP E", c.CP, []int{
			4,
		},
			c.E,
		)
	case 0xbc:
		return NewOp("CP H", c.CP, []int{
			4,
		},
			c.H,
		)
	case 0xbd:
		return NewOp("CP L", c.CP, []int{
			4,
		},
			c.L,
		)
	case 0xbe:
		return NewOp("CP (HL)", c.CP, []int{
			8,
		},
			c.MemoryAt(c.HL),
		)
	case 0xbf:
		return NewOp("CP A", c.CP, []int{
			4,
		},
			c.A,
		)
	case 0xc:
		return NewOp("INC C", c.INC, []int{
			4,
		},
			c.C,
		)
	case 0xc0:
		return NewOp("RET NZ", c.RETC, []int{
			20,

			8,
		},
			CaseNZ,
		)
	case 0xc1:
		return NewOp("POP BC", c.POP, []int{
			12,
		},
			c.BC,
		)
	case 0xc2:
		return NewOp("JP NZ,a16", c.JPC, []int{
			16,

			12,
		},
			CaseNZ,
			c.A16(),
		)
	case 0xc3:
		return NewOp("JP a16", c.JP, []int{
			16,
		},
			c.A16(),
		)
	case 0xc4:
		return NewOp("CALL NZ,a16", c.CALLC, []int{
			24,

			12,
		},
			CaseNZ,
			c.A16(),
		)
	case 0xc5:
		return NewOp("PUSH BC", c.PUSH, []int{
			16,
		},
			c.BC,
		)
	case 0xc6:
		return NewOp("ADD A,d8", c.ADD, []int{
			8,
		},
			c.A,
			c.D8(),
		)
	case 0xc7:
		return NewOp("RST 00H", c.RST, []int{
			16,
		},
			0x00,
		)
	case 0xc8:
		return NewOp("RET Z", c.RETC, []int{
			20,

			8,
		},
			CaseZ,
		)
	case 0xc9:
		return NewOp("RET", c.RET, []int{
			16,
		},
		)
	case 0xca:
		return NewOp("JP Z,a16", c.JPC, []int{
			16,

			12,
		},
			CaseZ,
			c.A16(),
		)
	case 0xcb:
		return NewOp("PREFIX CB", c.PREFIX, []int{
			4,
		},
			c.CB,
		)
	case 0xcc:
		return NewOp("CALL Z,a16", c.CALLC, []int{
			24,

			12,
		},
			CaseZ,
			c.A16(),
		)
	case 0xcd:
		return NewOp("CALL a16", c.CALL, []int{
			24,
		},
			c.A16(),
		)
	case 0xce:
		return NewOp("ADC A,d8", c.ADC, []int{
			8,
		},
			c.A,
			c.D8(),
		)
	case 0xcf:
		return NewOp("RST 08H", c.RST, []int{
			16,
		},
			0x08,
		)
	case 0xd:
		return NewOp("DEC C", c.DEC, []int{
			4,
		},
			c.C,
		)
	case 0xd0:
		return NewOp("RET NC", c.RETC, []int{
			20,

			8,
		},
			CaseNC,
		)
	case 0xd1:
		return NewOp("POP DE", c.POP, []int{
			12,
		},
			c.DE,
		)
	case 0xd2:
		return NewOp("JP NC,a16", c.JPC, []int{
			16,

			12,
		},
			CaseNC,
			c.A16(),
		)
	case 0xd4:
		return NewOp("CALL NC,a16", c.CALLC, []int{
			24,

			12,
		},
			CaseNC,
			c.A16(),
		)
	case 0xd5:
		return NewOp("PUSH DE", c.PUSH, []int{
			16,
		},
			c.DE,
		)
	case 0xd6:
		return NewOp("SUB d8", c.SUB, []int{
			8,
		},
			c.D8(),
		)
	case 0xd7:
		return NewOp("RST 10H", c.RST, []int{
			16,
		},
			0x10,
		)
	case 0xd8:
		return NewOp("RET C", c.RETC, []int{
			20,

			8,
		},
			CaseC,
		)
	case 0xd9:
		return NewOp("RETI", c.RETI, []int{
			16,
		},
		)
	case 0xda:
		return NewOp("JP C,a16", c.JPC, []int{
			16,

			12,
		},
			CaseC,
			c.A16(),
		)
	case 0xdc:
		return NewOp("CALL C,a16", c.CALLC, []int{
			24,

			12,
		},
			CaseC,
			c.A16(),
		)
	case 0xde:
		return NewOp("SBC A,d8", c.SBC, []int{
			8,
		},
			c.A,
			c.D8(),
		)
	case 0xdf:
		return NewOp("RST 18H", c.RST, []int{
			16,
		},
			0x18,
		)
	case 0xe:
		return NewOp("LD C,d8", c.LD, []int{
			8,
		},
			c.C,
			c.D8(),
		)
	case 0xe0:
		return NewOp("LDH (a8),A", c.LDH, []int{
			12,
		},
			c.MemoryAt(c.A8()),
			c.A,
		)
	case 0xe1:
		return NewOp("POP HL", c.POP, []int{
			12,
		},
			c.HL,
		)
	case 0xe2:
		return NewOp("LD (C),A", c.LD, []int{
			8,
		},
			c.MemoryAt(c.C),
			c.A,
		)
	case 0xe5:
		return NewOp("PUSH HL", c.PUSH, []int{
			16,
		},
			c.HL,
		)
	case 0xe6:
		return NewOp("AND d8", c.AND, []int{
			8,
		},
			c.D8(),
		)
	case 0xe7:
		return NewOp("RST 20H", c.RST, []int{
			16,
		},
			0x20,
		)
	case 0xe8:
		return NewOp("ADD SP,r8", c.ADD, []int{
			16,
		},
			c.SP,
			c.R8(),
		)
	case 0xe9:
		return NewOp("JP (HL)", c.JP, []int{
			4,
		},
			c.MemoryAt(c.HL),
		)
	case 0xea:
		return NewOp("LD (a16),A", c.LD, []int{
			16,
		},
			c.MemoryAt(c.A16()),
			c.A,
		)
	case 0xee:
		return NewOp("XOR d8", c.XOR, []int{
			8,
		},
			c.D8(),
		)
	case 0xef:
		return NewOp("RST 28H", c.RST, []int{
			16,
		},
			0x28,
		)
	case 0xf:
		return NewOp("RRCA", c.RRCA, []int{
			4,
		},
		)
	case 0xf0:
		return NewOp("LDH A,(a8)", c.LDH, []int{
			12,
		},
			c.A,
			c.MemoryAt(c.A8()),
		)
	case 0xf1:
		return NewOp("POP AF", c.POP, []int{
			12,
		},
			c.AF,
		)
	case 0xf2:
		return NewOp("LD A,(C)", c.LD, []int{
			8,
		},
			c.A,
			c.MemoryAt(c.C),
		)
	case 0xf3:
		return NewOp("DI", c.DI, []int{
			4,
		},
		)
	case 0xf5:
		return NewOp("PUSH AF", c.PUSH, []int{
			16,
		},
			c.AF,
		)
	case 0xf6:
		return NewOp("OR d8", c.OR, []int{
			8,
		},
			c.D8(),
		)
	case 0xf7:
		return NewOp("RST 30H", c.RST, []int{
			16,
		},
			0x30,
		)
	case 0xf8:
		return NewOp("LDHL SP,r8", c.LDHL, []int{
			12,
		},
			c.SP,
			c.R8(),
		)
	case 0xf9:
		return NewOp("LD SP,HL", c.LD, []int{
			8,
		},
			c.SP,
			c.HL,
		)
	case 0xfa:
		return NewOp("LD A,(a16)", c.LD, []int{
			16,
		},
			c.A,
			c.MemoryAt(c.A16()),
		)
	case 0xfb:
		return NewOp("EI", c.EI, []int{
			4,
		},
		)
	case 0xfe:
		return NewOp("CP d8", c.CP, []int{
			8,
		},
			c.D8(),
		)
	case 0xff:
		return NewOp("RST 38H", c.RST, []int{
			16,
		},
			0x38,
		)
	default:
		panic(fmt.Sprintf("unknown opcode: 0x%X", code))
	}
}
