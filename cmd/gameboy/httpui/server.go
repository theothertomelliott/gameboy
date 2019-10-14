package httpui

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"sync"

	packr "github.com/gobuffalo/packr/v2"
	"github.com/theothertomelliott/gameboy"
)

// NewServer creates a UI server for the provided DMG instance
func NewServer(gb *gameboy.DMG, debug bool) *Server {
	return &Server{
		gb:            gb,
		debugEnabled:  debug,
		decompilation: make(map[uint16]string),
		stack:         make(map[uint16]stackEntry),
		trace:         make([]traceEntry, 0, 10),
	}
}

type Server struct {
	gb *gameboy.DMG

	debugEnabled bool

	stack    map[uint16]stackEntry
	stackMtx sync.Mutex

	decompilation map[uint16]string
	decompileMtx  sync.Mutex

	trace []traceEntry
}

type traceEntry struct {
	Pos          Uint16
	Description  string
	MemoryValues Bytes
	Repeated     int
}

type Uint32 uint16

func (u Uint32) String() string {
	return fmt.Sprintf("%08X", uint16(u))
}

type Uint16 uint16

func (u Uint16) String() string {
	return fmt.Sprintf("%04X", uint16(u))
}

type Uint8 uint8

func (u Uint8) String() string {
	return fmt.Sprintf("%02X", uint8(u))
}

type Bytes []byte

func (b Bytes) String() string {
	var formatted = make([]string, 0, len(b))
	for _, v := range b {
		formatted = append(formatted, fmt.Sprintf("%02X '%s'", v, []byte{v}))
	}
	return strings.Join(formatted, ", ")
}

type stackEntry struct {
	Pos     Uint16
	Value   Uint16
	WriteBy Uint16
}

func (s *Server) addTraceOrRecordRepeat(t traceEntry) {
	if len(s.trace) != 0 {
		lastT := s.trace[len(s.trace)-1]
		lastTWithNoRepeat := lastT
		lastTWithNoRepeat.Repeated = 0

		if fmt.Sprint(lastTWithNoRepeat) == fmt.Sprint(t) {
			lastT.Repeated++
			s.trace[len(s.trace)-1] = lastT
			return
		}
	}
	s.trace = append(s.trace, t)
}

func (s *Server) Trace(ev gameboy.TraceMessage) {
	if !s.debugEnabled {
		return
	}

	if ev.CPU != nil {
		s.decompileMtx.Lock()
		s.decompilation[ev.CPU.PC] = ev.CPU.Description

		t := traceEntry{
			Pos:         Uint16(ev.CPU.PC),
			Description: ev.CPU.Description,
		}
		if ev.MMU != nil {
			t.MemoryValues = Bytes(ev.MMU.ValuesAfter)
		}

		s.addTraceOrRecordRepeat(t)
		s.decompileMtx.Unlock()
	}

	s.stackMtx.Lock()
	for _, st := range ev.Stack {
		if st.ValueIn == 0 || ev.CPU == nil {
			continue
		}
		s.stack[st.Pos] = stackEntry{
			Pos:     Uint16(st.Pos),
			Value:   Uint16(st.ValueIn),
			WriteBy: Uint16(ev.CPU.PC),
		}
	}
	s.stackMtx.Unlock()
}

// ListenAndServe starts a UI server on the specified port
func (s *Server) ListenAndServe(port int) error {
	http.HandleFunc("/memory", s.HandleMemory)
	http.HandleFunc("/debug", s.HandleDebug)
	http.HandleFunc("/reset", s.HandleReset)
	http.HandleFunc("/debug/togglepaused", s.HandleTogglePaused)
	http.HandleFunc("/debug/togglebreakpoint/", s.HandleToggleBreakpoint)
	http.HandleFunc("/debug/step", s.HandleStep)
	http.HandleFunc("/tiles", s.HandleTiles)
	http.HandleFunc("/trace/search", s.HandleSearchTrace)
	http.HandleFunc("/trace", s.HandleTrace)
	box := packr.New("public", "./public")
	fs := http.FileServer(box)
	http.Handle("/public/", http.StripPrefix("/public", fs))

	err := s.RegisterAPI()
	if err != nil {
		return err
	}
	http.HandleFunc("/", s.HandleIndex)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	box := packr.New("views", "./views")
	tpl, err := box.FindString("index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.New("index").Parse(tpl)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := struct {
		Items []string
	}{
		Items: []string{
			"/memory",
			"/debug",
			"/tiles",
			"/trace",
		},
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
