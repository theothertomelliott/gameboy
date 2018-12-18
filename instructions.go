package gameboy

// Instructions for the Gameboy
// Based on: http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf

// NOP is a no-op instruction
func (c *CPU) NOP(...Param) {}

// LD loads src into dst
//
// Use with:
//  LD nn,n:
//   nn (src) = B,C,D,E,H,L,BC,DE,HL,SP
//   n (dst = 8 bit immediate value
//  LD r1,r2:
//   r1, r2 = A,B,C,D,E,H,L,(HL)
//  LD A,n:
//   n (src) = A,B,C,D,E,H,L,(BC),(DE),(HL),(nn),#
//   nn (src) = two byte immediate value. (LS byte first.)
//  LD n,A:
//   n = A,B,C,D,E,H,L,(BC),(DE),(HL),(nn)
//   nn = two byte immediate value. (LS byte first.)
//
// VARIANTS:
//  LD A,(C):
//   Put value at address $FF00 + register C into A.
//   Same as: LD A,($FF00+C)
//  LD (C),A:
//   Put A into address $FF00 + register C.
//  LD A,(HLD):
//   Same as: LDD A,(HL)
//  LD A,(HL-):
//   Same as: LDD A,(HL)
//  LDD A,(HL):
//   Put value at address HL into A. Decrement HL.
//   Same as: LD A,(HL) - DEC HL
//  LD (HLD),A:
//   Same as: LDD (HL),A
//  LD (HL-),A
//   Same as: LDD (HL),A
//  LDD (HL),A
//   Put A into memory address HL. Decrement HL.
//   Same as: LD (HL),A - DEC HL
//  LD A,(HLI)
//   Same as: LDI A,(HL)
//  LD A,(HL+)
//   Same as: LDI A,(HL)
//  LD (HLI),A
//   Same as: LDI (HL),A
//  LD (HL+),A
//   Same as: LDI (HL),A
// LD n,nn
// Use with:
//  n = BC,DE,HL,SP
//  nn = 16 bit immediate value
//
// LD SP,HL
// Description:
// Put HL into Stack Pointer (SP).
//
// 3. LD HL,SP+n
// Description: Same as: LDHL SP,n
//
// 5. LD (nn),SP
// Description:
//  Put Stack Pointer (SP) at address n.
// Use with:
//  nn = two byte immediate address.
func (c *CPU) LD(params ...Param) {
	if dst, is8 := params[0].(Value8); is8 {
		if src, is8 := params[1].(Value8); is8 {
			dst.Write8(src.Read8())
			return
		}
	}
	if dst, is16 := params[0].(Value16); is16 {
		if params[1] == c.C {
			dst.Write16(0xFF00 + uint16(c.C.Read8()))
			return
		}
		src := params[1].(Value16)
		value := src.Read16()
		dst.Write16(value)
	}
}

// LDI loads src into dst and increments dst
// VARIANTS:
//  LDI A,(HL)
//   Same as: LD A,(HL) - INC HL
//  LDI (HL),A
//   Same as: LD (HL),A - INC HL
func (c *CPU) LDI(params ...Param) {
	dst := params[0].(Value8)
	src := params[1].(Value8)
	dst.Write8(src.Read8())
	if mem, isMem := dst.(*Memory); isMem {
		index := mem.GetIndex().(Value16)
		index.Write16(index.Read16() + 1)
	}
	if mem, isMem := src.(*Memory); isMem {
		index := mem.GetIndex().(Value16)
		index.Write16(index.Read16() + 1)
	}
}

// LDD loads src into dst and decrements src
func (c *CPU) LDD(params ...Param) {
	dst := params[0].(Value8)
	src := params[1].(Value8)
	dst.Write8(src.Read8())
	if mem, isMem := dst.(*Memory); isMem {
		index := mem.GetIndex().(Value16)
		index.Write16(index.Read16() - 1)
	}
	if mem, isMem := src.(*Memory); isMem {
		index := mem.GetIndex().(Value16)
		index.Write16(index.Read16() - 1)
	}
}

