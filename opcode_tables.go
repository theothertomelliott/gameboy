package gameboy

import "fmt"

func cbprefixedHandler(c *CPU, code Opcode) (string, []int) {
	switch code {
	case 0x0:
		c.RLC(
			c.B,
		)
		return "RLC B", []int{
			8}
	case 0x1:
		c.RLC(
			c.C,
		)
		return "RLC C", []int{
			8}
	case 0x10:
		c.RL(
			c.B,
		)
		return "RL B", []int{
			8}
	case 0x11:
		c.RL(
			c.C,
		)
		return "RL C", []int{
			8}
	case 0x12:
		c.RL(
			c.D,
		)
		return "RL D", []int{
			8}
	case 0x13:
		c.RL(
			c.E,
		)
		return "RL E", []int{
			8}
	case 0x14:
		c.RL(
			c.H,
		)
		return "RL H", []int{
			8}
	case 0x15:
		c.RL(
			c.L,
		)
		return "RL L", []int{
			8}
	case 0x16:
		c.RL(
			c.MemoryAt(c.HL),
		)
		return "RL (HL)", []int{
			16}
	case 0x17:
		c.RL(
			c.A,
		)
		return "RL A", []int{
			8}
	case 0x18:
		c.RR(
			c.B,
		)
		return "RR B", []int{
			8}
	case 0x19:
		c.RR(
			c.C,
		)
		return "RR C", []int{
			8}
	case 0x1a:
		c.RR(
			c.D,
		)
		return "RR D", []int{
			8}
	case 0x1b:
		c.RR(
			c.E,
		)
		return "RR E", []int{
			8}
	case 0x1c:
		c.RR(
			c.H,
		)
		return "RR H", []int{
			8}
	case 0x1d:
		c.RR(
			c.L,
		)
		return "RR L", []int{
			8}
	case 0x1e:
		c.RR(
			c.MemoryAt(c.HL),
		)
		return "RR (HL)", []int{
			16}
	case 0x1f:
		c.RR(
			c.A,
		)
		return "RR A", []int{
			8}
	case 0x2:
		c.RLC(
			c.D,
		)
		return "RLC D", []int{
			8}
	case 0x20:
		c.SLA(
			c.B,
		)
		return "SLA B", []int{
			8}
	case 0x21:
		c.SLA(
			c.C,
		)
		return "SLA C", []int{
			8}
	case 0x22:
		c.SLA(
			c.D,
		)
		return "SLA D", []int{
			8}
	case 0x23:
		c.SLA(
			c.E,
		)
		return "SLA E", []int{
			8}
	case 0x24:
		c.SLA(
			c.H,
		)
		return "SLA H", []int{
			8}
	case 0x25:
		c.SLA(
			c.L,
		)
		return "SLA L", []int{
			8}
	case 0x26:
		c.SLA(
			c.MemoryAt(c.HL),
		)
		return "SLA (HL)", []int{
			16}
	case 0x27:
		c.SLA(
			c.A,
		)
		return "SLA A", []int{
			8}
	case 0x28:
		c.SRA(
			c.B,
		)
		return "SRA B", []int{
			8}
	case 0x29:
		c.SRA(
			c.C,
		)
		return "SRA C", []int{
			8}
	case 0x2a:
		c.SRA(
			c.D,
		)
		return "SRA D", []int{
			8}
	case 0x2b:
		c.SRA(
			c.E,
		)
		return "SRA E", []int{
			8}
	case 0x2c:
		c.SRA(
			c.H,
		)
		return "SRA H", []int{
			8}
	case 0x2d:
		c.SRA(
			c.L,
		)
		return "SRA L", []int{
			8}
	case 0x2e:
		c.SRA(
			c.MemoryAt(c.HL),
		)
		return "SRA (HL)", []int{
			16}
	case 0x2f:
		c.SRA(
			c.A,
		)
		return "SRA A", []int{
			8}
	case 0x3:
		c.RLC(
			c.E,
		)
		return "RLC E", []int{
			8}
	case 0x30:
		c.SWAP(
			c.B,
		)
		return "SWAP B", []int{
			8}
	case 0x31:
		c.SWAP(
			c.C,
		)
		return "SWAP C", []int{
			8}
	case 0x32:
		c.SWAP(
			c.D,
		)
		return "SWAP D", []int{
			8}
	case 0x33:
		c.SWAP(
			c.E,
		)
		return "SWAP E", []int{
			8}
	case 0x34:
		c.SWAP(
			c.H,
		)
		return "SWAP H", []int{
			8}
	case 0x35:
		c.SWAP(
			c.L,
		)
		return "SWAP L", []int{
			8}
	case 0x36:
		c.SWAP(
			c.MemoryAt(c.HL),
		)
		return "SWAP (HL)", []int{
			16}
	case 0x37:
		c.SWAP(
			c.A,
		)
		return "SWAP A", []int{
			8}
	case 0x38:
		c.SRL(
			c.B,
		)
		return "SRL B", []int{
			8}
	case 0x39:
		c.SRL(
			c.C,
		)
		return "SRL C", []int{
			8}
	case 0x3a:
		c.SRL(
			c.D,
		)
		return "SRL D", []int{
			8}
	case 0x3b:
		c.SRL(
			c.E,
		)
		return "SRL E", []int{
			8}
	case 0x3c:
		c.SRL(
			c.H,
		)
		return "SRL H", []int{
			8}
	case 0x3d:
		c.SRL(
			c.L,
		)
		return "SRL L", []int{
			8}
	case 0x3e:
		c.SRL(
			c.MemoryAt(c.HL),
		)
		return "SRL (HL)", []int{
			16}
	case 0x3f:
		c.SRL(
			c.A,
		)
		return "SRL A", []int{
			8}
	case 0x4:
		c.RLC(
			c.H,
		)
		return "RLC H", []int{
			8}
	case 0x40:
		c.BIT(
			0,
			c.B,
		)
		return "BIT 0,B", []int{
			8}
	case 0x41:
		c.BIT(
			0,
			c.C,
		)
		return "BIT 0,C", []int{
			8}
	case 0x42:
		c.BIT(
			0,
			c.D,
		)
		return "BIT 0,D", []int{
			8}
	case 0x43:
		c.BIT(
			0,
			c.E,
		)
		return "BIT 0,E", []int{
			8}
	case 0x44:
		c.BIT(
			0,
			c.H,
		)
		return "BIT 0,H", []int{
			8}
	case 0x45:
		c.BIT(
			0,
			c.L,
		)
		return "BIT 0,L", []int{
			8}
	case 0x46:
		c.BIT(
			0,
			c.MemoryAt(c.HL),
		)
		return "BIT 0,(HL)", []int{
			16}
	case 0x47:
		c.BIT(
			0,
			c.A,
		)
		return "BIT 0,A", []int{
			8}
	case 0x48:
		c.BIT(
			1,
			c.B,
		)
		return "BIT 1,B", []int{
			8}
	case 0x49:
		c.BIT(
			1,
			c.C,
		)
		return "BIT 1,C", []int{
			8}
	case 0x4a:
		c.BIT(
			1,
			c.D,
		)
		return "BIT 1,D", []int{
			8}
	case 0x4b:
		c.BIT(
			1,
			c.E,
		)
		return "BIT 1,E", []int{
			8}
	case 0x4c:
		c.BIT(
			1,
			c.H,
		)
		return "BIT 1,H", []int{
			8}
	case 0x4d:
		c.BIT(
			1,
			c.L,
		)
		return "BIT 1,L", []int{
			8}
	case 0x4e:
		c.BIT(
			1,
			c.MemoryAt(c.HL),
		)
		return "BIT 1,(HL)", []int{
			16}
	case 0x4f:
		c.BIT(
			1,
			c.A,
		)
		return "BIT 1,A", []int{
			8}
	case 0x5:
		c.RLC(
			c.L,
		)
		return "RLC L", []int{
			8}
	case 0x50:
		c.BIT(
			2,
			c.B,
		)
		return "BIT 2,B", []int{
			8}
	case 0x51:
		c.BIT(
			2,
			c.C,
		)
		return "BIT 2,C", []int{
			8}
	case 0x52:
		c.BIT(
			2,
			c.D,
		)
		return "BIT 2,D", []int{
			8}
	case 0x53:
		c.BIT(
			2,
			c.E,
		)
		return "BIT 2,E", []int{
			8}
	case 0x54:
		c.BIT(
			2,
			c.H,
		)
		return "BIT 2,H", []int{
			8}
	case 0x55:
		c.BIT(
			2,
			c.L,
		)
		return "BIT 2,L", []int{
			8}
	case 0x56:
		c.BIT(
			2,
			c.MemoryAt(c.HL),
		)
		return "BIT 2,(HL)", []int{
			16}
	case 0x57:
		c.BIT(
			2,
			c.A,
		)
		return "BIT 2,A", []int{
			8}
	case 0x58:
		c.BIT(
			3,
			c.B,
		)
		return "BIT 3,B", []int{
			8}
	case 0x59:
		c.BIT(
			3,
			c.C,
		)
		return "BIT 3,C", []int{
			8}
	case 0x5a:
		c.BIT(
			3,
			c.D,
		)
		return "BIT 3,D", []int{
			8}
	case 0x5b:
		c.BIT(
			3,
			c.E,
		)
		return "BIT 3,E", []int{
			8}
	case 0x5c:
		c.BIT(
			3,
			c.H,
		)
		return "BIT 3,H", []int{
			8}
	case 0x5d:
		c.BIT(
			3,
			c.L,
		)
		return "BIT 3,L", []int{
			8}
	case 0x5e:
		c.BIT(
			3,
			c.MemoryAt(c.HL),
		)
		return "BIT 3,(HL)", []int{
			16}
	case 0x5f:
		c.BIT(
			3,
			c.A,
		)
		return "BIT 3,A", []int{
			8}
	case 0x6:
		c.RLC(
			c.MemoryAt(c.HL),
		)
		return "RLC (HL)", []int{
			16}
	case 0x60:
		c.BIT(
			4,
			c.B,
		)
		return "BIT 4,B", []int{
			8}
	case 0x61:
		c.BIT(
			4,
			c.C,
		)
		return "BIT 4,C", []int{
			8}
	case 0x62:
		c.BIT(
			4,
			c.D,
		)
		return "BIT 4,D", []int{
			8}
	case 0x63:
		c.BIT(
			4,
			c.E,
		)
		return "BIT 4,E", []int{
			8}
	case 0x64:
		c.BIT(
			4,
			c.H,
		)
		return "BIT 4,H", []int{
			8}
	case 0x65:
		c.BIT(
			4,
			c.L,
		)
		return "BIT 4,L", []int{
			8}
	case 0x66:
		c.BIT(
			4,
			c.MemoryAt(c.HL),
		)
		return "BIT 4,(HL)", []int{
			16}
	case 0x67:
		c.BIT(
			4,
			c.A,
		)
		return "BIT 4,A", []int{
			8}
	case 0x68:
		c.BIT(
			5,
			c.B,
		)
		return "BIT 5,B", []int{
			8}
	case 0x69:
		c.BIT(
			5,
			c.C,
		)
		return "BIT 5,C", []int{
			8}
	case 0x6a:
		c.BIT(
			5,
			c.D,
		)
		return "BIT 5,D", []int{
			8}
	case 0x6b:
		c.BIT(
			5,
			c.E,
		)
		return "BIT 5,E", []int{
			8}
	case 0x6c:
		c.BIT(
			5,
			c.H,
		)
		return "BIT 5,H", []int{
			8}
	case 0x6d:
		c.BIT(
			5,
			c.L,
		)
		return "BIT 5,L", []int{
			8}
	case 0x6e:
		c.BIT(
			5,
			c.MemoryAt(c.HL),
		)
		return "BIT 5,(HL)", []int{
			16}
	case 0x6f:
		c.BIT(
			5,
			c.A,
		)
		return "BIT 5,A", []int{
			8}
	case 0x7:
		c.RLC(
			c.A,
		)
		return "RLC A", []int{
			8}
	case 0x70:
		c.BIT(
			6,
			c.B,
		)
		return "BIT 6,B", []int{
			8}
	case 0x71:
		c.BIT(
			6,
			c.C,
		)
		return "BIT 6,C", []int{
			8}
	case 0x72:
		c.BIT(
			6,
			c.D,
		)
		return "BIT 6,D", []int{
			8}
	case 0x73:
		c.BIT(
			6,
			c.E,
		)
		return "BIT 6,E", []int{
			8}
	case 0x74:
		c.BIT(
			6,
			c.H,
		)
		return "BIT 6,H", []int{
			8}
	case 0x75:
		c.BIT(
			6,
			c.L,
		)
		return "BIT 6,L", []int{
			8}
	case 0x76:
		c.BIT(
			6,
			c.MemoryAt(c.HL),
		)
		return "BIT 6,(HL)", []int{
			16}
	case 0x77:
		c.BIT(
			6,
			c.A,
		)
		return "BIT 6,A", []int{
			8}
	case 0x78:
		c.BIT(
			7,
			c.B,
		)
		return "BIT 7,B", []int{
			8}
	case 0x79:
		c.BIT(
			7,
			c.C,
		)
		return "BIT 7,C", []int{
			8}
	case 0x7a:
		c.BIT(
			7,
			c.D,
		)
		return "BIT 7,D", []int{
			8}
	case 0x7b:
		c.BIT(
			7,
			c.E,
		)
		return "BIT 7,E", []int{
			8}
	case 0x7c:
		c.BIT(
			7,
			c.H,
		)
		return "BIT 7,H", []int{
			8}
	case 0x7d:
		c.BIT(
			7,
			c.L,
		)
		return "BIT 7,L", []int{
			8}
	case 0x7e:
		c.BIT(
			7,
			c.MemoryAt(c.HL),
		)
		return "BIT 7,(HL)", []int{
			16}
	case 0x7f:
		c.BIT(
			7,
			c.A,
		)
		return "BIT 7,A", []int{
			8}
	case 0x8:
		c.RRC(
			c.B,
		)
		return "RRC B", []int{
			8}
	case 0x80:
		c.RES(
			0,
			c.B,
		)
		return "RES 0,B", []int{
			8}
	case 0x81:
		c.RES(
			0,
			c.C,
		)
		return "RES 0,C", []int{
			8}
	case 0x82:
		c.RES(
			0,
			c.D,
		)
		return "RES 0,D", []int{
			8}
	case 0x83:
		c.RES(
			0,
			c.E,
		)
		return "RES 0,E", []int{
			8}
	case 0x84:
		c.RES(
			0,
			c.H,
		)
		return "RES 0,H", []int{
			8}
	case 0x85:
		c.RES(
			0,
			c.L,
		)
		return "RES 0,L", []int{
			8}
	case 0x86:
		c.RES(
			0,
			c.MemoryAt(c.HL),
		)
		return "RES 0,(HL)", []int{
			16}
	case 0x87:
		c.RES(
			0,
			c.A,
		)
		return "RES 0,A", []int{
			8}
	case 0x88:
		c.RES(
			1,
			c.B,
		)
		return "RES 1,B", []int{
			8}
	case 0x89:
		c.RES(
			1,
			c.C,
		)
		return "RES 1,C", []int{
			8}
	case 0x8a:
		c.RES(
			1,
			c.D,
		)
		return "RES 1,D", []int{
			8}
	case 0x8b:
		c.RES(
			1,
			c.E,
		)
		return "RES 1,E", []int{
			8}
	case 0x8c:
		c.RES(
			1,
			c.H,
		)
		return "RES 1,H", []int{
			8}
	case 0x8d:
		c.RES(
			1,
			c.L,
		)
		return "RES 1,L", []int{
			8}
	case 0x8e:
		c.RES(
			1,
			c.MemoryAt(c.HL),
		)
		return "RES 1,(HL)", []int{
			16}
	case 0x8f:
		c.RES(
			1,
			c.A,
		)
		return "RES 1,A", []int{
			8}
	case 0x9:
		c.RRC(
			c.C,
		)
		return "RRC C", []int{
			8}
	case 0x90:
		c.RES(
			2,
			c.B,
		)
		return "RES 2,B", []int{
			8}
	case 0x91:
		c.RES(
			2,
			c.C,
		)
		return "RES 2,C", []int{
			8}
	case 0x92:
		c.RES(
			2,
			c.D,
		)
		return "RES 2,D", []int{
			8}
	case 0x93:
		c.RES(
			2,
			c.E,
		)
		return "RES 2,E", []int{
			8}
	case 0x94:
		c.RES(
			2,
			c.H,
		)
		return "RES 2,H", []int{
			8}
	case 0x95:
		c.RES(
			2,
			c.L,
		)
		return "RES 2,L", []int{
			8}
	case 0x96:
		c.RES(
			2,
			c.MemoryAt(c.HL),
		)
		return "RES 2,(HL)", []int{
			16}
	case 0x97:
		c.RES(
			2,
			c.A,
		)
		return "RES 2,A", []int{
			8}
	case 0x98:
		c.RES(
			3,
			c.B,
		)
		return "RES 3,B", []int{
			8}
	case 0x99:
		c.RES(
			3,
			c.C,
		)
		return "RES 3,C", []int{
			8}
	case 0x9a:
		c.RES(
			3,
			c.D,
		)
		return "RES 3,D", []int{
			8}
	case 0x9b:
		c.RES(
			3,
			c.E,
		)
		return "RES 3,E", []int{
			8}
	case 0x9c:
		c.RES(
			3,
			c.H,
		)
		return "RES 3,H", []int{
			8}
	case 0x9d:
		c.RES(
			3,
			c.L,
		)
		return "RES 3,L", []int{
			8}
	case 0x9e:
		c.RES(
			3,
			c.MemoryAt(c.HL),
		)
		return "RES 3,(HL)", []int{
			16}
	case 0x9f:
		c.RES(
			3,
			c.A,
		)
		return "RES 3,A", []int{
			8}
	case 0xa:
		c.RRC(
			c.D,
		)
		return "RRC D", []int{
			8}
	case 0xa0:
		c.RES(
			4,
			c.B,
		)
		return "RES 4,B", []int{
			8}
	case 0xa1:
		c.RES(
			4,
			c.C,
		)
		return "RES 4,C", []int{
			8}
	case 0xa2:
		c.RES(
			4,
			c.D,
		)
		return "RES 4,D", []int{
			8}
	case 0xa3:
		c.RES(
			4,
			c.E,
		)
		return "RES 4,E", []int{
			8}
	case 0xa4:
		c.RES(
			4,
			c.H,
		)
		return "RES 4,H", []int{
			8}
	case 0xa5:
		c.RES(
			4,
			c.L,
		)
		return "RES 4,L", []int{
			8}
	case 0xa6:
		c.RES(
			4,
			c.MemoryAt(c.HL),
		)
		return "RES 4,(HL)", []int{
			16}
	case 0xa7:
		c.RES(
			4,
			c.A,
		)
		return "RES 4,A", []int{
			8}
	case 0xa8:
		c.RES(
			5,
			c.B,
		)
		return "RES 5,B", []int{
			8}
	case 0xa9:
		c.RES(
			5,
			c.C,
		)
		return "RES 5,C", []int{
			8}
	case 0xaa:
		c.RES(
			5,
			c.D,
		)
		return "RES 5,D", []int{
			8}
	case 0xab:
		c.RES(
			5,
			c.E,
		)
		return "RES 5,E", []int{
			8}
	case 0xac:
		c.RES(
			5,
			c.H,
		)
		return "RES 5,H", []int{
			8}
	case 0xad:
		c.RES(
			5,
			c.L,
		)
		return "RES 5,L", []int{
			8}
	case 0xae:
		c.RES(
			5,
			c.MemoryAt(c.HL),
		)
		return "RES 5,(HL)", []int{
			16}
	case 0xaf:
		c.RES(
			5,
			c.A,
		)
		return "RES 5,A", []int{
			8}
	case 0xb:
		c.RRC(
			c.E,
		)
		return "RRC E", []int{
			8}
	case 0xb0:
		c.RES(
			6,
			c.B,
		)
		return "RES 6,B", []int{
			8}
	case 0xb1:
		c.RES(
			6,
			c.C,
		)
		return "RES 6,C", []int{
			8}
	case 0xb2:
		c.RES(
			6,
			c.D,
		)
		return "RES 6,D", []int{
			8}
	case 0xb3:
		c.RES(
			6,
			c.E,
		)
		return "RES 6,E", []int{
			8}
	case 0xb4:
		c.RES(
			6,
			c.H,
		)
		return "RES 6,H", []int{
			8}
	case 0xb5:
		c.RES(
			6,
			c.L,
		)
		return "RES 6,L", []int{
			8}
	case 0xb6:
		c.RES(
			6,
			c.MemoryAt(c.HL),
		)
		return "RES 6,(HL)", []int{
			16}
	case 0xb7:
		c.RES(
			6,
			c.A,
		)
		return "RES 6,A", []int{
			8}
	case 0xb8:
		c.RES(
			7,
			c.B,
		)
		return "RES 7,B", []int{
			8}
	case 0xb9:
		c.RES(
			7,
			c.C,
		)
		return "RES 7,C", []int{
			8}
	case 0xba:
		c.RES(
			7,
			c.D,
		)
		return "RES 7,D", []int{
			8}
	case 0xbb:
		c.RES(
			7,
			c.E,
		)
		return "RES 7,E", []int{
			8}
	case 0xbc:
		c.RES(
			7,
			c.H,
		)
		return "RES 7,H", []int{
			8}
	case 0xbd:
		c.RES(
			7,
			c.L,
		)
		return "RES 7,L", []int{
			8}
	case 0xbe:
		c.RES(
			7,
			c.MemoryAt(c.HL),
		)
		return "RES 7,(HL)", []int{
			16}
	case 0xbf:
		c.RES(
			7,
			c.A,
		)
		return "RES 7,A", []int{
			8}
	case 0xc:
		c.RRC(
			c.H,
		)
		return "RRC H", []int{
			8}
	case 0xc0:
		c.SET(
			0,
			c.B,
		)
		return "SET 0,B", []int{
			8}
	case 0xc1:
		c.SET(
			0,
			c.C,
		)
		return "SET 0,C", []int{
			8}
	case 0xc2:
		c.SET(
			0,
			c.D,
		)
		return "SET 0,D", []int{
			8}
	case 0xc3:
		c.SET(
			0,
			c.E,
		)
		return "SET 0,E", []int{
			8}
	case 0xc4:
		c.SET(
			0,
			c.H,
		)
		return "SET 0,H", []int{
			8}
	case 0xc5:
		c.SET(
			0,
			c.L,
		)
		return "SET 0,L", []int{
			8}
	case 0xc6:
		c.SET(
			0,
			c.MemoryAt(c.HL),
		)
		return "SET 0,(HL)", []int{
			16}
	case 0xc7:
		c.SET(
			0,
			c.A,
		)
		return "SET 0,A", []int{
			8}
	case 0xc8:
		c.SET(
			1,
			c.B,
		)
		return "SET 1,B", []int{
			8}
	case 0xc9:
		c.SET(
			1,
			c.C,
		)
		return "SET 1,C", []int{
			8}
	case 0xca:
		c.SET(
			1,
			c.D,
		)
		return "SET 1,D", []int{
			8}
	case 0xcb:
		c.SET(
			1,
			c.E,
		)
		return "SET 1,E", []int{
			8}
	case 0xcc:
		c.SET(
			1,
			c.H,
		)
		return "SET 1,H", []int{
			8}
	case 0xcd:
		c.SET(
			1,
			c.L,
		)
		return "SET 1,L", []int{
			8}
	case 0xce:
		c.SET(
			1,
			c.MemoryAt(c.HL),
		)
		return "SET 1,(HL)", []int{
			16}
	case 0xcf:
		c.SET(
			1,
			c.A,
		)
		return "SET 1,A", []int{
			8}
	case 0xd:
		c.RRC(
			c.L,
		)
		return "RRC L", []int{
			8}
	case 0xd0:
		c.SET(
			2,
			c.B,
		)
		return "SET 2,B", []int{
			8}
	case 0xd1:
		c.SET(
			2,
			c.C,
		)
		return "SET 2,C", []int{
			8}
	case 0xd2:
		c.SET(
			2,
			c.D,
		)
		return "SET 2,D", []int{
			8}
	case 0xd3:
		c.SET(
			2,
			c.E,
		)
		return "SET 2,E", []int{
			8}
	case 0xd4:
		c.SET(
			2,
			c.H,
		)
		return "SET 2,H", []int{
			8}
	case 0xd5:
		c.SET(
			2,
			c.L,
		)
		return "SET 2,L", []int{
			8}
	case 0xd6:
		c.SET(
			2,
			c.MemoryAt(c.HL),
		)
		return "SET 2,(HL)", []int{
			16}
	case 0xd7:
		c.SET(
			2,
			c.A,
		)
		return "SET 2,A", []int{
			8}
	case 0xd8:
		c.SET(
			3,
			c.B,
		)
		return "SET 3,B", []int{
			8}
	case 0xd9:
		c.SET(
			3,
			c.C,
		)
		return "SET 3,C", []int{
			8}
	case 0xda:
		c.SET(
			3,
			c.D,
		)
		return "SET 3,D", []int{
			8}
	case 0xdb:
		c.SET(
			3,
			c.E,
		)
		return "SET 3,E", []int{
			8}
	case 0xdc:
		c.SET(
			3,
			c.H,
		)
		return "SET 3,H", []int{
			8}
	case 0xdd:
		c.SET(
			3,
			c.L,
		)
		return "SET 3,L", []int{
			8}
	case 0xde:
		c.SET(
			3,
			c.MemoryAt(c.HL),
		)
		return "SET 3,(HL)", []int{
			16}
	case 0xdf:
		c.SET(
			3,
			c.A,
		)
		return "SET 3,A", []int{
			8}
	case 0xe:
		c.RRC(
			c.MemoryAt(c.HL),
		)
		return "RRC (HL)", []int{
			16}
	case 0xe0:
		c.SET(
			4,
			c.B,
		)
		return "SET 4,B", []int{
			8}
	case 0xe1:
		c.SET(
			4,
			c.C,
		)
		return "SET 4,C", []int{
			8}
	case 0xe2:
		c.SET(
			4,
			c.D,
		)
		return "SET 4,D", []int{
			8}
	case 0xe3:
		c.SET(
			4,
			c.E,
		)
		return "SET 4,E", []int{
			8}
	case 0xe4:
		c.SET(
			4,
			c.H,
		)
		return "SET 4,H", []int{
			8}
	case 0xe5:
		c.SET(
			4,
			c.L,
		)
		return "SET 4,L", []int{
			8}
	case 0xe6:
		c.SET(
			4,
			c.MemoryAt(c.HL),
		)
		return "SET 4,(HL)", []int{
			16}
	case 0xe7:
		c.SET(
			4,
			c.A,
		)
		return "SET 4,A", []int{
			8}
	case 0xe8:
		c.SET(
			5,
			c.B,
		)
		return "SET 5,B", []int{
			8}
	case 0xe9:
		c.SET(
			5,
			c.C,
		)
		return "SET 5,C", []int{
			8}
	case 0xea:
		c.SET(
			5,
			c.D,
		)
		return "SET 5,D", []int{
			8}
	case 0xeb:
		c.SET(
			5,
			c.E,
		)
		return "SET 5,E", []int{
			8}
	case 0xec:
		c.SET(
			5,
			c.H,
		)
		return "SET 5,H", []int{
			8}
	case 0xed:
		c.SET(
			5,
			c.L,
		)
		return "SET 5,L", []int{
			8}
	case 0xee:
		c.SET(
			5,
			c.MemoryAt(c.HL),
		)
		return "SET 5,(HL)", []int{
			16}
	case 0xef:
		c.SET(
			5,
			c.A,
		)
		return "SET 5,A", []int{
			8}
	case 0xf:
		c.RRC(
			c.A,
		)
		return "RRC A", []int{
			8}
	case 0xf0:
		c.SET(
			6,
			c.B,
		)
		return "SET 6,B", []int{
			8}
	case 0xf1:
		c.SET(
			6,
			c.C,
		)
		return "SET 6,C", []int{
			8}
	case 0xf2:
		c.SET(
			6,
			c.D,
		)
		return "SET 6,D", []int{
			8}
	case 0xf3:
		c.SET(
			6,
			c.E,
		)
		return "SET 6,E", []int{
			8}
	case 0xf4:
		c.SET(
			6,
			c.H,
		)
		return "SET 6,H", []int{
			8}
	case 0xf5:
		c.SET(
			6,
			c.L,
		)
		return "SET 6,L", []int{
			8}
	case 0xf6:
		c.SET(
			6,
			c.MemoryAt(c.HL),
		)
		return "SET 6,(HL)", []int{
			16}
	case 0xf7:
		c.SET(
			6,
			c.A,
		)
		return "SET 6,A", []int{
			8}
	case 0xf8:
		c.SET(
			7,
			c.B,
		)
		return "SET 7,B", []int{
			8}
	case 0xf9:
		c.SET(
			7,
			c.C,
		)
		return "SET 7,C", []int{
			8}
	case 0xfa:
		c.SET(
			7,
			c.D,
		)
		return "SET 7,D", []int{
			8}
	case 0xfb:
		c.SET(
			7,
			c.E,
		)
		return "SET 7,E", []int{
			8}
	case 0xfc:
		c.SET(
			7,
			c.H,
		)
		return "SET 7,H", []int{
			8}
	case 0xfd:
		c.SET(
			7,
			c.L,
		)
		return "SET 7,L", []int{
			8}
	case 0xfe:
		c.SET(
			7,
			c.MemoryAt(c.HL),
		)
		return "SET 7,(HL)", []int{
			16}
	case 0xff:
		c.SET(
			7,
			c.A,
		)
		return "SET 7,A", []int{
			8}
	default:
		panic(fmt.Sprintf("unknown opcode: 0x%X", code))
	}
}
func unprefixedHandler(c *CPU, code Opcode) (string, []int) {
	switch code {
	case 0x0:
		c.NOP()
		return "NOP", []int{
			4}
	case 0x1:
		c.LD(
			c.BC,
			c.D16(),
		)
		return "LD BC,d16", []int{
			12}
	case 0x10:
		c.STOP(
			0,
		)
		return "STOP 0", []int{
			4}
	case 0x11:
		c.LD(
			c.DE,
			c.D16(),
		)
		return "LD DE,d16", []int{
			12}
	case 0x12:
		c.LD(
			c.MemoryAt(c.DE),
			c.A,
		)
		return "LD (DE),A", []int{
			8}
	case 0x13:
		c.INC(
			c.DE,
		)
		return "INC DE", []int{
			8}
	case 0x14:
		c.INC(
			c.D,
		)
		return "INC D", []int{
			4}
	case 0x15:
		c.DEC(
			c.D,
		)
		return "DEC D", []int{
			4}
	case 0x16:
		c.LD(
			c.D,
			c.D8(),
		)
		return "LD D,d8", []int{
			8}
	case 0x17:
		c.RL(
			c.A,
		)
		return "RLA", []int{
			4}
	case 0x18:
		c.JR(
			c.R8(),
		)
		return "JR r8", []int{
			12}
	case 0x19:
		c.ADD(
			c.HL,
			c.DE,
		)
		return "ADD HL,DE", []int{
			8}
	case 0x1a:
		c.LD(
			c.A,
			c.MemoryAt(c.DE),
		)
		return "LD A,(DE)", []int{
			8}
	case 0x1b:
		c.DEC(
			c.DE,
		)
		return "DEC DE", []int{
			8}
	case 0x1c:
		c.INC(
			c.E,
		)
		return "INC E", []int{
			4}
	case 0x1d:
		c.DEC(
			c.E,
		)
		return "DEC E", []int{
			4}
	case 0x1e:
		c.LD(
			c.E,
			c.D8(),
		)
		return "LD E,d8", []int{
			8}
	case 0x1f:
		c.RRA()
		return "RRA", []int{
			4}
	case 0x2:
		c.LD(
			c.MemoryAt(c.BC),
			c.A,
		)
		return "LD (BC),A", []int{
			8}
	case 0x20:
		c.JRC(
			CaseNZ,
			c.R8(),
		)
		return "JR NZ,r8", []int{
			12,
			8}
	case 0x21:
		c.LD(
			c.HL,
			c.D16(),
		)
		return "LD HL,d16", []int{
			12}
	case 0x22:
		c.LDI(
			c.MemoryAt(c.HL),
			c.A,
		)
		return "LDI (HL),A", []int{
			8}
	case 0x23:
		c.INC(
			c.HL,
		)
		return "INC HL", []int{
			8}
	case 0x24:
		c.INC(
			c.H,
		)
		return "INC H", []int{
			4}
	case 0x25:
		c.DEC(
			c.H,
		)
		return "DEC H", []int{
			4}
	case 0x26:
		c.LD(
			c.H,
			c.D8(),
		)
		return "LD H,d8", []int{
			8}
	case 0x27:
		c.DAA()
		return "DAA", []int{
			4}
	case 0x28:
		c.JRC(
			CaseZ,
			c.R8(),
		)
		return "JR Z,r8", []int{
			12,
			8}
	case 0x29:
		c.ADD(
			c.HL,
			c.HL,
		)
		return "ADD HL,HL", []int{
			8}
	case 0x2a:
		c.LDI(
			c.A,
			c.MemoryAt(c.HL),
		)
		return "LDI A,(HL)", []int{
			8}
	case 0x2b:
		c.DEC(
			c.HL,
		)
		return "DEC HL", []int{
			8}
	case 0x2c:
		c.INC(
			c.L,
		)
		return "INC L", []int{
			4}
	case 0x2d:
		c.DEC(
			c.L,
		)
		return "DEC L", []int{
			4}
	case 0x2e:
		c.LD(
			c.L,
			c.D8(),
		)
		return "LD L,d8", []int{
			8}
	case 0x2f:
		c.CPL()
		return "CPL", []int{
			4}
	case 0x3:
		c.INC(
			c.BC,
		)
		return "INC BC", []int{
			8}
	case 0x30:
		c.JRC(
			CaseNC,
			c.R8(),
		)
		return "JR NC,r8", []int{
			12,
			8}
	case 0x31:
		c.LD(
			c.SP,
			c.D16(),
		)
		return "LD SP,d16", []int{
			12}
	case 0x32:
		c.LDD(
			c.MemoryAt(c.HL),
			c.A,
		)
		return "LDD (HL),A", []int{
			8}
	case 0x33:
		c.INC(
			c.SP,
		)
		return "INC SP", []int{
			8}
	case 0x34:
		c.INC(
			c.MemoryAt(c.HL),
		)
		return "INC (HL)", []int{
			12}
	case 0x35:
		c.DEC(
			c.MemoryAt(c.HL),
		)
		return "DEC (HL)", []int{
			12}
	case 0x36:
		c.LD(
			c.MemoryAt(c.HL),
			c.D8(),
		)
		return "LD (HL),d8", []int{
			12}
	case 0x37:
		c.SCF()
		return "SCF", []int{
			4}
	case 0x38:
		c.JRC(
			CaseC,
			c.R8(),
		)
		return "JR C,r8", []int{
			12,
			8}
	case 0x39:
		c.ADD(
			c.HL,
			c.SP,
		)
		return "ADD HL,SP", []int{
			8}
	case 0x3a:
		c.LDD(
			c.A,
			c.MemoryAt(c.HL),
		)
		return "LDD A,(HL)", []int{
			8}
	case 0x3b:
		c.DEC(
			c.SP,
		)
		return "DEC SP", []int{
			8}
	case 0x3c:
		c.INC(
			c.A,
		)
		return "INC A", []int{
			4}
	case 0x3d:
		c.DEC(
			c.A,
		)
		return "DEC A", []int{
			4}
	case 0x3e:
		c.LD(
			c.A,
			c.D8(),
		)
		return "LD A,d8", []int{
			8}
	case 0x3f:
		c.CCF()
		return "CCF", []int{
			4}
	case 0x4:
		c.INC(
			c.B,
		)
		return "INC B", []int{
			4}
	case 0x40:
		c.LD(
			c.B,
			c.B,
		)
		return "LD B,B", []int{
			4}
	case 0x41:
		c.LD(
			c.B,
			c.C,
		)
		return "LD B,C", []int{
			4}
	case 0x42:
		c.LD(
			c.B,
			c.D,
		)
		return "LD B,D", []int{
			4}
	case 0x43:
		c.LD(
			c.B,
			c.E,
		)
		return "LD B,E", []int{
			4}
	case 0x44:
		c.LD(
			c.B,
			c.H,
		)
		return "LD B,H", []int{
			4}
	case 0x45:
		c.LD(
			c.B,
			c.L,
		)
		return "LD B,L", []int{
			4}
	case 0x46:
		c.LD(
			c.B,
			c.MemoryAt(c.HL),
		)
		return "LD B,(HL)", []int{
			8}
	case 0x47:
		c.LD(
			c.B,
			c.A,
		)
		return "LD B,A", []int{
			4}
	case 0x48:
		c.LD(
			c.C,
			c.B,
		)
		return "LD C,B", []int{
			4}
	case 0x49:
		c.LD(
			c.C,
			c.C,
		)
		return "LD C,C", []int{
			4}
	case 0x4a:
		c.LD(
			c.C,
			c.D,
		)
		return "LD C,D", []int{
			4}
	case 0x4b:
		c.LD(
			c.C,
			c.E,
		)
		return "LD C,E", []int{
			4}
	case 0x4c:
		c.LD(
			c.C,
			c.H,
		)
		return "LD C,H", []int{
			4}
	case 0x4d:
		c.LD(
			c.C,
			c.L,
		)
		return "LD C,L", []int{
			4}
	case 0x4e:
		c.LD(
			c.C,
			c.MemoryAt(c.HL),
		)
		return "LD C,(HL)", []int{
			8}
	case 0x4f:
		c.LD(
			c.C,
			c.A,
		)
		return "LD C,A", []int{
			4}
	case 0x5:
		c.DEC(
			c.B,
		)
		return "DEC B", []int{
			4}
	case 0x50:
		c.LD(
			c.D,
			c.B,
		)
		return "LD D,B", []int{
			4}
	case 0x51:
		c.LD(
			c.D,
			c.C,
		)
		return "LD D,C", []int{
			4}
	case 0x52:
		c.LD(
			c.D,
			c.D,
		)
		return "LD D,D", []int{
			4}
	case 0x53:
		c.LD(
			c.D,
			c.E,
		)
		return "LD D,E", []int{
			4}
	case 0x54:
		c.LD(
			c.D,
			c.H,
		)
		return "LD D,H", []int{
			4}
	case 0x55:
		c.LD(
			c.D,
			c.L,
		)
		return "LD D,L", []int{
			4}
	case 0x56:
		c.LD(
			c.D,
			c.MemoryAt(c.HL),
		)
		return "LD D,(HL)", []int{
			8}
	case 0x57:
		c.LD(
			c.D,
			c.A,
		)
		return "LD D,A", []int{
			4}
	case 0x58:
		c.LD(
			c.E,
			c.B,
		)
		return "LD E,B", []int{
			4}
	case 0x59:
		c.LD(
			c.E,
			c.C,
		)
		return "LD E,C", []int{
			4}
	case 0x5a:
		c.LD(
			c.E,
			c.D,
		)
		return "LD E,D", []int{
			4}
	case 0x5b:
		c.LD(
			c.E,
			c.E,
		)
		return "LD E,E", []int{
			4}
	case 0x5c:
		c.LD(
			c.E,
			c.H,
		)
		return "LD E,H", []int{
			4}
	case 0x5d:
		c.LD(
			c.E,
			c.L,
		)
		return "LD E,L", []int{
			4}
	case 0x5e:
		c.LD(
			c.E,
			c.MemoryAt(c.HL),
		)
		return "LD E,(HL)", []int{
			8}
	case 0x5f:
		c.LD(
			c.E,
			c.A,
		)
		return "LD E,A", []int{
			4}
	case 0x6:
		c.LD(
			c.B,
			c.D8(),
		)
		return "LD B,d8", []int{
			8}
	case 0x60:
		c.LD(
			c.H,
			c.B,
		)
		return "LD H,B", []int{
			4}
	case 0x61:
		c.LD(
			c.H,
			c.C,
		)
		return "LD H,C", []int{
			4}
	case 0x62:
		c.LD(
			c.H,
			c.D,
		)
		return "LD H,D", []int{
			4}
	case 0x63:
		c.LD(
			c.H,
			c.E,
		)
		return "LD H,E", []int{
			4}
	case 0x64:
		c.LD(
			c.H,
			c.H,
		)
		return "LD H,H", []int{
			4}
	case 0x65:
		c.LD(
			c.H,
			c.L,
		)
		return "LD H,L", []int{
			4}
	case 0x66:
		c.LD(
			c.H,
			c.MemoryAt(c.HL),
		)
		return "LD H,(HL)", []int{
			8}
	case 0x67:
		c.LD(
			c.H,
			c.A,
		)
		return "LD H,A", []int{
			4}
	case 0x68:
		c.LD(
			c.L,
			c.B,
		)
		return "LD L,B", []int{
			4}
	case 0x69:
		c.LD(
			c.L,
			c.C,
		)
		return "LD L,C", []int{
			4}
	case 0x6a:
		c.LD(
			c.L,
			c.D,
		)
		return "LD L,D", []int{
			4}
	case 0x6b:
		c.LD(
			c.L,
			c.E,
		)
		return "LD L,E", []int{
			4}
	case 0x6c:
		c.LD(
			c.L,
			c.H,
		)
		return "LD L,H", []int{
			4}
	case 0x6d:
		c.LD(
			c.L,
			c.L,
		)
		return "LD L,L", []int{
			4}
	case 0x6e:
		c.LD(
			c.L,
			c.MemoryAt(c.HL),
		)
		return "LD L,(HL)", []int{
			8}
	case 0x6f:
		c.LD(
			c.L,
			c.A,
		)
		return "LD L,A", []int{
			4}
	case 0x7:
		c.RLCA()
		return "RLCA", []int{
			4}
	case 0x70:
		c.LD(
			c.MemoryAt(c.HL),
			c.B,
		)
		return "LD (HL),B", []int{
			8}
	case 0x71:
		c.LD(
			c.MemoryAt(c.HL),
			c.C,
		)
		return "LD (HL),C", []int{
			8}
	case 0x72:
		c.LD(
			c.MemoryAt(c.HL),
			c.D,
		)
		return "LD (HL),D", []int{
			8}
	case 0x73:
		c.LD(
			c.MemoryAt(c.HL),
			c.E,
		)
		return "LD (HL),E", []int{
			8}
	case 0x74:
		c.LD(
			c.MemoryAt(c.HL),
			c.H,
		)
		return "LD (HL),H", []int{
			8}
	case 0x75:
		c.LD(
			c.MemoryAt(c.HL),
			c.L,
		)
		return "LD (HL),L", []int{
			8}
	case 0x76:
		c.HALT()
		return "HALT", []int{
			4}
	case 0x77:
		c.LD(
			c.MemoryAt(c.HL),
			c.A,
		)
		return "LD (HL),A", []int{
			8}
	case 0x78:
		c.LD(
			c.A,
			c.B,
		)
		return "LD A,B", []int{
			4}
	case 0x79:
		c.LD(
			c.A,
			c.C,
		)
		return "LD A,C", []int{
			4}
	case 0x7a:
		c.LD(
			c.A,
			c.D,
		)
		return "LD A,D", []int{
			4}
	case 0x7b:
		c.LD(
			c.A,
			c.E,
		)
		return "LD A,E", []int{
			4}
	case 0x7c:
		c.LD(
			c.A,
			c.H,
		)
		return "LD A,H", []int{
			4}
	case 0x7d:
		c.LD(
			c.A,
			c.L,
		)
		return "LD A,L", []int{
			4}
	case 0x7e:
		c.LD(
			c.A,
			c.MemoryAt(c.HL),
		)
		return "LD A,(HL)", []int{
			8}
	case 0x7f:
		c.LD(
			c.A,
			c.A,
		)
		return "LD A,A", []int{
			4}
	case 0x8:
		c.LD(
			c.MemoryAt(c.A16()),
			c.SP,
		)
		return "LD (a16),SP", []int{
			20}
	case 0x80:
		c.ADD(
			c.A,
			c.B,
		)
		return "ADD A,B", []int{
			4}
	case 0x81:
		c.ADD(
			c.A,
			c.C,
		)
		return "ADD A,C", []int{
			4}
	case 0x82:
		c.ADD(
			c.A,
			c.D,
		)
		return "ADD A,D", []int{
			4}
	case 0x83:
		c.ADD(
			c.A,
			c.E,
		)
		return "ADD A,E", []int{
			4}
	case 0x84:
		c.ADD(
			c.A,
			c.H,
		)
		return "ADD A,H", []int{
			4}
	case 0x85:
		c.ADD(
			c.A,
			c.L,
		)
		return "ADD A,L", []int{
			4}
	case 0x86:
		c.ADD(
			c.A,
			c.MemoryAt(c.HL),
		)
		return "ADD A,(HL)", []int{
			8}
	case 0x87:
		c.ADD(
			c.A,
			c.A,
		)
		return "ADD A,A", []int{
			4}
	case 0x88:
		c.ADC(
			c.A,
			c.B,
		)
		return "ADC A,B", []int{
			4}
	case 0x89:
		c.ADC(
			c.A,
			c.C,
		)
		return "ADC A,C", []int{
			4}
	case 0x8a:
		c.ADC(
			c.A,
			c.D,
		)
		return "ADC A,D", []int{
			4}
	case 0x8b:
		c.ADC(
			c.A,
			c.E,
		)
		return "ADC A,E", []int{
			4}
	case 0x8c:
		c.ADC(
			c.A,
			c.H,
		)
		return "ADC A,H", []int{
			4}
	case 0x8d:
		c.ADC(
			c.A,
			c.L,
		)
		return "ADC A,L", []int{
			4}
	case 0x8e:
		c.ADC(
			c.A,
			c.MemoryAt(c.HL),
		)
		return "ADC A,(HL)", []int{
			8}
	case 0x8f:
		c.ADC(
			c.A,
			c.A,
		)
		return "ADC A,A", []int{
			4}
	case 0x9:
		c.ADD(
			c.HL,
			c.BC,
		)
		return "ADD HL,BC", []int{
			8}
	case 0x90:
		c.SUB(
			c.B,
		)
		return "SUB B", []int{
			4}
	case 0x91:
		c.SUB(
			c.C,
		)
		return "SUB C", []int{
			4}
	case 0x92:
		c.SUB(
			c.D,
		)
		return "SUB D", []int{
			4}
	case 0x93:
		c.SUB(
			c.E,
		)
		return "SUB E", []int{
			4}
	case 0x94:
		c.SUB(
			c.H,
		)
		return "SUB H", []int{
			4}
	case 0x95:
		c.SUB(
			c.L,
		)
		return "SUB L", []int{
			4}
	case 0x96:
		c.SUB(
			c.MemoryAt(c.HL),
		)
		return "SUB (HL)", []int{
			8}
	case 0x97:
		c.SUB(
			c.A,
		)
		return "SUB A", []int{
			4}
	case 0x98:
		c.SBC(
			c.A,
			c.B,
		)
		return "SBC A,B", []int{
			4}
	case 0x99:
		c.SBC(
			c.A,
			c.C,
		)
		return "SBC A,C", []int{
			4}
	case 0x9a:
		c.SBC(
			c.A,
			c.D,
		)
		return "SBC A,D", []int{
			4}
	case 0x9b:
		c.SBC(
			c.A,
			c.E,
		)
		return "SBC A,E", []int{
			4}
	case 0x9c:
		c.SBC(
			c.A,
			c.H,
		)
		return "SBC A,H", []int{
			4}
	case 0x9d:
		c.SBC(
			c.A,
			c.L,
		)
		return "SBC A,L", []int{
			4}
	case 0x9e:
		c.SBC(
			c.A,
			c.MemoryAt(c.HL),
		)
		return "SBC A,(HL)", []int{
			8}
	case 0x9f:
		c.SBC(
			c.A,
			c.A,
		)
		return "SBC A,A", []int{
			4}
	case 0xa:
		c.LD(
			c.A,
			c.MemoryAt(c.BC),
		)
		return "LD A,(BC)", []int{
			8}
	case 0xa0:
		c.AND(
			c.B,
		)
		return "AND B", []int{
			4}
	case 0xa1:
		c.AND(
			c.C,
		)
		return "AND C", []int{
			4}
	case 0xa2:
		c.AND(
			c.D,
		)
		return "AND D", []int{
			4}
	case 0xa3:
		c.AND(
			c.E,
		)
		return "AND E", []int{
			4}
	case 0xa4:
		c.AND(
			c.H,
		)
		return "AND H", []int{
			4}
	case 0xa5:
		c.AND(
			c.L,
		)
		return "AND L", []int{
			4}
	case 0xa6:
		c.AND(
			c.MemoryAt(c.HL),
		)
		return "AND (HL)", []int{
			8}
	case 0xa7:
		c.AND(
			c.A,
		)
		return "AND A", []int{
			4}
	case 0xa8:
		c.XOR(
			c.B,
		)
		return "XOR B", []int{
			4}
	case 0xa9:
		c.XOR(
			c.C,
		)
		return "XOR C", []int{
			4}
	case 0xaa:
		c.XOR(
			c.D,
		)
		return "XOR D", []int{
			4}
	case 0xab:
		c.XOR(
			c.E,
		)
		return "XOR E", []int{
			4}
	case 0xac:
		c.XOR(
			c.H,
		)
		return "XOR H", []int{
			4}
	case 0xad:
		c.XOR(
			c.L,
		)
		return "XOR L", []int{
			4}
	case 0xae:
		c.XOR(
			c.MemoryAt(c.HL),
		)
		return "XOR (HL)", []int{
			8}
	case 0xaf:
		c.XOR(
			c.A,
		)
		return "XOR A", []int{
			4}
	case 0xb:
		c.DEC(
			c.BC,
		)
		return "DEC BC", []int{
			8}
	case 0xb0:
		c.OR(
			c.B,
		)
		return "OR B", []int{
			4}
	case 0xb1:
		c.OR(
			c.C,
		)
		return "OR C", []int{
			4}
	case 0xb2:
		c.OR(
			c.D,
		)
		return "OR D", []int{
			4}
	case 0xb3:
		c.OR(
			c.E,
		)
		return "OR E", []int{
			4}
	case 0xb4:
		c.OR(
			c.H,
		)
		return "OR H", []int{
			4}
	case 0xb5:
		c.OR(
			c.L,
		)
		return "OR L", []int{
			4}
	case 0xb6:
		c.OR(
			c.MemoryAt(c.HL),
		)
		return "OR (HL)", []int{
			8}
	case 0xb7:
		c.OR(
			c.A,
		)
		return "OR A", []int{
			4}
	case 0xb8:
		c.CP(
			c.B,
		)
		return "CP B", []int{
			4}
	case 0xb9:
		c.CP(
			c.C,
		)
		return "CP C", []int{
			4}
	case 0xba:
		c.CP(
			c.D,
		)
		return "CP D", []int{
			4}
	case 0xbb:
		c.CP(
			c.E,
		)
		return "CP E", []int{
			4}
	case 0xbc:
		c.CP(
			c.H,
		)
		return "CP H", []int{
			4}
	case 0xbd:
		c.CP(
			c.L,
		)
		return "CP L", []int{
			4}
	case 0xbe:
		c.CP(
			c.MemoryAt(c.HL),
		)
		return "CP (HL)", []int{
			8}
	case 0xbf:
		c.CP(
			c.A,
		)
		return "CP A", []int{
			4}
	case 0xc:
		c.INC(
			c.C,
		)
		return "INC C", []int{
			4}
	case 0xc0:
		c.RETC(
			CaseNZ,
		)
		return "RET NZ", []int{
			20,
			8}
	case 0xc1:
		c.POP(
			c.BC,
		)
		return "POP BC", []int{
			12}
	case 0xc2:
		c.JPC(
			CaseNZ,
			c.A16(),
		)
		return "JP NZ,a16", []int{
			16,
			12}
	case 0xc3:
		c.JP(
			c.A16(),
		)
		return "JP a16", []int{
			16}
	case 0xc4:
		c.CALLC(
			CaseNZ,
			c.A16(),
		)
		return "CALL NZ,a16", []int{
			24,
			12}
	case 0xc5:
		c.PUSH(
			c.BC,
		)
		return "PUSH BC", []int{
			16}
	case 0xc6:
		c.ADD(
			c.A,
			c.D8(),
		)
		return "ADD A,d8", []int{
			8}
	case 0xc7:
		c.RST(
			0x00,
		)
		return "RST 00H", []int{
			16}
	case 0xc8:
		c.RETC(
			CaseZ,
		)
		return "RET Z", []int{
			20,
			8}
	case 0xc9:
		c.RET()
		return "RET", []int{
			16}
	case 0xca:
		c.JPC(
			CaseZ,
			c.A16(),
		)
		return "JP Z,a16", []int{
			16,
			12}
	case 0xcb:
		c.PREFIX(
			c.CB,
		)
		return "PREFIX CB", []int{
			4}
	case 0xcc:
		c.CALLC(
			CaseZ,
			c.A16(),
		)
		return "CALL Z,a16", []int{
			24,
			12}
	case 0xcd:
		c.CALL(
			c.A16(),
		)
		return "CALL a16", []int{
			24}
	case 0xce:
		c.ADC(
			c.A,
			c.D8(),
		)
		return "ADC A,d8", []int{
			8}
	case 0xcf:
		c.RST(
			0x08,
		)
		return "RST 08H", []int{
			16}
	case 0xd:
		c.DEC(
			c.C,
		)
		return "DEC C", []int{
			4}
	case 0xd0:
		c.RETC(
			CaseNC,
		)
		return "RET NC", []int{
			20,
			8}
	case 0xd1:
		c.POP(
			c.DE,
		)
		return "POP DE", []int{
			12}
	case 0xd2:
		c.JPC(
			CaseNC,
			c.A16(),
		)
		return "JP NC,a16", []int{
			16,
			12}
	case 0xd4:
		c.CALLC(
			CaseNC,
			c.A16(),
		)
		return "CALL NC,a16", []int{
			24,
			12}
	case 0xd5:
		c.PUSH(
			c.DE,
		)
		return "PUSH DE", []int{
			16}
	case 0xd6:
		c.SUB(
			c.D8(),
		)
		return "SUB d8", []int{
			8}
	case 0xd7:
		c.RST(
			0x10,
		)
		return "RST 10H", []int{
			16}
	case 0xd8:
		c.RETC(
			CaseC,
		)
		return "RET C", []int{
			20,
			8}
	case 0xd9:
		c.RETI()
		return "RETI", []int{
			16}
	case 0xda:
		c.JPC(
			CaseC,
			c.A16(),
		)
		return "JP C,a16", []int{
			16,
			12}
	case 0xdc:
		c.CALLC(
			CaseC,
			c.A16(),
		)
		return "CALL C,a16", []int{
			24,
			12}
	case 0xde:
		c.SBC(
			c.A,
			c.D8(),
		)
		return "SBC A,d8", []int{
			8}
	case 0xdf:
		c.RST(
			0x18,
		)
		return "RST 18H", []int{
			16}
	case 0xe:
		c.LD(
			c.C,
			c.D8(),
		)
		return "LD C,d8", []int{
			8}
	case 0xe0:
		c.LDH(
			c.MemoryAt(c.A8()),
			c.A,
		)
		return "LDH (a8),A", []int{
			12}
	case 0xe1:
		c.POP(
			c.HL,
		)
		return "POP HL", []int{
			12}
	case 0xe2:
		c.LD(
			c.MemoryAt(c.C),
			c.A,
		)
		return "LD (C),A", []int{
			8}
	case 0xe5:
		c.PUSH(
			c.HL,
		)
		return "PUSH HL", []int{
			16}
	case 0xe6:
		c.AND(
			c.D8(),
		)
		return "AND d8", []int{
			8}
	case 0xe7:
		c.RST(
			0x20,
		)
		return "RST 20H", []int{
			16}
	case 0xe8:
		c.ADD(
			c.SP,
			c.R8(),
		)
		return "ADD SP,r8", []int{
			16}
	case 0xe9:
		c.JP(
			c.MemoryAt(c.HL),
		)
		return "JP (HL)", []int{
			4}
	case 0xea:
		c.LD(
			c.MemoryAt(c.A16()),
			c.A,
		)
		return "LD (a16),A", []int{
			16}
	case 0xee:
		c.XOR(
			c.D8(),
		)
		return "XOR d8", []int{
			8}
	case 0xef:
		c.RST(
			0x28,
		)
		return "RST 28H", []int{
			16}
	case 0xf:
		c.RRCA()
		return "RRCA", []int{
			4}
	case 0xf0:
		c.LDH(
			c.A,
			c.MemoryAt(c.A8()),
		)
		return "LDH A,(a8)", []int{
			12}
	case 0xf1:
		c.POP(
			c.AF,
		)
		return "POP AF", []int{
			12}
	case 0xf2:
		c.LD(
			c.A,
			c.MemoryAt(c.C),
		)
		return "LD A,(C)", []int{
			8}
	case 0xf3:
		c.DI()
		return "DI", []int{
			4}
	case 0xf5:
		c.PUSH(
			c.AF,
		)
		return "PUSH AF", []int{
			16}
	case 0xf6:
		c.OR(
			c.D8(),
		)
		return "OR d8", []int{
			8}
	case 0xf7:
		c.RST(
			0x30,
		)
		return "RST 30H", []int{
			16}
	case 0xf8:
		c.LDHL(
			c.SP,
			c.R8(),
		)
		return "LDHL SP,r8", []int{
			12}
	case 0xf9:
		c.LD(
			c.SP,
			c.HL,
		)
		return "LD SP,HL", []int{
			8}
	case 0xfa:
		c.LD(
			c.A,
			c.MemoryAt(c.A16()),
		)
		return "LD A,(a16)", []int{
			16}
	case 0xfb:
		c.EI()
		return "EI", []int{
			4}
	case 0xfe:
		c.CP(
			c.D8(),
		)
		return "CP d8", []int{
			8}
	case 0xff:
		c.RST(
			0x38,
		)
		return "RST 38H", []int{
			16}
	default:
		panic(fmt.Sprintf("unknown opcode: 0x%X", code))
	}
}
