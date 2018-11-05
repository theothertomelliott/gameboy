package gameboy

type TraceMessage struct {
	Count int64
	Event TraceEvent
}

type TraceEvent struct {
	PC          uint16
	Description string
	FlagsBefore string
	FlagsAfter  string
}

type Tracer struct {
	Event chan TraceMessage
	Count int64
}

func NewTracer() *Tracer {
	return &Tracer{
		Event: make(chan TraceMessage, 8),
	}
}

func (t *Tracer) Log(ev TraceEvent) {
	if t.Event == nil {
		return
	}

	select {
	case t.Event <- TraceMessage{
		Count: t.Count,
		Event: ev,
	}:
	default:
	}
	t.Count++
}

func (t *Tracer) Close() {
	close(t.Event)
	t.Event = nil
}