// LDH loads src into memory address $FF00+dst
//
// 19. LDH (n),A
// Description:
// Put A into memory address $FF00+n.
//
// 20. LDH A,(n)
// Description:
// Put memory address $FF00+n into A.
func (c *CPU) LDH(params ...Param) {
	c.LD(params...)
}

// LDHL puts a + b effective address into HL
//
// 4. LDHL SP,n
// Description:
// Put SP + n effective address into HL.
// Use with:
// n = one byte signed immediate value.
// Flags affected:
//  Z - Reset.
//  N - Reset.
//  H - Set or reset according to operation.
//  C - Set or reset according to operation.
func (c *CPU) LDHL(params ...Param) {
	sp := params[0].(Value16)
	n := params[1].(ValueSigned8)

	c.F.SetZ(false)
	c.F.SetN(false)

	vSP := sp.Read16()
	vN := int16(n.ReadSigned8())

	total := int32(vSP) + int32(vN)

	halfCarry := int32(vSP&0xFFF)+int32(vN&0xFFF) > 0xFFF
	carry := total > 0xFFFF

	c.F.SetH(halfCarry)
	c.F.SetC(carry)

	valueOut := c.MMU.Read16(uint16(total))
	c.HL.Write16(valueOut)
}

// PUSH pushes nn onto the stack.
// The stack pointer is decremented twice.
//
// Use with:
//  nn = AF,BC,DE,HL
func (c *CPU) PUSH(params ...Param) {
	nn := params[0].(Value16)
	v := nn.Read16()
	high := byte((0xFF00 & v) >> 8)
	low := byte(0xFF & v)
	c.SP.Inc(-1)
	m := c.MemoryAt(c.SP)
	m.Write8(high)
	c.SP.Inc(-1)
	m = c.MemoryAt(c.SP)
	m.Write8(low)
}

// POP pops from the stack into nn
//
// Use with:
// nn = AF,BC,DE,HL
func (c *CPU) POP(params ...Param) {
	nn := params[0].(Value16)
	m := c.MemoryAt(c.SP)
	low := m.Read8()
	m = c.MemoryAt(c.SP)
	c.SP.Inc(1)
	high := m.Read8()
	v := uint16(low) + (uint16(high) << 8)
	nn.Write16(v)
	c.SP.Inc(1)
}

// ADD adds n to A
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Set if carry from bit 3.
//  C - Set if carry from bit 7.
//
// 1. ADD HL,n
// Description:
//  Add n to HL.
// Use with:
//  n = BC,DE,HL,SP
// Flags affected:
//  Z - Not affected.
//  N - Reset.
//  H - Set if carry from bit 11.
//  C - Set if carry from bit 15.
//
// 2. ADD SP,n
// Description:
//  Add n to Stack Pointer (SP).
// Use with:
//  n = one byte signed immediate value (#).
// Flags affected:
//  Z - Reset.
//  N - Reset.
//  H - Set or reset according to operation.
//  C - Set or reset according to operation.
func (c *CPU) ADD(params ...Param) {
	// 16 bit
	if dst, is16Bit := params[0].(Value16); is16Bit {
		x := dst.Read16()
		var y uint16
		if src, is16Bit := params[1].(Value16); is16Bit {
			y = src.Read16()
		}
		if src, is8Bit := params[1].(Value8); is8Bit {
			y = uint16(src.Read8())
		}
		result := x + y
		c.F.SetZ(result == 0)
		c.F.SetN(false)

		halfCarry := ((x & 0xFFF) + (y & 0xFFF)) > 0xFFF
		carry := (uint32(x) + uint32(y)) > 0xFFFF
		c.F.SetH(halfCarry)
		c.F.SetC(carry)

		dst.Write16(result)
		return
	}

	// 8 bit
	if dst, is8Bit := params[0].(Value8); is8Bit {
		n := params[1].(Value8)
		x := n.Read8()
		y := dst.Read8()
		result := y + x
		c.F.SetZ(result == 0)
		c.F.SetN(false)

		halfCarry := ((y & 0xF) + (x & 0xF)) > 0xF
		carry := (uint16(y) + uint16(x)) > 0xFF
		c.F.SetH(halfCarry)
		c.F.SetC(carry)

		dst.Write8(result)
		return
	}
}

