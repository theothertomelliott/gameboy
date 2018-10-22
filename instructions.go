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
		src := params[1].(Value8)
		dst.Write(src.Read())
	}
	if dst, is16 := params[0].(Value16); is16 {
		src := params[1].(Value16)
		dst.Write(src.Read())
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
	dst.Write(src.Read() + 1)
}

// LDD loads src into dst and decrements src
func (c *CPU) LDD(params ...Param) {
	dst := params[0].(Value8)
	src := params[1].(Value8)
	dst.Write(src.Read())
	dst.Write(dst.Read() - 1)
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
	n := params[1].(Value8)

	c.F.SetZ(false)
	c.F.SetN(false)

	// TODO: Set carry and half carry flag

	c.HL.Write(sp.Read() + uint16(n.Read()))
}

// PUSH pushes nn onto the stack.
// The stack pointer is decremented twice.
//
// Use with:
//  nn = AF,BC,DE,HL
func (c *CPU) PUSH(params ...Param) {
	nn := params[0].(Value16)
	m := c.MemoryAt16(c.SP)
	m.Write(nn.Read())
	c.SP.Inc(-2)
}

// POP pops from the stack into nn
//
// Use with:
// nn = AF,BC,DE,HL
func (c *CPU) POP(params ...Param) {
	nn := params[0].(Value16)
	m := c.MemoryAt16(c.SP)
	nn.Write(m.Read())
	c.SP.Inc(2)
}

// ADD adds n to A
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Set if carry from bit 3.
//  C - Set if carry from bit 7.
func (c *CPU) ADD(params ...Param) {
	n := params[0].(Value8)
	a := c.A.Read()
	in := n.Read()
	result := a + in
	c.F.SetZ(result == 0)
	c.F.SetN(false)

	halfCarry := ((a & 0xF) + (in & 0xF)) > 0xF
	carry := (uint16(a) + uint16(in)) > 0xFF
	c.F.SetH(halfCarry)
	c.F.SetC(carry)

	c.A.Write(result)
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
	a := c.A.Read()
	in := n.Read()

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
	c.A.Write(result)
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
	a := c.A.Read()
	in := n.Read()
	result := a - in
	c.F.SetZ(result == 0)
	c.F.SetN(true)

	halfCarry := ((a & 0xF) - (in & 0xF)) > 0xF
	carry := (uint16(a) - uint16(in)) > 0xFF
	c.F.SetH(!halfCarry)
	c.F.SetC(!carry)

	c.A.Write(result)
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
	a := c.A.Read()
	in := n.Read()

	var carryValue byte
	if c.F.C() {
		carryValue = 1
	}

	result := a - (in + carryValue)
	c.F.SetZ(result == 0)
	c.F.SetN(true)

	halfCarry := ((a & 0xF) - (in & 0xF)) > 0xF
	carry := (uint16(a) - uint16(in)) > 0xFF
	c.F.SetH(!halfCarry)
	c.F.SetC(!carry)

	c.A.Write(result)
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
	result := n.Read() & c.A.Read()
	c.F.SetZ(result == 0)
	c.F.SetN(false)
	c.F.SetH(true)
	c.F.SetC(false)
	c.A.Write(result)
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
	result := n.Read() | c.A.Read()
	c.F.SetZ(result == 0)
	c.F.SetN(false)
	c.F.SetH(false)
	c.F.SetC(false)
	c.A.Write(result)
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
	result := n.Read() ^ c.A.Read()
	c.F.SetZ(result == 0)
	c.F.SetN(false)
	c.F.SetH(false)
	c.F.SetC(false)
	c.A.Write(result)
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
	a := c.A.Read()
	in := n.Read()
	result := a - in
	c.F.SetZ(result == 0)
	c.F.SetN(true)

	halfCarry := ((a & 0xF) - (in & 0xF)) > 0xF
	carry := (uint16(a) - uint16(in)) > 0xFF
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
func (c *CPU) INC(...Param) {

}

// DEC decrements n
//
// n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Set.
//  H - Set if no borrow from bit 4.
//  C - Not affected.
func (c *CPU) DEC(...Param) {}

// ADDHL adds n to HL
//
// Use with:
//  n = BC,DE,HL,SP
// Flags affected:
//  Z - Not affected.
//  N - Reset.
//  H - Set if carry from bit 11.
//  C - Set if carry from bit 15.
func (c *CPU) ADDHL(...Param) {}

// ADDSP adds n to the stack pointer
//
// Use with:
//  n = one byte signed immediate value (#).
// Flags affected:
//  Z - Reset.
//  N - Reset.
//  H - Set or reset according to operation.
//  C - Set or reset according to operation.
func (c *CPU) ADDSP(...Param) {}

// SWAP swaps the uppper and lower nibbles of n
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Reset.
func (c *CPU) SWAP(...Param) {}

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
func (c *CPU) DAA(...Param) {}

// CPL complements register A (flips all bits)
//
// Flags affected:
//  Z - Not affected.
//  N - Set.
//  H - Set.
//  C - Not affected.
func (c *CPU) CPL(...Param) {}

// CCF complements the carry flag
// If C flag is set, then reset it.
// If C flag is reset, then set it.
//
// Flags affected:
//  Z - Not affected.
//  N - Reset.
//  H - Reset.
//  C - Complemented.
func (c *CPU) CCF(...Param) {}

// SCF sets the carry flag
//
// Flags affected:
//  Z - Not affected.
//  N - Reset.
//  H - Reset.
//  C - Set.
func (c *CPU) SCF(...Param) {}

// HALT powers down the CPU until an interrupt occurs
// Used to reduce energy consumption
func (c *CPU) HALT(...Param) {}

// STOP halts the CPU and LCD until a button is pressed
func (c *CPU) STOP(...Param) {}

// DI disables interrupts but not
// immediately. Interrupts are disabled after
// instruction after DI is executed.
func (c *CPU) DI(...Param) {}

// EI enables interrupts. This intruction enables interrupts
// but not immediately. Interrupts are enabled after
// instruction after EI is executed.
func (c *CPU) EI(...Param) {}

// RLCA rotates A left.
// The carry flag is set to the previous bit 7.
//
// Flags affected:
// Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 7 data.
func (c *CPU) RLCA(...Param) {}

// RLA rotates A left through carry flag.
//
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 7 data.
func (c *CPU) RLA(...Param) {}

// RRCA rotates A right.
// Old bit 0 to Carry flag.
//
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) RRCA(...Param) {}

