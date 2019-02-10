package httpui

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"

	packr "github.com/gobuffalo/packr/v2"
)

// HandleDebug renders the debugger view
func (s *Server) HandleDebug(w http.ResponseWriter, r *http.Request) {
	s.decompileMtx.Lock()
	defer s.decompileMtx.Unlock()

	box := packr.New("views", "./views")
	tpl, err := box.FindString("debug.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.New("debug.html").Parse(tpl)
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
		Registers: s.getRegisters(),
		Paused:    s.gb.IsPaused(),
	}
	for _, index := range indices {
		r := row{}
		r.Index = fmt.Sprintf("%04X", index)
		r.Description = s.decompilation[index]
		r.Flags = []string{r.Index}
		r.Id = r.Index
		if index == s.gb.CPU().PC.Read16() {
			r.Id = "PC"
			r.Flags = append(r.Flags, "PC")
		}

		_, r.Breakpoint = s.gb.Breakpoints[index]

		data.Op = append(data.Op, r)
	}

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
		A: fmt.Sprintf("0x%02X", cpu.A.Read8()),
		F: fmt.Sprintf("0x%02X", cpu.F.Read8()),
		B: fmt.Sprintf("0x%02X", cpu.B.Read8()),
		C: fmt.Sprintf("0x%02X", cpu.C.Read8()),
		D: fmt.Sprintf("0x%02X", cpu.D.Read8()),
		E: fmt.Sprintf("0x%02X", cpu.E.Read8()),
		H: fmt.Sprintf("0x%02X", cpu.H.Read8()),
		L: fmt.Sprintf("0x%02X", cpu.L.Read8()),
	}
}

type (
	registers struct {
		A, F string
		B, C string
		D, E string
		H, L string
	}
	row struct {
		Index       string
		Description string
		Flags       []string
		Id          string
		Breakpoint  bool
	}
	table struct {
		Registers registers
		Op        []row
		Paused    bool
	}
)
