package httpui

import (
	"net/http"
	"sort"
)

// HandleDebug renders the debugger view
func (s *Server) HandleDebug(w http.ResponseWriter, r *http.Request) {
	s.decompileMtx.Lock()
	defer s.decompileMtx.Unlock()

	t, err := loadTemplate("debug.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var indices []uint16
	for index := range s.decompilation {
		indices = append(indices, index)
	}
	sort.Slice(indices, func(i, j int) bool {
		return indices[i] < indices[j]
	})

	data := table{
		TestOutput: s.gb.CPU().MMU.TestOutput(),
		Registers:  s.getRegisters(),
		Paused:     s.gb.IsPaused(),
	}
	for _, index := range indices {
		r := row{}
		r.Index = Uint16(index)
		r.Description = s.decompilation[index]
		r.Flags = []string{r.Index.String()}
		r.Id = r.Index.String()
		if index == s.gb.CPU().PC.Read16() {
			r.Id = "PC"
			r.Flags = append(r.Flags, "PC")
		}

		_, r.Breakpoint = s.gb.Breakpoints[index]

		data.Op = append(data.Op, r)
	}

	s.stackMtx.Lock()
	for _, sp := range s.stack {
		data.Stack = append(data.Stack, sp)
	}
	sort.Slice(data.Stack, func(i, j int) bool {
		return data.Stack[i].Pos < data.Stack[j].Pos
	})
	s.stackMtx.Unlock()

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// HandleReset resets the gameboy
func (s *Server) HandleReset(w http.ResponseWriter, r *http.Request) {
	s.gb.Reset()
	http.Redirect(w, r, "/debug", 302)
}

// HandleTogglePaused toggles the paused state of the DMB emulator and redirects to the debugger
func (s *Server) HandleTogglePaused(w http.ResponseWriter, r *http.Request) {
	s.gb.SetPaused(!s.gb.IsPaused())
	http.Redirect(w, r, "/debug", 302)
}

// HandleStep steps the DMB and redirects to the debugger
func (s *Server) HandleStep(w http.ResponseWriter, r *http.Request) {
	err := s.gb.Step()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/debug", 302)
}

func (s *Server) getRegisters() registers {
	cpu := s.gb.CPU()
	return registers{
		A: Uint8(cpu.A.Read8()),
		F: Uint8(cpu.F.Read8()),
		B: Uint8(cpu.B.Read8()),
		C: Uint8(cpu.C.Read8()),
		D: Uint8(cpu.D.Read8()),
		E: Uint8(cpu.E.Read8()),
		H: Uint8(cpu.H.Read8()),
		L: Uint8(cpu.L.Read8()),

		SP: Uint16(cpu.SP.Read16()),
		PC: Uint16(cpu.PC.Read16()),
	}
}

type (
	registers struct {
		A, F Uint8
		B, C Uint8
		D, E Uint8
		H, L Uint8

		SP Uint16
		PC Uint16
	}
	row struct {
		Index       Uint16
		Description string
		Flags       []string
		Id          string
		Breakpoint  bool
	}
	table struct {
		Stack      []stackEntry
		TestOutput string
		Registers  registers
		Op         []row
		Paused     bool
	}
)