// RRA rotates A right through carry flag.
//
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) RRA(...Param) {}

// RLC rotates n left. Old bit 7 to Carry flag.
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 7 data.
func (c *CPU) RLC(...Param) {}

// RL rotates n left through carry flag.
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 7 data.
func (c *CPU) RL(...Param) {}

// RRC rotates n right. Old bit 0 to Carry flag.
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) RRC(...Param) {}

// RR rotates n right through Carry flag.
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) RR(...Param) {}

// SLA shifts n left into Carry. LSB of n set to 0
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 7 data.
func (c *CPU) SLA(...Param) {}

// SRA shifts n right into Carry. MSB doesn't change.
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) SRA(...Param) {}

// SRL shifts n right into Carry. MSB set to 0.
//
// Use with:
// n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) SRL(...Param) {}

// BIT tests bit b in register r.
//
// Use with:
//  b = 0 - 7, r = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if bit b of register r is 0.
//  N - Reset.
//  H - Set.
//  C - Not affected.
func (c *CPU) BIT(...Param) {}

// SET sets bit b in register r.
//
// Use with:
//  b = 0 - 7, r = A,B,C,D,E,H,L,(HL)
func (c *CPU) SET(...Param) {}

// RES resets bit b in register r.
//
// Use with:
//  b = 0 - 7, r = A,B,C,D,E,H,L,(HL)
func (c *CPU) RES(...Param) {}

// JP jumps to address nn
//
// Use with:
//  nn = two byte immediate value. (LS byte first.)
func (c *CPU) JP(...Param) {}

// JPC jumps to address nn if following condition is true:
// cc = NZ, Jump if Z flag is reset.
// cc = Z, Jump if Z flag is set.
// cc = NC, Jump if C flag is reset.
// cc = C, Jump if C flag is set.
//
// Use with:
//  nn = two byte immediate value. (LS byte first.)
func (c *CPU) JPC(...Param) {}

// JPHL jumps to the address contained in HL
func (c *CPU) JPHL(...Param) {}

// JR adds n to current address and jumps to it.
//
// Use with:
//  n = one byte signed immediate value
func (c *CPU) JR(...Param) {}

// JRC will, if following condition is true, add n to current
// address and jump to it.
// cc = NZ, Jump if Z flag is reset.
// cc = Z, Jump if Z flag is set.
// cc = NC, Jump if C flag is reset.
// cc = C, Jump if C flag is set.
func (c *CPU) JRC(...Param) {}

// CALL pushes address of next instruction onto stack and then
// jumps to address n.
//
// Use with:
//  nn = two byte immediate value. (LS byte first.)
func (c *CPU) CALL(...Param) {}

// CALLC calls address n if following condition is true:
// cc = NZ, Call if Z flag is reset.
// cc = Z, Call if Z flag is set.
// cc = NC, Call if C flag is reset.
// cc = C, Call if C flag is set.
//
// Use with:
//  nn = two byte immediate value. (LS byte first.)
func (c *CPU) CALLC(...Param) {}

// RST pushes present address onto stack.
// Jumps to address $0000 + n.
//
// Use with:
//  n = $00,$08,$10,$18,$20,$28,$30,$38
func (c *CPU) RST(...Param) {}

// RET pops two bytes from stack & jumps to that address.
func (c *CPU) RET(...Param) {}

// RETC returns if following condition is true:
// cc = NZ, Return if Z flag is reset.
// cc = Z, Return if Z flag is set.
// cc = NC, Return if C flag is reset.
// cc = C, Return if C flag is set.
func (c *CPU) RETC(...Param) {}

// RETI pops two bytes from stack & jumps to that address then
// enables interrupts.
func (c *CPU) RETI(...Param) {}

// PREFIX is a placeholder for prefixing an opcode
func (c *CPU) PREFIX(...Param) {}