// ADC adds src+carry flag to A
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
//  Flags affected:
//   Z - Set if result is zero.
//   N - Reset.
//   H - Set if carry from bit 3.
//   C - Set if carry from bit 7.
func (c *CPU) ADC(params ...Param) {
	n := params[0].(Value8)
	a := c.A.Read8()
	in := n.Read8()

	var carryValue byte
	if c.F.C() {
		carryValue = 1
	}

	result := a + in + carryValue
	c.F.SetZ(result == 0)
	c.F.SetN(false)

	halfCarry := ((a & 0xF) + (in & 0xF) + carryValue) > 0xF
	carry := (uint16(a) + uint16(in) + uint16(carryValue)) > 0xFF
	c.F.SetH(halfCarry)
	c.F.SetC(carry)
	c.A.Write8(result)
}

// SUB subtracts n from A
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero.
//  N - Set.
//  H - Set if no borrow from bit 4.
//  C - Set if no borrow.
func (c *CPU) SUB(params ...Param) {
	n := params[0].(Value8)
	a := c.A.Read8()
	in := n.Read8()
	carry := a < in
	result := a - in
	c.F.SetZ(result == 0)
	c.F.SetN(true)

	halfCarry := (a & 0xF) < (in & 0xF)
	c.F.SetH(!halfCarry)
	c.F.SetC(!carry)

	c.A.Write8(result)
}

// SBC subtracts n+carry flag from A
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//   Z - Set if result is zero.
//   N - Set.
//   H - Set if no borrow from bit 4.
//   C - Set if no borrow.
func (c *CPU) SBC(params ...Param) {
	n := params[0].(Value8)
	a := c.A.Read8()
	in := n.Read8()

	var carryValue byte
	if c.F.C() {
		carryValue = 1
	}
	sub := (in + carryValue)
	result := a - sub
	c.F.SetZ(result == 0)
	c.F.SetN(true)

	halfCarry := (a & 0xF) < (sub & 0xF)
	carry := uint16(a) < uint16(sub)
	c.F.SetH(!halfCarry)
	c.F.SetC(!carry)

	c.A.Write8(result)
}

// AND locally ANDs n with A and stores the result in A
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Set.
//  C - Reset.
func (c *CPU) AND(params ...Param) {
	n := params[0].(Value8)
	result := n.Read8() & c.A.Read8()
	c.F.SetZ(result == 0)
	c.F.SetN(false)
	c.F.SetH(true)
	c.F.SetC(false)
	c.A.Write8(result)
}

// OR locally ORs n with A and stores the result in A
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Reset.
func (c *CPU) OR(params ...Param) {
	n := params[0].(Value8)
	result := n.Read8() | c.A.Read8()
	c.F.SetZ(result == 0)
	c.F.SetN(false)
	c.F.SetH(false)
	c.F.SetC(false)
	c.A.Write8(result)
}

// XOR locally XORs n with A and stores the result in A
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Reset.
func (c *CPU) XOR(params ...Param) {
	n := params[0].(Value8)
	result := n.Read8() ^ c.A.Read8()
	c.F.SetZ(result == 0)
	c.F.SetN(false)
	c.F.SetH(false)
	c.F.SetC(false)
	c.A.Write8(result)
}

// CP compares A with n. This is basically A-n but the results are thrown away.
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero. (Set if A = n.)
//  N - Set.
//  H - Set if no borrow from bit 4.
//  C - Set for no borrow. (Set if A < n.)
func (c *CPU) CP(params ...Param) {
	n := params[0].(Value8)
	a := c.A.Read8()
	in := n.Read8()
	c.F.SetZ(a == in)
	c.F.SetN(true)

	halfCarry := (a & 0xF) < (in & 0xF)
	carry := uint16(a) < uint16(in)
	c.F.SetH(!halfCarry)
	c.F.SetC(!carry)
}

