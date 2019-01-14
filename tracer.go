package gameboy

import (
	"fmt"
	"runtime/debug"
	"strings"
)

type LogTracer interface {
	Logf(string, ...interface{})
}
type LogMessage struct {
	Text  string
	Stack []byte
}

type TraceMessage struct {
	Count     int64
	CPU       *CPUEvent
	MMU       *MMUEvent
	Registers []RegisterEvent
	Stack     []StackEvent
	Log       []LogMessage
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

type StackEvent struct {
	Pos      uint16
	ValueIn  uint16
	ValueOut uint16
}

func (s StackEvent) String() string {
	if s.ValueIn != 0 {
		return fmt.Sprintf("(SP)=0x%04X<-0x%04X", s.Pos, s.ValueIn)
	}
	return fmt.Sprintf("(SP)=0x%04X->0x%04X", s.Pos, s.ValueOut)
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

	currentMessage TraceMessage
}

func NewTracer() *Tracer {
	return &Tracer{}
}

func (t *Tracer) AddRegister(name string, valueBefore byte, valueAfter byte) {
	t.currentMessage.Registers = append(
		t.currentMessage.Registers,
		RegisterEvent{
			Name:        name,
			ValueBefore: valueBefore,
			ValueAfter:  valueAfter,
		},
	)
}

func (t *Tracer) AddCPU(pc uint16, description string) {
	t.currentMessage.CPU = &CPUEvent{
		PC:          pc,
		Description: description,
	}
}

func (t *Tracer) AddStack(pos, in, out uint16) {
	t.currentMessage.Stack = append(
		t.currentMessage.Stack,
		StackEvent{
			Pos:      pos,
			ValueIn:  in,
			ValueOut: out,
		},
	)
}

func (t *Tracer) AddMMU(pos uint16, values ...byte) {
	t.currentMessage.MMU = &MMUEvent{
		Pos:         pos,
		ValuesAfter: values,
	}
}

func (t *Tracer) Logf(message string, args ...interface{}) {
	t.currentMessage.Log = append(
		t.currentMessage.Log,
		LogMessage{
			Text:  fmt.Sprintf(message, args...),
			Stack: debug.Stack(),
		},
	)
}

func (t *Tracer) Reset() {
	t.currentMessage = TraceMessage{}
}

// Flush passes the trace data for this cycle to the relevant handler
func (t *Tracer) Flush() {
	if t.Logger == nil {
		return
	}

	msg := t.currentMessage
	msg.Count = t.Count
	t.Logger(msg)
	t.Count++
	t.Reset()
}
