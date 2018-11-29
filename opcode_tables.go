package gameboy

import "fmt"

func cbprefixedHandler(c *CPU, code Opcode) (string, []int, error) {
	switch code {
	case 0x0:
		o1 := c.B
		c.RLC(
			o1,
		)
		description := fmt.Sprint(
			"RLC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x1:
		o1 := c.C
		c.RLC(
			o1,
		)
		description := fmt.Sprint(
			"RLC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x10:
		o1 := c.B
		c.RL(
			o1,
		)
		description := fmt.Sprint(
			"RL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x11:
		o1 := c.C
		c.RL(
			o1,
		)
		description := fmt.Sprint(
			"RL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x12:
		o1 := c.D
		c.RL(
			o1,
		)
		description := fmt.Sprint(
			"RL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x13:
		o1 := c.E
		c.RL(
			o1,
		)
		description := fmt.Sprint(
			"RL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x14:
		o1 := c.H
		c.RL(
			o1,
		)
		description := fmt.Sprint(
			"RL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x15:
		o1 := c.L
		c.RL(
			o1,
		)
		description := fmt.Sprint(
			"RL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x16:
		o1 := c.MemoryAt(c.HL)
		c.RL(
			o1,
		)
		description := fmt.Sprint(
			"RL ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0x17:
		o1 := c.A
		c.RL(
			o1,
		)
		description := fmt.Sprint(
			"RL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x18:
		o1 := c.B
		c.RR(
			o1,
		)
		description := fmt.Sprint(
			"RR ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x19:
		o1 := c.C
		c.RR(
			o1,
		)
		description := fmt.Sprint(
			"RR ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x1a:
		o1 := c.D
		c.RR(
			o1,
		)
		description := fmt.Sprint(
			"RR ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x1b:
		o1 := c.E
		c.RR(
			o1,
		)
		description := fmt.Sprint(
			"RR ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x1c:
		o1 := c.H
		c.RR(
			o1,
		)
		description := fmt.Sprint(
			"RR ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x1d:
		o1 := c.L
		c.RR(
			o1,
		)
		description := fmt.Sprint(
			"RR ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x1e:
		o1 := c.MemoryAt(c.HL)
		c.RR(
			o1,
		)
		description := fmt.Sprint(
			"RR ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0x1f:
		o1 := c.A
		c.RR(
			o1,
		)
		description := fmt.Sprint(
			"RR ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x2:
		o1 := c.D
		c.RLC(
			o1,
		)
		description := fmt.Sprint(
			"RLC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x20:
		o1 := c.B
		c.SLA(
			o1,
		)
		description := fmt.Sprint(
			"SLA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x21:
		o1 := c.C
		c.SLA(
			o1,
		)
		description := fmt.Sprint(
			"SLA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x22:
		o1 := c.D
		c.SLA(
			o1,
		)
		description := fmt.Sprint(
			"SLA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x23:
		o1 := c.E
		c.SLA(
			o1,
		)
		description := fmt.Sprint(
			"SLA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x24:
		o1 := c.H
		c.SLA(
			o1,
		)
		description := fmt.Sprint(
			"SLA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x25:
		o1 := c.L
		c.SLA(
			o1,
		)
		description := fmt.Sprint(
			"SLA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x26:
		o1 := c.MemoryAt(c.HL)
		c.SLA(
			o1,
		)
		description := fmt.Sprint(
			"SLA ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0x27:
		o1 := c.A
		c.SLA(
			o1,
		)
		description := fmt.Sprint(
			"SLA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x28:
		o1 := c.B
		c.SRA(
			o1,
		)
		description := fmt.Sprint(
			"SRA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x29:
		o1 := c.C
		c.SRA(
			o1,
		)
		description := fmt.Sprint(
			"SRA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x2a:
		o1 := c.D
		c.SRA(
			o1,
		)
		description := fmt.Sprint(
			"SRA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x2b:
		o1 := c.E
		c.SRA(
			o1,
		)
		description := fmt.Sprint(
			"SRA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x2c:
		o1 := c.H
		c.SRA(
			o1,
		)
		description := fmt.Sprint(
			"SRA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x2d:
		o1 := c.L
		c.SRA(
			o1,
		)
		description := fmt.Sprint(
			"SRA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x2e:
		o1 := c.MemoryAt(c.HL)
		c.SRA(
			o1,
		)
		description := fmt.Sprint(
			"SRA ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0x2f:
		o1 := c.A
		c.SRA(
			o1,
		)
		description := fmt.Sprint(
			"SRA ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x3:
		o1 := c.E
		c.RLC(
			o1,
		)
		description := fmt.Sprint(
			"RLC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x30:
		o1 := c.B
		c.SWAP(
			o1,
		)
		description := fmt.Sprint(
			"SWAP ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x31:
		o1 := c.C
		c.SWAP(
			o1,
		)
		description := fmt.Sprint(
			"SWAP ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x32:
		o1 := c.D
		c.SWAP(
			o1,
		)
		description := fmt.Sprint(
			"SWAP ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x33:
		o1 := c.E
		c.SWAP(
			o1,
		)
		description := fmt.Sprint(
			"SWAP ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x34:
		o1 := c.H
		c.SWAP(
			o1,
		)
		description := fmt.Sprint(
			"SWAP ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x35:
		o1 := c.L
		c.SWAP(
			o1,
		)
		description := fmt.Sprint(
			"SWAP ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x36:
		o1 := c.MemoryAt(c.HL)
		c.SWAP(
			o1,
		)
		description := fmt.Sprint(
			"SWAP ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0x37:
		o1 := c.A
		c.SWAP(
			o1,
		)
		description := fmt.Sprint(
			"SWAP ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x38:
		o1 := c.B
		c.SRL(
			o1,
		)
		description := fmt.Sprint(
			"SRL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x39:
		o1 := c.C
		c.SRL(
			o1,
		)
		description := fmt.Sprint(
			"SRL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x3a:
		o1 := c.D
		c.SRL(
			o1,
		)
		description := fmt.Sprint(
			"SRL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x3b:
		o1 := c.E
		c.SRL(
			o1,
		)
		description := fmt.Sprint(
			"SRL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x3c:
		o1 := c.H
		c.SRL(
			o1,
		)
		description := fmt.Sprint(
			"SRL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x3d:
		o1 := c.L
		c.SRL(
			o1,
		)
		description := fmt.Sprint(
			"SRL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x3e:
		o1 := c.MemoryAt(c.HL)
		c.SRL(
			o1,
		)
		description := fmt.Sprint(
			"SRL ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0x3f:
		o1 := c.A
		c.SRL(
			o1,
		)
		description := fmt.Sprint(
			"SRL ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x4:
		o1 := c.H
		c.RLC(
			o1,
		)
		description := fmt.Sprint(
			"RLC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x40:
		o1 := 0
		o2 := c.B
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x41:
		o1 := 0
		o2 := c.C
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x42:
		o1 := 0
		o2 := c.D
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x43:
		o1 := 0
		o2 := c.E
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x44:
		o1 := 0
		o2 := c.H
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x45:
		o1 := 0
		o2 := c.L
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x46:
		o1 := 0
		o2 := c.MemoryAt(c.HL)
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0x47:
		o1 := 0
		o2 := c.A
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x48:
		o1 := 1
		o2 := c.B
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x49:
		o1 := 1
		o2 := c.C
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x4a:
		o1 := 1
		o2 := c.D
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x4b:
		o1 := 1
		o2 := c.E
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x4c:
		o1 := 1
		o2 := c.H
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x4d:
		o1 := 1
		o2 := c.L
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x4e:
		o1 := 1
		o2 := c.MemoryAt(c.HL)
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0x4f:
		o1 := 1
		o2 := c.A
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x5:
		o1 := c.L
		c.RLC(
			o1,
		)
		description := fmt.Sprint(
			"RLC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x50:
		o1 := 2
		o2 := c.B
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x51:
		o1 := 2
		o2 := c.C
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x52:
		o1 := 2
		o2 := c.D
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x53:
		o1 := 2
		o2 := c.E
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x54:
		o1 := 2
		o2 := c.H
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x55:
		o1 := 2
		o2 := c.L
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x56:
		o1 := 2
		o2 := c.MemoryAt(c.HL)
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0x57:
		o1 := 2
		o2 := c.A
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x58:
		o1 := 3
		o2 := c.B
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x59:
		o1 := 3
		o2 := c.C
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x5a:
		o1 := 3
		o2 := c.D
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x5b:
		o1 := 3
		o2 := c.E
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x5c:
		o1 := 3
		o2 := c.H
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x5d:
		o1 := 3
		o2 := c.L
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x5e:
		o1 := 3
		o2 := c.MemoryAt(c.HL)
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0x5f:
		o1 := 3
		o2 := c.A
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x6:
		o1 := c.MemoryAt(c.HL)
		c.RLC(
			o1,
		)
		description := fmt.Sprint(
			"RLC ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0x60:
		o1 := 4
		o2 := c.B
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x61:
		o1 := 4
		o2 := c.C
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x62:
		o1 := 4
		o2 := c.D
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x63:
		o1 := 4
		o2 := c.E
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x64:
		o1 := 4
		o2 := c.H
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x65:
		o1 := 4
		o2 := c.L
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x66:
		o1 := 4
		o2 := c.MemoryAt(c.HL)
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0x67:
		o1 := 4
		o2 := c.A
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x68:
		o1 := 5
		o2 := c.B
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x69:
		o1 := 5
		o2 := c.C
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x6a:
		o1 := 5
		o2 := c.D
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x6b:
		o1 := 5
		o2 := c.E
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x6c:
		o1 := 5
		o2 := c.H
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x6d:
		o1 := 5
		o2 := c.L
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x6e:
		o1 := 5
		o2 := c.MemoryAt(c.HL)
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0x6f:
		o1 := 5
		o2 := c.A
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x7:
		o1 := c.A
		c.RLC(
			o1,
		)
		description := fmt.Sprint(
			"RLC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x70:
		o1 := 6
		o2 := c.B
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x71:
		o1 := 6
		o2 := c.C
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x72:
		o1 := 6
		o2 := c.D
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x73:
		o1 := 6
		o2 := c.E
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x74:
		o1 := 6
		o2 := c.H
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x75:
		o1 := 6
		o2 := c.L
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x76:
		o1 := 6
		o2 := c.MemoryAt(c.HL)
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0x77:
		o1 := 6
		o2 := c.A
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x78:
		o1 := 7
		o2 := c.B
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x79:
		o1 := 7
		o2 := c.C
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x7a:
		o1 := 7
		o2 := c.D
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x7b:
		o1 := 7
		o2 := c.E
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x7c:
		o1 := 7
		o2 := c.H
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x7d:
		o1 := 7
		o2 := c.L
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x7e:
		o1 := 7
		o2 := c.MemoryAt(c.HL)
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0x7f:
		o1 := 7
		o2 := c.A
		c.BIT(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"BIT ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x8:
		o1 := c.B
		c.RRC(
			o1,
		)
		description := fmt.Sprint(
			"RRC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x80:
		o1 := 0
		o2 := c.B
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x81:
		o1 := 0
		o2 := c.C
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x82:
		o1 := 0
		o2 := c.D
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x83:
		o1 := 0
		o2 := c.E
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x84:
		o1 := 0
		o2 := c.H
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x85:
		o1 := 0
		o2 := c.L
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x86:
		o1 := 0
		o2 := c.MemoryAt(c.HL)
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0x87:
		o1 := 0
		o2 := c.A
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x88:
		o1 := 1
		o2 := c.B
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x89:
		o1 := 1
		o2 := c.C
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x8a:
		o1 := 1
		o2 := c.D
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x8b:
		o1 := 1
		o2 := c.E
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x8c:
		o1 := 1
		o2 := c.H
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x8d:
		o1 := 1
		o2 := c.L
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x8e:
		o1 := 1
		o2 := c.MemoryAt(c.HL)
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0x8f:
		o1 := 1
		o2 := c.A
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x9:
		o1 := c.C
		c.RRC(
			o1,
		)
		description := fmt.Sprint(
			"RRC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x90:
		o1 := 2
		o2 := c.B
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x91:
		o1 := 2
		o2 := c.C
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x92:
		o1 := 2
		o2 := c.D
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x93:
		o1 := 2
		o2 := c.E
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x94:
		o1 := 2
		o2 := c.H
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x95:
		o1 := 2
		o2 := c.L
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x96:
		o1 := 2
		o2 := c.MemoryAt(c.HL)
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0x97:
		o1 := 2
		o2 := c.A
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x98:
		o1 := 3
		o2 := c.B
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x99:
		o1 := 3
		o2 := c.C
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x9a:
		o1 := 3
		o2 := c.D
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x9b:
		o1 := 3
		o2 := c.E
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x9c:
		o1 := 3
		o2 := c.H
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x9d:
		o1 := 3
		o2 := c.L
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x9e:
		o1 := 3
		o2 := c.MemoryAt(c.HL)
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0x9f:
		o1 := 3
		o2 := c.A
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xa:
		o1 := c.D
		c.RRC(
			o1,
		)
		description := fmt.Sprint(
			"RRC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xa0:
		o1 := 4
		o2 := c.B
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xa1:
		o1 := 4
		o2 := c.C
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xa2:
		o1 := 4
		o2 := c.D
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xa3:
		o1 := 4
		o2 := c.E
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xa4:
		o1 := 4
		o2 := c.H
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xa5:
		o1 := 4
		o2 := c.L
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xa6:
		o1 := 4
		o2 := c.MemoryAt(c.HL)
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xa7:
		o1 := 4
		o2 := c.A
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xa8:
		o1 := 5
		o2 := c.B
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xa9:
		o1 := 5
		o2 := c.C
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xaa:
		o1 := 5
		o2 := c.D
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xab:
		o1 := 5
		o2 := c.E
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xac:
		o1 := 5
		o2 := c.H
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xad:
		o1 := 5
		o2 := c.L
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xae:
		o1 := 5
		o2 := c.MemoryAt(c.HL)
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xaf:
		o1 := 5
		o2 := c.A
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xb:
		o1 := c.E
		c.RRC(
			o1,
		)
		description := fmt.Sprint(
			"RRC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xb0:
		o1 := 6
		o2 := c.B
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xb1:
		o1 := 6
		o2 := c.C
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xb2:
		o1 := 6
		o2 := c.D
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xb3:
		o1 := 6
		o2 := c.E
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xb4:
		o1 := 6
		o2 := c.H
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xb5:
		o1 := 6
		o2 := c.L
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xb6:
		o1 := 6
		o2 := c.MemoryAt(c.HL)
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xb7:
		o1 := 6
		o2 := c.A
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xb8:
		o1 := 7
		o2 := c.B
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xb9:
		o1 := 7
		o2 := c.C
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xba:
		o1 := 7
		o2 := c.D
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xbb:
		o1 := 7
		o2 := c.E
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xbc:
		o1 := 7
		o2 := c.H
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xbd:
		o1 := 7
		o2 := c.L
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xbe:
		o1 := 7
		o2 := c.MemoryAt(c.HL)
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xbf:
		o1 := 7
		o2 := c.A
		c.RES(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"RES ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xc:
		o1 := c.H
		c.RRC(
			o1,
		)
		description := fmt.Sprint(
			"RRC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xc0:
		o1 := 0
		o2 := c.B
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xc1:
		o1 := 0
		o2 := c.C
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xc2:
		o1 := 0
		o2 := c.D
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xc3:
		o1 := 0
		o2 := c.E
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xc4:
		o1 := 0
		o2 := c.H
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xc5:
		o1 := 0
		o2 := c.L
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xc6:
		o1 := 0
		o2 := c.MemoryAt(c.HL)
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xc7:
		o1 := 0
		o2 := c.A
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xc8:
		o1 := 1
		o2 := c.B
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xc9:
		o1 := 1
		o2 := c.C
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xca:
		o1 := 1
		o2 := c.D
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xcb:
		o1 := 1
		o2 := c.E
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xcc:
		o1 := 1
		o2 := c.H
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xcd:
		o1 := 1
		o2 := c.L
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xce:
		o1 := 1
		o2 := c.MemoryAt(c.HL)
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xcf:
		o1 := 1
		o2 := c.A
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xd:
		o1 := c.L
		c.RRC(
			o1,
		)
		description := fmt.Sprint(
			"RRC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xd0:
		o1 := 2
		o2 := c.B
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xd1:
		o1 := 2
		o2 := c.C
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xd2:
		o1 := 2
		o2 := c.D
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xd3:
		o1 := 2
		o2 := c.E
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xd4:
		o1 := 2
		o2 := c.H
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xd5:
		o1 := 2
		o2 := c.L
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xd6:
		o1 := 2
		o2 := c.MemoryAt(c.HL)
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xd7:
		o1 := 2
		o2 := c.A
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xd8:
		o1 := 3
		o2 := c.B
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xd9:
		o1 := 3
		o2 := c.C
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xda:
		o1 := 3
		o2 := c.D
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xdb:
		o1 := 3
		o2 := c.E
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xdc:
		o1 := 3
		o2 := c.H
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xdd:
		o1 := 3
		o2 := c.L
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xde:
		o1 := 3
		o2 := c.MemoryAt(c.HL)
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xdf:
		o1 := 3
		o2 := c.A
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xe:
		o1 := c.MemoryAt(c.HL)
		c.RRC(
			o1,
		)
		description := fmt.Sprint(
			"RRC ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xe0:
		o1 := 4
		o2 := c.B
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xe1:
		o1 := 4
		o2 := c.C
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xe2:
		o1 := 4
		o2 := c.D
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xe3:
		o1 := 4
		o2 := c.E
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xe4:
		o1 := 4
		o2 := c.H
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xe5:
		o1 := 4
		o2 := c.L
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xe6:
		o1 := 4
		o2 := c.MemoryAt(c.HL)
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xe7:
		o1 := 4
		o2 := c.A
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xe8:
		o1 := 5
		o2 := c.B
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xe9:
		o1 := 5
		o2 := c.C
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xea:
		o1 := 5
		o2 := c.D
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xeb:
		o1 := 5
		o2 := c.E
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xec:
		o1 := 5
		o2 := c.H
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xed:
		o1 := 5
		o2 := c.L
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xee:
		o1 := 5
		o2 := c.MemoryAt(c.HL)
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xef:
		o1 := 5
		o2 := c.A
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xf:
		o1 := c.A
		c.RRC(
			o1,
		)
		description := fmt.Sprint(
			"RRC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xf0:
		o1 := 6
		o2 := c.B
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xf1:
		o1 := 6
		o2 := c.C
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xf2:
		o1 := 6
		o2 := c.D
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xf3:
		o1 := 6
		o2 := c.E
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xf4:
		o1 := 6
		o2 := c.H
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xf5:
		o1 := 6
		o2 := c.L
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xf6:
		o1 := 6
		o2 := c.MemoryAt(c.HL)
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xf7:
		o1 := 6
		o2 := c.A
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xf8:
		o1 := 7
		o2 := c.B
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xf9:
		o1 := 7
		o2 := c.C
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xfa:
		o1 := 7
		o2 := c.D
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xfb:
		o1 := 7
		o2 := c.E
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xfc:
		o1 := 7
		o2 := c.H
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xfd:
		o1 := 7
		o2 := c.L
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xfe:
		o1 := 7
		o2 := c.MemoryAt(c.HL)
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xff:
		o1 := 7
		o2 := c.A
		c.SET(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SET ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	default:
		return "", nil, fmt.Errorf("unknown opcode: 0x%X", code)
	}
}
func unprefixedHandler(c *CPU, code Opcode) (string, []int, error) {
	switch code {
	case 0x0:
		c.NOP()
		description := fmt.Sprint(
			"NOP ",
		)
		return description, []int{
			4}, nil
	case 0x1:
		o1 := c.BC
		o2 := c.D16()
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			12}, nil
	case 0x10:
		o1 := 0
		c.STOP(
			o1,
		)
		description := fmt.Sprint(
			"STOP ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x11:
		o1 := c.DE
		o2 := c.D16()
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			12}, nil
	case 0x12:
		o1 := c.MemoryAt(c.DE)
		o2 := c.A
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x13:
		o1 := c.DE
		c.INC(
			o1,
		)
		description := fmt.Sprint(
			"INC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x14:
		o1 := c.D
		c.INC(
			o1,
		)
		description := fmt.Sprint(
			"INC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x15:
		o1 := c.D
		c.DEC(
			o1,
		)
		description := fmt.Sprint(
			"DEC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x16:
		o1 := c.D
		o2 := c.D8()
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x17:
		o1 := c.A
		c.RL(
			o1,
		)
		description := fmt.Sprint(
			"RLA ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x18:
		o1 := c.R8()
		c.JR(
			o1,
		)
		description := fmt.Sprint(
			"JR ",
			o1,
		)
		return description, []int{
			12}, nil
	case 0x19:
		o1 := c.HL
		o2 := c.DE
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x1a:
		o1 := c.A
		o2 := c.MemoryAt(c.DE)
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x1b:
		o1 := c.DE
		c.DEC(
			o1,
		)
		description := fmt.Sprint(
			"DEC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x1c:
		o1 := c.E
		c.INC(
			o1,
		)
		description := fmt.Sprint(
			"INC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x1d:
		o1 := c.E
		c.DEC(
			o1,
		)
		description := fmt.Sprint(
			"DEC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x1e:
		o1 := c.E
		o2 := c.D8()
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x1f:
		c.RRA()
		description := fmt.Sprint(
			"RRA ",
		)
		return description, []int{
			4}, nil
	case 0x2:
		o1 := c.MemoryAt(c.BC)
		o2 := c.A
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x20:
		o1 := CaseNZ
		o2 := c.R8()
		c.JRC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"JR ",
			o1,
			o2,
		)
		return description, []int{
			12,
			8}, nil
	case 0x21:
		o1 := c.HL
		o2 := c.D16()
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			12}, nil
	case 0x22:
		o1 := c.MemoryAt(c.HL)
		o2 := c.A
		c.LDI(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LDI ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x23:
		o1 := c.HL
		c.INC(
			o1,
		)
		description := fmt.Sprint(
			"INC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x24:
		o1 := c.H
		c.INC(
			o1,
		)
		description := fmt.Sprint(
			"INC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x25:
		o1 := c.H
		c.DEC(
			o1,
		)
		description := fmt.Sprint(
			"DEC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x26:
		o1 := c.H
		o2 := c.D8()
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x27:
		c.DAA()
		description := fmt.Sprint(
			"DAA ",
		)
		return description, []int{
			4}, nil
	case 0x28:
		o1 := CaseZ
		o2 := c.R8()
		c.JRC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"JR ",
			o1,
			o2,
		)
		return description, []int{
			12,
			8}, nil
	case 0x29:
		o1 := c.HL
		o2 := c.HL
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x2a:
		o1 := c.A
		o2 := c.MemoryAt(c.HL)
		c.LDI(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LDI ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x2b:
		o1 := c.HL
		c.DEC(
			o1,
		)
		description := fmt.Sprint(
			"DEC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x2c:
		o1 := c.L
		c.INC(
			o1,
		)
		description := fmt.Sprint(
			"INC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x2d:
		o1 := c.L
		c.DEC(
			o1,
		)
		description := fmt.Sprint(
			"DEC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x2e:
		o1 := c.L
		o2 := c.D8()
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x2f:
		c.CPL()
		description := fmt.Sprint(
			"CPL ",
		)
		return description, []int{
			4}, nil
	case 0x3:
		o1 := c.BC
		c.INC(
			o1,
		)
		description := fmt.Sprint(
			"INC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x30:
		o1 := CaseNC
		o2 := c.R8()
		c.JRC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"JR ",
			o1,
			o2,
		)
		return description, []int{
			12,
			8}, nil
	case 0x31:
		o1 := c.SP
		o2 := c.D16()
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			12}, nil
	case 0x32:
		o1 := c.MemoryAt(c.HL)
		o2 := c.A
		c.LDD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LDD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x33:
		o1 := c.SP
		c.INC(
			o1,
		)
		description := fmt.Sprint(
			"INC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x34:
		o1 := c.MemoryAt(c.HL)
		c.INC(
			o1,
		)
		description := fmt.Sprint(
			"INC ",
			o1,
		)
		return description, []int{
			12}, nil
	case 0x35:
		o1 := c.MemoryAt(c.HL)
		c.DEC(
			o1,
		)
		description := fmt.Sprint(
			"DEC ",
			o1,
		)
		return description, []int{
			12}, nil
	case 0x36:
		o1 := c.MemoryAt(c.HL)
		o2 := c.D8()
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			12}, nil
	case 0x37:
		c.SCF()
		description := fmt.Sprint(
			"SCF ",
		)
		return description, []int{
			4}, nil
	case 0x38:
		o1 := CaseC
		o2 := c.R8()
		c.JRC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"JR ",
			o1,
			o2,
		)
		return description, []int{
			12,
			8}, nil
	case 0x39:
		o1 := c.HL
		o2 := c.SP
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x3a:
		o1 := c.A
		o2 := c.MemoryAt(c.HL)
		c.LDD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LDD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x3b:
		o1 := c.SP
		c.DEC(
			o1,
		)
		description := fmt.Sprint(
			"DEC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x3c:
		o1 := c.A
		c.INC(
			o1,
		)
		description := fmt.Sprint(
			"INC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x3d:
		o1 := c.A
		c.DEC(
			o1,
		)
		description := fmt.Sprint(
			"DEC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x3e:
		o1 := c.A
		o2 := c.D8()
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x3f:
		c.CCF()
		description := fmt.Sprint(
			"CCF ",
		)
		return description, []int{
			4}, nil
	case 0x4:
		o1 := c.B
		c.INC(
			o1,
		)
		description := fmt.Sprint(
			"INC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x40:
		o1 := c.B
		o2 := c.B
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x41:
		o1 := c.B
		o2 := c.C
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x42:
		o1 := c.B
		o2 := c.D
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x43:
		o1 := c.B
		o2 := c.E
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x44:
		o1 := c.B
		o2 := c.H
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x45:
		o1 := c.B
		o2 := c.L
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x46:
		o1 := c.B
		o2 := c.MemoryAt(c.HL)
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x47:
		o1 := c.B
		o2 := c.A
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x48:
		o1 := c.C
		o2 := c.B
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x49:
		o1 := c.C
		o2 := c.C
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x4a:
		o1 := c.C
		o2 := c.D
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x4b:
		o1 := c.C
		o2 := c.E
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x4c:
		o1 := c.C
		o2 := c.H
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x4d:
		o1 := c.C
		o2 := c.L
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x4e:
		o1 := c.C
		o2 := c.MemoryAt(c.HL)
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x4f:
		o1 := c.C
		o2 := c.A
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x5:
		o1 := c.B
		c.DEC(
			o1,
		)
		description := fmt.Sprint(
			"DEC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x50:
		o1 := c.D
		o2 := c.B
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x51:
		o1 := c.D
		o2 := c.C
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x52:
		o1 := c.D
		o2 := c.D
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x53:
		o1 := c.D
		o2 := c.E
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x54:
		o1 := c.D
		o2 := c.H
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x55:
		o1 := c.D
		o2 := c.L
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x56:
		o1 := c.D
		o2 := c.MemoryAt(c.HL)
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x57:
		o1 := c.D
		o2 := c.A
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x58:
		o1 := c.E
		o2 := c.B
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x59:
		o1 := c.E
		o2 := c.C
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x5a:
		o1 := c.E
		o2 := c.D
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x5b:
		o1 := c.E
		o2 := c.E
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x5c:
		o1 := c.E
		o2 := c.H
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x5d:
		o1 := c.E
		o2 := c.L
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x5e:
		o1 := c.E
		o2 := c.MemoryAt(c.HL)
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x5f:
		o1 := c.E
		o2 := c.A
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x6:
		o1 := c.B
		o2 := c.D8()
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x60:
		o1 := c.H
		o2 := c.B
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x61:
		o1 := c.H
		o2 := c.C
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x62:
		o1 := c.H
		o2 := c.D
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x63:
		o1 := c.H
		o2 := c.E
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x64:
		o1 := c.H
		o2 := c.H
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x65:
		o1 := c.H
		o2 := c.L
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x66:
		o1 := c.H
		o2 := c.MemoryAt(c.HL)
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x67:
		o1 := c.H
		o2 := c.A
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x68:
		o1 := c.L
		o2 := c.B
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x69:
		o1 := c.L
		o2 := c.C
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x6a:
		o1 := c.L
		o2 := c.D
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x6b:
		o1 := c.L
		o2 := c.E
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x6c:
		o1 := c.L
		o2 := c.H
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x6d:
		o1 := c.L
		o2 := c.L
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x6e:
		o1 := c.L
		o2 := c.MemoryAt(c.HL)
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x6f:
		o1 := c.L
		o2 := c.A
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x7:
		c.RLCA()
		description := fmt.Sprint(
			"RLCA ",
		)
		return description, []int{
			4}, nil
	case 0x70:
		o1 := c.MemoryAt(c.HL)
		o2 := c.B
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x71:
		o1 := c.MemoryAt(c.HL)
		o2 := c.C
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x72:
		o1 := c.MemoryAt(c.HL)
		o2 := c.D
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x73:
		o1 := c.MemoryAt(c.HL)
		o2 := c.E
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x74:
		o1 := c.MemoryAt(c.HL)
		o2 := c.H
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x75:
		o1 := c.MemoryAt(c.HL)
		o2 := c.L
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x76:
		c.HALT()
		description := fmt.Sprint(
			"HALT ",
		)
		return description, []int{
			4}, nil
	case 0x77:
		o1 := c.MemoryAt(c.HL)
		o2 := c.A
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x78:
		o1 := c.A
		o2 := c.B
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x79:
		o1 := c.A
		o2 := c.C
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x7a:
		o1 := c.A
		o2 := c.D
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x7b:
		o1 := c.A
		o2 := c.E
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x7c:
		o1 := c.A
		o2 := c.H
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x7d:
		o1 := c.A
		o2 := c.L
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x7e:
		o1 := c.A
		o2 := c.MemoryAt(c.HL)
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x7f:
		o1 := c.A
		o2 := c.A
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x8:
		o1 := c.MemoryAt(c.A16())
		o2 := c.SP
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			20}, nil
	case 0x80:
		o1 := c.A
		o2 := c.B
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x81:
		o1 := c.A
		o2 := c.C
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x82:
		o1 := c.A
		o2 := c.D
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x83:
		o1 := c.A
		o2 := c.E
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x84:
		o1 := c.A
		o2 := c.H
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x85:
		o1 := c.A
		o2 := c.L
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x86:
		o1 := c.A
		o2 := c.MemoryAt(c.HL)
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x87:
		o1 := c.A
		o2 := c.A
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x88:
		o1 := c.A
		o2 := c.B
		c.ADC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x89:
		o1 := c.A
		o2 := c.C
		c.ADC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x8a:
		o1 := c.A
		o2 := c.D
		c.ADC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x8b:
		o1 := c.A
		o2 := c.E
		c.ADC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x8c:
		o1 := c.A
		o2 := c.H
		c.ADC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x8d:
		o1 := c.A
		o2 := c.L
		c.ADC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x8e:
		o1 := c.A
		o2 := c.MemoryAt(c.HL)
		c.ADC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADC ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x8f:
		o1 := c.A
		o2 := c.A
		c.ADC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x9:
		o1 := c.HL
		o2 := c.BC
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x90:
		o1 := c.B
		c.SUB(
			o1,
		)
		description := fmt.Sprint(
			"SUB ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x91:
		o1 := c.C
		c.SUB(
			o1,
		)
		description := fmt.Sprint(
			"SUB ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x92:
		o1 := c.D
		c.SUB(
			o1,
		)
		description := fmt.Sprint(
			"SUB ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x93:
		o1 := c.E
		c.SUB(
			o1,
		)
		description := fmt.Sprint(
			"SUB ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x94:
		o1 := c.H
		c.SUB(
			o1,
		)
		description := fmt.Sprint(
			"SUB ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x95:
		o1 := c.L
		c.SUB(
			o1,
		)
		description := fmt.Sprint(
			"SUB ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x96:
		o1 := c.MemoryAt(c.HL)
		c.SUB(
			o1,
		)
		description := fmt.Sprint(
			"SUB ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0x97:
		o1 := c.A
		c.SUB(
			o1,
		)
		description := fmt.Sprint(
			"SUB ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0x98:
		o1 := c.A
		o2 := c.B
		c.SBC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SBC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x99:
		o1 := c.A
		o2 := c.C
		c.SBC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SBC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x9a:
		o1 := c.A
		o2 := c.D
		c.SBC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SBC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x9b:
		o1 := c.A
		o2 := c.E
		c.SBC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SBC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x9c:
		o1 := c.A
		o2 := c.H
		c.SBC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SBC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x9d:
		o1 := c.A
		o2 := c.L
		c.SBC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SBC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0x9e:
		o1 := c.A
		o2 := c.MemoryAt(c.HL)
		c.SBC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SBC ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0x9f:
		o1 := c.A
		o2 := c.A
		c.SBC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SBC ",
			o1,
			o2,
		)
		return description, []int{
			4}, nil
	case 0xa:
		o1 := c.A
		o2 := c.MemoryAt(c.BC)
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xa0:
		o1 := c.B
		c.AND(
			o1,
		)
		description := fmt.Sprint(
			"AND ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xa1:
		o1 := c.C
		c.AND(
			o1,
		)
		description := fmt.Sprint(
			"AND ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xa2:
		o1 := c.D
		c.AND(
			o1,
		)
		description := fmt.Sprint(
			"AND ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xa3:
		o1 := c.E
		c.AND(
			o1,
		)
		description := fmt.Sprint(
			"AND ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xa4:
		o1 := c.H
		c.AND(
			o1,
		)
		description := fmt.Sprint(
			"AND ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xa5:
		o1 := c.L
		c.AND(
			o1,
		)
		description := fmt.Sprint(
			"AND ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xa6:
		o1 := c.MemoryAt(c.HL)
		c.AND(
			o1,
		)
		description := fmt.Sprint(
			"AND ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xa7:
		o1 := c.A
		c.AND(
			o1,
		)
		description := fmt.Sprint(
			"AND ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xa8:
		o1 := c.B
		c.XOR(
			o1,
		)
		description := fmt.Sprint(
			"XOR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xa9:
		o1 := c.C
		c.XOR(
			o1,
		)
		description := fmt.Sprint(
			"XOR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xaa:
		o1 := c.D
		c.XOR(
			o1,
		)
		description := fmt.Sprint(
			"XOR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xab:
		o1 := c.E
		c.XOR(
			o1,
		)
		description := fmt.Sprint(
			"XOR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xac:
		o1 := c.H
		c.XOR(
			o1,
		)
		description := fmt.Sprint(
			"XOR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xad:
		o1 := c.L
		c.XOR(
			o1,
		)
		description := fmt.Sprint(
			"XOR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xae:
		o1 := c.MemoryAt(c.HL)
		c.XOR(
			o1,
		)
		description := fmt.Sprint(
			"XOR ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xaf:
		o1 := c.A
		c.XOR(
			o1,
		)
		description := fmt.Sprint(
			"XOR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xb:
		o1 := c.BC
		c.DEC(
			o1,
		)
		description := fmt.Sprint(
			"DEC ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xb0:
		o1 := c.B
		c.OR(
			o1,
		)
		description := fmt.Sprint(
			"OR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xb1:
		o1 := c.C
		c.OR(
			o1,
		)
		description := fmt.Sprint(
			"OR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xb2:
		o1 := c.D
		c.OR(
			o1,
		)
		description := fmt.Sprint(
			"OR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xb3:
		o1 := c.E
		c.OR(
			o1,
		)
		description := fmt.Sprint(
			"OR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xb4:
		o1 := c.H
		c.OR(
			o1,
		)
		description := fmt.Sprint(
			"OR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xb5:
		o1 := c.L
		c.OR(
			o1,
		)
		description := fmt.Sprint(
			"OR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xb6:
		o1 := c.MemoryAt(c.HL)
		c.OR(
			o1,
		)
		description := fmt.Sprint(
			"OR ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xb7:
		o1 := c.A
		c.OR(
			o1,
		)
		description := fmt.Sprint(
			"OR ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xb8:
		o1 := c.B
		c.CP(
			o1,
		)
		description := fmt.Sprint(
			"CP ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xb9:
		o1 := c.C
		c.CP(
			o1,
		)
		description := fmt.Sprint(
			"CP ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xba:
		o1 := c.D
		c.CP(
			o1,
		)
		description := fmt.Sprint(
			"CP ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xbb:
		o1 := c.E
		c.CP(
			o1,
		)
		description := fmt.Sprint(
			"CP ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xbc:
		o1 := c.H
		c.CP(
			o1,
		)
		description := fmt.Sprint(
			"CP ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xbd:
		o1 := c.L
		c.CP(
			o1,
		)
		description := fmt.Sprint(
			"CP ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xbe:
		o1 := c.MemoryAt(c.HL)
		c.CP(
			o1,
		)
		description := fmt.Sprint(
			"CP ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xbf:
		o1 := c.A
		c.CP(
			o1,
		)
		description := fmt.Sprint(
			"CP ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xc:
		o1 := c.C
		c.INC(
			o1,
		)
		description := fmt.Sprint(
			"INC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xc0:
		o1 := CaseNZ
		c.RETC(
			o1,
		)
		description := fmt.Sprint(
			"RET ",
			o1,
		)
		return description, []int{
			20,
			8}, nil
	case 0xc1:
		o1 := c.BC
		c.POP(
			o1,
		)
		description := fmt.Sprint(
			"POP ",
			o1,
		)
		return description, []int{
			12}, nil
	case 0xc2:
		o1 := CaseNZ
		o2 := c.A16()
		c.JPC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"JP ",
			o1,
			o2,
		)
		return description, []int{
			16,
			12}, nil
	case 0xc3:
		o1 := c.A16()
		c.JP(
			o1,
		)
		description := fmt.Sprint(
			"JP ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xc4:
		o1 := CaseNZ
		o2 := c.A16()
		c.CALLC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"CALL ",
			o1,
			o2,
		)
		return description, []int{
			24,
			12}, nil
	case 0xc5:
		o1 := c.BC
		c.PUSH(
			o1,
		)
		description := fmt.Sprint(
			"PUSH ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xc6:
		o1 := c.A
		o2 := c.D8()
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xc7:
		o1 := 0x00
		c.RST(
			o1,
		)
		description := fmt.Sprint(
			"RST ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xc8:
		o1 := CaseZ
		c.RETC(
			o1,
		)
		description := fmt.Sprint(
			"RET ",
			o1,
		)
		return description, []int{
			20,
			8}, nil
	case 0xc9:
		c.RET()
		description := fmt.Sprint(
			"RET ",
		)
		return description, []int{
			16}, nil
	case 0xca:
		o1 := CaseZ
		o2 := c.A16()
		c.JPC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"JP ",
			o1,
			o2,
		)
		return description, []int{
			16,
			12}, nil
	case 0xcb:
		o1 := c.CB
		c.PREFIX(
			o1,
		)
		description := fmt.Sprint(
			"PREFIX ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xcc:
		o1 := CaseZ
		o2 := c.A16()
		c.CALLC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"CALL ",
			o1,
			o2,
		)
		return description, []int{
			24,
			12}, nil
	case 0xcd:
		o1 := c.A16()
		c.CALL(
			o1,
		)
		description := fmt.Sprint(
			"CALL ",
			o1,
		)
		return description, []int{
			24}, nil
	case 0xce:
		o1 := c.A
		o2 := c.D8()
		c.ADC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADC ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xcf:
		o1 := 0x08
		c.RST(
			o1,
		)
		description := fmt.Sprint(
			"RST ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xd:
		o1 := c.C
		c.DEC(
			o1,
		)
		description := fmt.Sprint(
			"DEC ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xd0:
		o1 := CaseNC
		c.RETC(
			o1,
		)
		description := fmt.Sprint(
			"RET ",
			o1,
		)
		return description, []int{
			20,
			8}, nil
	case 0xd1:
		o1 := c.DE
		c.POP(
			o1,
		)
		description := fmt.Sprint(
			"POP ",
			o1,
		)
		return description, []int{
			12}, nil
	case 0xd2:
		o1 := CaseNC
		o2 := c.A16()
		c.JPC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"JP ",
			o1,
			o2,
		)
		return description, []int{
			16,
			12}, nil
	case 0xd4:
		o1 := CaseNC
		o2 := c.A16()
		c.CALLC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"CALL ",
			o1,
			o2,
		)
		return description, []int{
			24,
			12}, nil
	case 0xd5:
		o1 := c.DE
		c.PUSH(
			o1,
		)
		description := fmt.Sprint(
			"PUSH ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xd6:
		o1 := c.D8()
		c.SUB(
			o1,
		)
		description := fmt.Sprint(
			"SUB ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xd7:
		o1 := 0x10
		c.RST(
			o1,
		)
		description := fmt.Sprint(
			"RST ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xd8:
		o1 := CaseC
		c.RETC(
			o1,
		)
		description := fmt.Sprint(
			"RET ",
			o1,
		)
		return description, []int{
			20,
			8}, nil
	case 0xd9:
		c.RETI()
		description := fmt.Sprint(
			"RETI ",
		)
		return description, []int{
			16}, nil
	case 0xda:
		o1 := CaseC
		o2 := c.A16()
		c.JPC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"JP ",
			o1,
			o2,
		)
		return description, []int{
			16,
			12}, nil
	case 0xdc:
		o1 := CaseC
		o2 := c.A16()
		c.CALLC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"CALL ",
			o1,
			o2,
		)
		return description, []int{
			24,
			12}, nil
	case 0xde:
		o1 := c.A
		o2 := c.D8()
		c.SBC(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"SBC ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xdf:
		o1 := 0x18
		c.RST(
			o1,
		)
		description := fmt.Sprint(
			"RST ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xe:
		o1 := c.C
		o2 := c.D8()
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xe0:
		o1 := c.MemoryAt(c.A8())
		o2 := c.A
		c.LDH(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LDH ",
			o1,
			o2,
		)
		return description, []int{
			12}, nil
	case 0xe1:
		o1 := c.HL
		c.POP(
			o1,
		)
		description := fmt.Sprint(
			"POP ",
			o1,
		)
		return description, []int{
			12}, nil
	case 0xe2:
		o1 := c.MemoryAtH(c.C)
		o2 := c.A
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xe5:
		o1 := c.HL
		c.PUSH(
			o1,
		)
		description := fmt.Sprint(
			"PUSH ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xe6:
		o1 := c.D8()
		c.AND(
			o1,
		)
		description := fmt.Sprint(
			"AND ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xe7:
		o1 := 0x20
		c.RST(
			o1,
		)
		description := fmt.Sprint(
			"RST ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xe8:
		o1 := c.SP
		o2 := c.R8()
		c.ADD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"ADD ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xe9:
		o1 := c.HL
		c.JP(
			o1,
		)
		description := fmt.Sprint(
			"JP ",
			o1,
		)
		return description, []int{
			4}, nil
	case 0xea:
		o1 := c.MemoryAt(c.A16())
		o2 := c.A
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xee:
		o1 := c.D8()
		c.XOR(
			o1,
		)
		description := fmt.Sprint(
			"XOR ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xef:
		o1 := 0x28
		c.RST(
			o1,
		)
		description := fmt.Sprint(
			"RST ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xf:
		c.RRCA()
		description := fmt.Sprint(
			"RRCA ",
		)
		return description, []int{
			4}, nil
	case 0xf0:
		o1 := c.A
		o2 := c.MemoryAt(c.A8())
		c.LDH(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LDH ",
			o1,
			o2,
		)
		return description, []int{
			12}, nil
	case 0xf1:
		o1 := c.AF
		c.POP(
			o1,
		)
		description := fmt.Sprint(
			"POP ",
			o1,
		)
		return description, []int{
			12}, nil
	case 0xf2:
		o1 := c.A
		o2 := c.MemoryAtH(c.C)
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xf3:
		c.DI()
		description := fmt.Sprint(
			"DI ",
		)
		return description, []int{
			4}, nil
	case 0xf5:
		o1 := c.AF
		c.PUSH(
			o1,
		)
		description := fmt.Sprint(
			"PUSH ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xf6:
		o1 := c.D8()
		c.OR(
			o1,
		)
		description := fmt.Sprint(
			"OR ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xf7:
		o1 := 0x30
		c.RST(
			o1,
		)
		description := fmt.Sprint(
			"RST ",
			o1,
		)
		return description, []int{
			16}, nil
	case 0xf8:
		o1 := c.SP
		o2 := c.R8()
		c.LDHL(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LDHL ",
			o1,
			o2,
		)
		return description, []int{
			12}, nil
	case 0xf9:
		o1 := c.SP
		o2 := c.HL
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			8}, nil
	case 0xfa:
		o1 := c.A
		o2 := c.MemoryAt(c.A16())
		c.LD(
			o1,
			o2,
		)
		description := fmt.Sprint(
			"LD ",
			o1,
			o2,
		)
		return description, []int{
			16}, nil
	case 0xfb:
		c.EI()
		description := fmt.Sprint(
			"EI ",
		)
		return description, []int{
			4}, nil
	case 0xfe:
		o1 := c.D8()
		c.CP(
			o1,
		)
		description := fmt.Sprint(
			"CP ",
			o1,
		)
		return description, []int{
			8}, nil
	case 0xff:
		o1 := 0x38
		c.RST(
			o1,
		)
		description := fmt.Sprint(
			"RST ",
			o1,
		)
		return description, []int{
			16}, nil
	default:
		return "", nil, fmt.Errorf("unknown opcode: 0x%X", code)
	}
}