// INC increments n
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Set if carry from bit 3.
//  C - Not affected.
//
// 3. INC nn
// Description:
//  Increment register nn.
// Use with:
//  nn = BC,DE,HL,SP
// Flags affected:
//  None.
func (c *CPU) INC(params ...Param) {
	if n, is8Bit := params[0].(Value8); is8Bit {
		in := n.Read8()
		result := in + 1
		c.F.SetZ(result == 0)
		c.F.SetN(false)

		halfCarry := (1 + (in & 0xF)) > 0xF
		c.F.SetH(halfCarry)

		n.Write8(result)
	}

	if n, is16Bit := params[0].(Value16); is16Bit {
		in := n.Read16()
		result := in + 1
		n.Write16(result)
	}
}

// DEC decrements n
//
// n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Set.
//  H - Set if no borrow from bit 4.
//  C - Not affected.
//
// 4. DEC nn
// Description:
//  Decrement register nn.
// Use with:
//  nn = BC,DE,HL,SP
// Flags affected:
//  None.
func (c *CPU) DEC(params ...Param) {
	if n, is8Bit := params[0].(Value8); is8Bit {
		in := n.Read8()
		c.F.SetZ(in == 1)
		result := in - 1
		if in == 0 {
			result = 0xFF
		}
		c.F.SetN(true)

		halfCarry := (in & 0xF) < 1
		c.F.SetH(!halfCarry)

		n.Write8(result)
		return
	}

	if n, is16Bit := params[0].(Value16); is16Bit {
		in := n.Read16()
		result := in - 1
		n.Write16(result)
	}
}

// SWAP swaps the uppper and lower nibbles of n
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Reset.
func (c *CPU) SWAP(params ...Param) {
	n := params[0].(Value8)
	high := (n.Read8() & 0xF0) >> 4
	low := (n.Read8() & 0xF)

	n.Write8(low<<4 | high)
}

// DAA decimal adjusts A.
// This instruction adjusts register A so that the
// correct representation of Binary Coded Decimal (BCD)
// is obtained
//
// Flags affected:
//  Z - Set if register A is zero.
//  N - Not affected.
//  H - Reset.
//  C - Set or reset according to operation.
func (c *CPU) DAA(...Param) {
	value := c.A.Read8()
	var correction byte
	if c.F.H() || (!c.F.N() && (value&0xF) > 0x9) {
		correction |= 0x06
	}
	if c.F.C() || (!c.F.N() && (value) > 0x99) {
		correction |= 0x60
		c.F.SetC(true)
	} else {
		c.F.SetC(false)
	}

	if c.F.N() {
		value -= correction
	} else {
		value += correction
	}

	c.F.SetZ(value == 0)
	c.F.SetH(false)
	c.A.Write8(value)
}

// CPL complements register A (flips all bits)
//
// Flags affected:
//  Z - Not affected.
//  N - Set.
//  H - Set.
//  C - Not affected.
func (c *CPU) CPL(...Param) {
	c.A.Write8(c.A.Read8() ^ 0xFF)
}

// CCF complements the carry flag
// If C flag is set, then reset it.
// If C flag is reset, then set it.
//
// Flags affected:
//  Z - Not affected.
//  N - Reset.
//  H - Reset.
//  C - Complemented.
func (c *CPU) CCF(...Param) {
	c.F.SetN(false)
	c.F.SetH(false)
	c.F.SetC(!c.F.C())
}

// SCF sets the carry flag
//
// Flags affected:
//  Z - Not affected.
//  N - Reset.
//  H - Reset.
//  C - Set.
func (c *CPU) SCF(...Param) {
	c.F.SetN(false)
	c.F.SetH(false)
	c.F.SetC(true)
}

// HALT powers down the CPU until an interrupt occurs
// Used to reduce energy consumption
func (c *CPU) HALT(...Param) {
	c.isHalted = true
}

// STOP halts the CPU and LCD until a button is pressed
func (c *CPU) STOP(...Param) {
	c.isStopped = true
}

// DI disables interrupts but not
// immediately. Interrupts are disabled after
// instruction after DI is executed.
func (c *CPU) DI(...Param) {
	// TODO: Delay interrupts
	c.IME = false
}

// EI enables interrupts. This intruction enables interrupts
// but not immediately. Interrupts are enabled after
// instruction after EI is executed.
func (c *CPU) EI(...Param) {
	// TODO: Delay interrupts
	c.IME = true
}

