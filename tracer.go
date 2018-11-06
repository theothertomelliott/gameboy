package gameboy

type TraceMessage struct {
	Count int64
	Event TraceEvent
}

type TraceEvent struct {
	PC          uint16
	Description string
}

type Tracer struct {
	Count int64

	Logger func(ev TraceMessage)
}

func NewTracer() *Tracer {
	return &Tracer{}
}

func (t *Tracer) Log(ev TraceEvent) {
	if t.Logger == nil {
		return
	}

	t.Logger(TraceMessage{
		Count: t.Count,
		Event: ev,
	})
	t.Count++
}
