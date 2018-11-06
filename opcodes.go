package gameboy

type ConditionalCase int

func (c ConditionalCase) String() string {
	switch c {
	case CaseZ:
		return "Z"
	case CaseNZ:
		return "NZ"
	case CaseC:
		return "C"
	case CaseNC:
		return "NC"
	default:
		return "Unknown Conditional"
	}
}

const (
	CaseZ ConditionalCase = iota
	CaseNZ
	CaseC
	CaseNC
)

type Param interface{}

type Opcode byte
