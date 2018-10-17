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
func (c *CPU) LD(params ...Param) {
	dst := params[0].(Value8)
	src := params[1].(Value8)
	dst.Write(src.Read())
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

}

// LD16 loads src into dst
//
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
func (c *CPU) LD16(dst, src *uint16) {}

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
func (c *CPU) LDHL(a, b *uint64) {}

// PUSH pushes nn onto the stack.
// The stack pointer is decremented twice.
//
// Use with:
//  nn = AF,BC,DE,HL
func (c *CPU) PUSH(nn *uint16) {}

// POP pops from the stack into nn
//
// Use with:
// nn = AF,BC,DE,HL
func (c *CPU) POP(nn *uint16) {}

// ADD adds n to A
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Set if carry from bit 3.
//  C - Set if carry from bit 7.
func (c *CPU) ADD(n *byte) {}

// ADDC adds src+carry flag to A
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
//  Flags affected:
//   Z - Set if result is zero.
//   N - Reset.
//   H - Set if carry from bit 3.
//   C - Set if carry from bit 7.
func (c *CPU) ADDC(n *byte) {}

// SUB subtracts n from A
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero.
//  N - Set.
//  H - Set if no borrow from bit 4.
//  C - Set if no borrow.
func (c *CPU) SUB(n *byte) {}

// SBC subtracts n+carry flag from A
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//   Z - Set if result is zero.
//   N - Set.
//   H - Set if no borrow from bit 4.
//   C - Set if no borrow.
func (c *CPU) SBC(n *byte) {}

// AND locally ANDs n with A and stores the result in A
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Set.
//  C - Reset.
func (c *CPU) AND(n *byte) {}

// OR locally ORs n with A and stores the result in A
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Reset.
func (c *CPU) OR(n *byte) {}

// XOR locally XORs n with A and stores the result in A
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Reset.
func (c *CPU) XOR(n *byte) {}

// CP compares A with n. This is basically A-n but the results are thrown away.
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL),#
// Flags affected:
//  Z - Set if result is zero. (Set if A = n.)
//  N - Set.
//  H - Set if no borrow from bit 4.
//  C - Set for no borrow. (Set if A < n.)
func (c *CPU) CP(n *byte) {}

// INC increments n
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Set if carry from bit 3.
//  C - Not affected.
func (c *CPU) INC(...Param) {}

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
func (c *CPU) ADDSP(n *byte) {}

// INC16 increments nn
//
// Use with:
//  nn = BC,DE,HL,SP
func (c *CPU) INC16(n *uint16) {}

// DEC16 decrements nn
//
// Use with:
//  nn = BC,DE,HL,SP
func (c *CPU) DEC16(nn *uint16) {}

// SWAP swaps the uppper and lower nibbles of n
//
// Use with:
//  n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Reset.
func (c *CPU) SWAP(n *byte) {}

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
func (c *CPU) DAA() {}

// CPL complements register A (flips all bits)
//
// Flags affected:
//  Z - Not affected.
//  N - Set.
//  H - Set.
//  C - Not affected.
func (c *CPU) CPL() {}

// CCF complements the carry flag
// If C flag is set, then reset it.
// If C flag is reset, then set it.
//
// Flags affected:
//  Z - Not affected.
//  N - Reset.
//  H - Reset.
//  C - Complemented.
func (c *CPU) CCF() {}

// SCF sets the carry flag
//
// Flags affected:
//  Z - Not affected.
//  N - Reset.
//  H - Reset.
//  C - Set.
func (c *CPU) SCF() {}

// HALT powers down the CPU until an interrupt occurs
// Used to reduce energy consumption
func (c *CPU) HALT() {}

// STOP halts the CPU and LCD until a button is pressed
func (c *CPU) STOP() {}

// DI disables interrupts but not
// immediately. Interrupts are disabled after
// instruction after DI is executed.
func (c *CPU) DI() {}

// EI enables interrupts. This intruction enables interrupts
// but not immediately. Interrupts are enabled after
// instruction after EI is executed.
func (c *CPU) EI() {}

// RCLA rotates A left.
// The carry flag is set to the previous bit 7.
//
// Flags affected:
// Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 7 data.
func (c *CPU) RCLA(...Param) {}

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
func (c *CPU) SRA(n *byte) {}

// SRL shifts n right into Carry. MSB set to 0.
//
// Use with:
// n = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if result is zero.
//  N - Reset.
//  H - Reset.
//  C - Contains old bit 0 data.
func (c *CPU) SRL(n *byte) {}

// BIT tests bit b in register r.
//
// Use with:
//  b = 0 - 7, r = A,B,C,D,E,H,L,(HL)
// Flags affected:
//  Z - Set if bit b of register r is 0.
//  N - Reset.
//  H - Set.
//  C - Not affected.
func (c *CPU) BIT(b, r *byte) {}

// SET sets bit b in register r.
//
// Use with:
//  b = 0 - 7, r = A,B,C,D,E,H,L,(HL)
func (c *CPU) SET(b, r *byte) {}

// RES resets bit b in register r.
//
// Use with:
//  b = 0 - 7, r = A,B,C,D,E,H,L,(HL)
func (c *CPU) RES(b, r *byte) {}

// JP jumps to address nn
//
// Use with:
//  nn = two byte immediate value. (LS byte first.)
func (c *CPU) JP(nn *uint16) {}

// JPC jumps to address nn if following condition is true:
// cc = NZ, Jump if Z flag is reset.
// cc = Z, Jump if Z flag is set.
// cc = NC, Jump if C flag is reset.
// cc = C, Jump if C flag is set.
//
// Use with:
//  nn = two byte immediate value. (LS byte first.)
func (c *CPU) JPC(cc, nn *uint16) {}

// JPHL jumps to the address contained in HL
func (c *CPU) JPHL() {}

// JR adds n to current address and jumps to it.
//
// Use with:
//  n = one byte signed immediate value
func (c *CPU) JR(n *byte) {}

// JRC will, if following condition is true, add n to current
// address and jump to it.
// cc = NZ, Jump if Z flag is reset.
// cc = Z, Jump if Z flag is set.
// cc = NC, Jump if C flag is reset.
// cc = C, Jump if C flag is set.
func (c *CPU) JRC(cc *uint16, n *byte) {}

// CALL pushes address of next instruction onto stack and then
// jumps to address n.
//
// Use with:
//  nn = two byte immediate value. (LS byte first.)
func (c *CPU) CALL(nn *uint16) {}

// CALLC calls address n if following condition is true:
// cc = NZ, Call if Z flag is reset.
// cc = Z, Call if Z flag is set.
// cc = NC, Call if C flag is reset.
// cc = C, Call if C flag is set.
//
// Use with:
//  nn = two byte immediate value. (LS byte first.)
func (c *CPU) CALLC(cc *uint16, nn *uint16) {}

// RST pushes present address onto stack.
// Jumps to address $0000 + n.
//
// Use with:
//  n = $00,$08,$10,$18,$20,$28,$30,$38
func (c *CPU) RST(n *byte) {}

// RET pops two bytes from stack & jumps to that address.
func (c *CPU) RET() {}

// RETC returns if following condition is true:
// cc = NZ, Return if Z flag is reset.
// cc = Z, Return if Z flag is set.
// cc = NC, Return if C flag is reset.
// cc = C, Return if C flag is set.
func (c *CPU) RETC(cc *uint16) {}

// RETI pops two bytes from stack & jumps to that address then
// enables interrupts.
func (c *CPU) RETI() {}
