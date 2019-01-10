package gameboy

import (
	"fmt"
	"strings"
)

type TraceMessage struct {
	Count     int64
	CPU       *CPUEvent
	MMU       *MMUEvent
	Registers []RegisterEvent
}

type CPUEvent struct {
	PC          uint16
	Description string
}

type RegisterEvent struct {
	Name        string
	ValueBefore byte
	ValueAfter  byte
}

func (r RegisterEvent) String() string {
	return fmt.Sprintf("%v: 0x%02X -> 0x%02X", r.Name, r.ValueBefore, r.ValueAfter)
}

type MMUEvent struct {
	Pos         uint16
	ValuesAfter []byte
}

func (m *MMUEvent) String() string {
	if m == nil {
		return ""
	}
	var valuesAsStr []string
	for _, value := range m.ValuesAfter {
		valuesAsStr = append(valuesAsStr, fmt.Sprintf("0x%02X", value))
	}
	return fmt.Sprintf("0x%02X: %v", m.Pos, strings.Join(valuesAsStr, " "))
}

type Tracer struct {
	Count int64

	Logger func(ev TraceMessage)

	CurrentCPU      *CPUEvent
	CurrentMMU      *MMUEvent
	CurrentRegister []RegisterEvent
}

func NewTracer() *Tracer {
	return &Tracer{}
}

func (t *Tracer) AddRegister(name string, valueBefore byte, valueAfter byte) {
	t.CurrentRegister = append(t.CurrentRegister, RegisterEvent{
		Name:        name,
		ValueBefore: valueBefore,
		ValueAfter:  valueAfter,
	})
}

func (t *Tracer) AddCPU(pc uint16, description string) {
	t.CurrentCPU = &CPUEvent{
		PC:          pc,
		Description: description,
	}
}

func (t *Tracer) AddMMU(pos uint16, values ...byte) {
	t.CurrentMMU = &MMUEvent{
		Pos:         pos,
		ValuesAfter: values,
	}
}

func (t *Tracer) Reset() {
	t.CurrentCPU = nil
	t.CurrentMMU = nil
	t.CurrentRegister = nil
}

func (t *Tracer) Log() {
	if t.Logger == nil {
		return
	}

	t.Logger(TraceMessage{
		Count:     t.Count,
		CPU:       t.CurrentCPU,
		MMU:       t.CurrentMMU,
		Registers: t.CurrentRegister,
	})
	t.Count++
	t.Reset()
}