// RLCA rotates A left.
// The carry flag is set to the previous bit 7.
//
// Flags affected:
// Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 7 data.
func (c *CPU) RLCA(...Param) {
	c.RLC(c.A)
	c.F.SetZ(false)
}

// RLA rotates A left through carry flag.
//
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 7 data.
func (c *CPU) RLA(...Param) {
	c.RL(c.A)
	c.F.SetZ(false)
}

// RRCA rotates A right.
// Old bit 0 to Carry flag.
//
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) RRCA(...Param) {
	c.RRC(c.A)
	c.F.SetZ(false)
}

// RRA rotates A right through carry flag.
//
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) RRA(...Param) {
	c.RR(c.A)
	c.F.SetZ(false)
}

// RLC rotates n left. Old bit 7 to Carry flag.
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 7 data.
func (c *CPU) RLC(params ...Param) {
	n := params[0].(Value8)
	value := n.Read8()
	msb := value & (0x1 << 7)
	result := value<<1 | (msb >> 7)
	c.F.SetZ(result == 0)
	c.F.SetN(false)
	c.F.SetH(false)
	c.F.SetC(msb > 0)
	n.Write8(result)
}

// RL rotates n left through carry flag.
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 7 data.
func (c *CPU) RL(params ...Param) {
	n := params[0].(Value8)
	value := n.Read8()
	msb := value & (0x1 << 7)
	result := value << 1
	// Add carry bit
	if c.F.C() {
		result++
	}
	c.F.SetZ(result == 0)
	c.F.SetN(false)
	c.F.SetH(false)
	c.F.SetC(msb > 0)
	n.Write8(result)
}

// RRC rotates n right. Old bit 0 to Carry flag.
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) RRC(params ...Param) {
	n := params[0].(Value8)
	value := n.Read8()
	lsb := value & 0x1
	result := value>>1 | lsb
	c.F.SetZ(result == 0)
	c.F.SetN(false)
	c.F.SetH(false)
	c.F.SetC(lsb > 0)
	n.Write8(result)
}

// RR rotates n right through Carry flag.
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) RR(params ...Param) {
	n := params[0].(Value8)
	value := n.Read8()
	lsb := value & 0x1
	result := value >> 1
	// Add carry bit
	if c.F.C() {
		result = result | (0x1 << 7)
	}
	c.F.SetZ(result == 0)
	c.F.SetN(false)
	c.F.SetH(false)
	c.F.SetC(lsb > 0)
	n.Write8(result)
}

// SLA shifts n left into Carry. LSB of n set to 0
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 7 data.
func (c *CPU) SLA(params ...Param) {
	n := params[0].(Value8)
	shifted := uint16(n.Read8()) << 1
	c.F.SetZ(shifted == 0)
	c.F.SetN(false)
	c.F.SetH(false)
	c.F.SetC(shifted > 0xFF)
	n.Write8(byte(shifted & 0xFF))
}

// SRA shifts n right into Carry. MSB doesn't change.
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) SRA(params ...Param) {
	n := params[0].(Value8)
	value := n.Read8()
	msb := value & (0x1 << 7)
	c.F.SetC(value&0x1 > 0)
	shifted := value >> 1
	c.F.SetZ(shifted == 0)
	c.F.SetN(false)
	c.F.SetH(false)
	n.Write8(byte(shifted | msb))
}

// SRL shifts n right into Carry. MSB set to 0.
//
// Use with:
// n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) SRL(params ...Param) {
	n := params[0].(Value8)
	value := n.Read8()
	c.F.SetC(value&0x1 > 0)
	shifted := value >> 1
	c.F.SetZ(shifted == 0)
	c.F.SetN(false)
	c.F.SetH(false)
	n.Write8(shifted)
}

// BIT tests bit b in register r.
//
// Use with:
//  b = 0 - 7, r = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if bit b of register r is 0.
//  N - Reset.
//  H - Set.
//  C - Not affected.
func (c *CPU) BIT(params ...Param) {
	pos := byte(params[0].(int))
	value := params[1].(Value8).Read8()
	result := bitValue(pos, value) > 0
	c.F.SetZ(result)
}

