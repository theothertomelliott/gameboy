package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Opcode struct {
	Addr        string  `json:"addr"`
	Description string  `json:"-"`
	Mnemonic    string  `json:"mnemonic"`
	Cycles      []int   `json:"cycles"`
	Operand1    Operand `json:"operand1"`
	Operand2    Operand `json:"operand2"`
}

func (o *Opcode) UnmarshalJSON(d []byte) error {
	type Alias Opcode
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	err := json.Unmarshal(d, aux)
	if err != nil {
		return err
	}

	textValues := struct {
		Mnemonic string `json:"mnemonic"`
		Operand1 string `json:"operand1"`
		Operand2 string `json:"operand2"`
	}{}
	err = json.Unmarshal(d, &textValues)
	if err != nil {
		return err
	}
	o.Description = textValues.Mnemonic

	// Set RLA to RL A
	if o.Mnemonic == "RLA" {
		o.Mnemonic = "RL"
		o.Operand1 = "c.A"
	}

	if o.isConditional() {
		o.Mnemonic = fmt.Sprintf("%sC", o.Mnemonic)
		switch o.Operand1 {
		case "c.Z":
			o.Operand1 = "CaseZ"
		case "c.NZ":
			o.Operand1 = "CaseNZ"
		case "c.C":
			o.Operand1 = "CaseC"
		case "c.NC":
			o.Operand1 = "CaseNC"
		}
	}
	return nil
}

func (o *Opcode) isConditional() bool {
	switch o.Mnemonic {
	case "JP", "JR", "CALL":
		return o.Operand1 != "" && o.Operand2 != ""
	case "RET":
		return o.Operand1 != ""
	default:
		return false
	}
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

	s = strings.Replace(s, "r8", "R8()", 1)
	s = strings.Replace(s, "d8", "D8()", 1)
	s = strings.Replace(s, "a16", "A16()", 1)
	s = strings.Replace(s, "d16", "D16()", 1)
	s = strings.Replace(s, "a8", "A8()", 1)

	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		s = strings.Replace(s, "(", "", 1)
		s = strings.Replace(s, ")", "", 1)
		*o = Operand(fmt.Sprintf("c.MemoryAt(c.%v)", s))
		return nil
	}

	*o = Operand(fmt.Sprintf("c.%v", s))
	return nil
}
