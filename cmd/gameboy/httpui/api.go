package httpui

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func (s *Server) RegisterAPI() error {
	http.HandleFunc("/api/cpu", s.HandleAPICPU)
	http.HandleFunc("/api/decompile", s.HandleAPIDecompile)
	http.HandleFunc("/api/stack", s.HandleAPIStack)
	return nil
}

func (s *Server) HandleAPICPU(w http.ResponseWriter, r *http.Request) {
	data := cpuData{
		Registers: s.getRegisters(),
		Paused:    s.gb.IsPaused(),
		IME:       s.gb.CPU().IME,
	}
	_ = jsonResponse(w, data)
}

func jsonResponse(w http.ResponseWriter, body interface{}) error {
	out, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return err
	}
	fmt.Fprint(w, string(out))
	return nil
}

func (s *Server) HandleAPIDecompile(w http.ResponseWriter, r *http.Request) {
	s.decompileMtx.Lock()
	defer s.decompileMtx.Unlock()

	var indices []uint16
	for index := range s.decompilation {
		indices = append(indices, index)
	}
	sort.Slice(indices, func(i, j int) bool {
		return indices[i] < indices[j]
	})

	var ops []op
	for _, index := range indices {
		r := op{}
		r.Index = Uint16(index)
		r.Description = s.decompilation[index]
		ops = append(ops, r)
	}

	_ = jsonResponse(w, ops)
}

func (s *Server) HandleAPIStack(w http.ResponseWriter, r *http.Request) {
	s.stackMtx.Lock()
	defer s.stackMtx.Unlock()

	var stack []stackEntry
	for _, sp := range s.stack {
		stack = append(stack, sp)
	}
	sort.Slice(stack, func(i, j int) bool {
		return stack[i].Pos < stack[j].Pos
	})

	_ = jsonResponse(w, stack)
}

type (
	cpuData struct {
		Registers registers
		Paused    bool
		IME       bool
	}

	op struct {
		Index       Uint16
		Description string
	}
)
