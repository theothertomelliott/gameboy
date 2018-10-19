package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Opcode struct {
	Addr     string  `addr`
	Mnemonic string  `mnemonic`
	Cycles   []int   `cycles`
	Operand1 Operand `operand1`
	Operand2 Operand `operand2`
}

type Operand string

func (o *Operand) UnmarshalJSON(d []byte) error {
	var s string
	err := json.Unmarshal(d, &s)
	if err != nil {
		return err
	}

	if _, err := strconv.Atoi(s); err == nil {
		*o = Operand(s)
		return err
	}
	if strings.HasSuffix(s, "H") {
		n := strings.Replace(s, "H", "", 1)
		h := fmt.Sprintf("0x%v", n)
		if _, err := strconv.ParseInt(h, 0, 64); err == nil {
			*o = Operand(h)
			return err
		}
	}

	s = strings.Replace(s, "r8", "D8", 1)
	s = strings.Replace(s, "d8", "D8", 1)
	s = strings.Replace(s, "a16", "D16", 1)
	s = strings.Replace(s, "d16", "D16", 1)

	if s == "(a8)" {
		*o = Operand("c.MemoryAtH(c.D8)")
		return nil
	}
	s = strings.Replace(s, "a8", "A8", 1)

	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		s = strings.Replace(s, "(", "", 1)
		s = strings.Replace(s, ")", "", 1)
		*o = Operand(fmt.Sprintf("c.MemoryAt(c.%v)", s))
		return nil
	}

	*o = Operand(fmt.Sprintf("c.%v", s))
	return nil
}
