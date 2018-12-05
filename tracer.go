package gameboy

type TraceMessage struct {
	Count int64
	CPU   *CPUEvent
	MMU   *MMUEvent
}

type CPUEvent struct {
	PC          uint16
	Description string
}

type MMUEvent struct {
	Pos    uint16
	Values []byte
}

type Tracer struct {
	Count int64

	Logger func(ev TraceMessage)

	CurrentCPU *CPUEvent
	CurrentMMU *MMUEvent
}

func NewTracer() *Tracer {
	return &Tracer{}
}

func (t *Tracer) AddCPU(pc uint16, description string) {
	t.CurrentCPU = &CPUEvent{
		PC:          pc,
		Description: description,
	}
}

func (t *Tracer) AddMMU(pos uint16, values ...byte) {
	t.CurrentMMU = &MMUEvent{
		Pos:    pos,
		Values: values,
	}
}

func (t *Tracer) Log() {
	if t.Logger == nil {
		return
	}

	t.Logger(TraceMessage{
		Count: t.Count,
		CPU:   t.CurrentCPU,
		MMU:   t.CurrentMMU,
	})
	t.CurrentCPU = nil
	t.CurrentMMU = nil
	t.Count++
}