// SET sets bit b in register r.
//
// Use with:
//  b = 0 - 7, r = A,B,C,D,E,H,L,(HL)
func (c *CPU) SET(params ...Param) {
	pos := byte(params[0].(int))
	r := params[1].(Value8)
	value := r.Read8()
	value |= (1 << pos)
	r.Write8(value)
}

// RES resets bit b in register r.
//
// Use with:
//  b = 0 - 7, r = A,B,C,D,E,H,L,(HL)
func (c *CPU) RES(params ...Param) {
	pos := byte(params[0].(int))
	r := params[1].(Value8)
	value := r.Read8()
	value &^= (1 << pos)
	r.Write8(value)
}

// JP jumps to address nn
//
// Use with:
//  nn = two byte immediate value. (LS byte first.)
func (c *CPU) JP(params ...Param) {
	nn := params[0].(Value16)
	c.PC.Write16(nn.Read16())
}

// JPC jumps to address nn if following condition is true:
// cc = NZ, Jump if Z flag is reset.
// cc = Z, Jump if Z flag is set.
// cc = NC, Jump if C flag is reset.
// cc = C, Jump if C flag is set.
//
// Use with:
//  nn = two byte immediate value. (LS byte first.)
func (c *CPU) JPC(params ...Param) {
	if c.conditionMet(params...) {
		c.JP(params[1:]...)
	}
}

// JR adds n to current address and jumps to it.
//
// Use with:
//  n = one byte signed immediate value
func (c *CPU) JR(params ...Param) {
	n := params[0].(ValueSigned8)
	current := c.PC.Read16()
	v := int16(current) + int16(n.ReadSigned8())
	c.PC.Write16(uint16(v))
}

// JRC will, if following condition is true, add n to current
// address and jump to it.
// cc = NZ, Jump if Z flag is reset.
// cc = Z, Jump if Z flag is set.
// cc = NC, Jump if C flag is reset.
// cc = C, Jump if C flag is set.
func (c *CPU) JRC(params ...Param) {
	if c.conditionMet(params...) {
		c.JR(params[1:]...)
	}
}

// CALL pushes address of next instruction onto stack and then
// jumps to address n.
//
// Use with:
//  nn = two byte immediate value. (LS byte first.)
func (c *CPU) CALL(params ...Param) {
	n := params[0].(Value16)
	dst := n.Read16()
	c.PC.Inc(1)
	c.PUSH(c.PC)
	c.PC.Write16(dst)
}

// CALLC calls address n if following condition is true:
// cc = NZ, Call if Z flag is reset.
// cc = Z, Call if Z flag is set.
// cc = NC, Call if C flag is reset.
// cc = C, Call if C flag is set.
//
// Use with:
//  nn = two byte immediate value. (LS byte first.)
func (c *CPU) CALLC(params ...Param) {
	if c.conditionMet(params...) {
		c.CALL(params[1:]...)
	}
}

// RST pushes present address onto stack.
// Jumps to address $0000 + n.
//
// Use with:
//  n = $00,$08,$10,$18,$20,$28,$30,$38
func (c *CPU) RST(params ...Param) {
	c.PUSH(c.PC)
	i := params[0].(int)
	c.PC.Write16(uint16(i))
}

// RET pops two bytes from stack & jumps to that address.
func (c *CPU) RET(...Param) {
	c.POP(c.PC)
	c.PC.Inc(-1)
}

// RETC returns if following condition is true:
// cc = NZ, Return if Z flag is reset.
// cc = Z, Return if Z flag is set.
// cc = NC, Return if C flag is reset.
// cc = C, Return if C flag is set.
func (c *CPU) RETC(params ...Param) {
	if c.conditionMet(params...) {
		c.RET(params...)
	}
}

// RETI pops two bytes from stack & jumps to that address then
// enables interrupts.
func (c *CPU) RETI(...Param) {
	c.RET()
	c.EI()
}

// PREFIX is a placeholder for prefixing an opcode
func (c *CPU) PREFIX(...Param) {}

func (c *CPU) conditionMet(params ...Param) bool {
	switch params[0] {
	case CaseC:
		return c.F.C()
	case CaseNC:
		return !c.F.C()
	case CaseZ:
		return c.F.Z()
	case CaseNZ:
		return !c.F.Z()
	default:
		return false
	}
}
