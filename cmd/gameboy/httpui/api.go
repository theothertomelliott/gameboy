package httpui

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

func (s *Server) RegisterAPI() error {
	http.HandleFunc("/api/cpu", s.HandleAPICPU)
	http.HandleFunc("/api/decompile", s.HandleAPIDecompile)
	http.HandleFunc("/api/stack", s.HandleAPIStack)
	http.HandleFunc("/api/trace", s.HandleAPITrace)
	return nil
}

func (s *Server) HandleAPICPU(w http.ResponseWriter, r *http.Request) {
	data := cpuData{
		Registers: s.getRegisters(),
		Paused:    s.gb.IsPaused(),
		IME:       s.gb.CPU().IME,
	}
	jsonResponse(w, data)
}

func jsonResponse(w http.ResponseWriter, body interface{}) {
	out, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(out))
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

	jsonResponse(w, ops)
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

	jsonResponse(w, stack)
}

func (s *Server) HandleAPITrace(w http.ResponseWriter, r *http.Request) {
	const pageSize = 1000

	offsetStr := r.FormValue("offset")
	var (
		offset int64
		err    error
	)
	switch offsetStr {
	case "", "last":
		lastPageLength := len(s.trace) % pageSize
		if lastPageLength == 0 {
			lastPageLength = pageSize
		}
		offset = int64(len(s.trace) - lastPageLength)
	default:
		offset, err = strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if offset < 0 {
			offset = 0
		}
	}

	var trace []traceEntry
	if int(offset) < len(s.trace) {
		if (len(s.trace) - int(offset)) < pageSize {
			trace = s.trace[offset:]
		} else {
			trace = s.trace[offset : offset+pageSize]
		}
	}

	data := searchData{
		Start:    offset,
		End:      offset + int64(len(trace)),
		Previous: offset - 1000,
		Next:     offset + 1000,
		Total:    len(s.trace),
	}

	for index, t := range trace {
		data.Trace = append(data.Trace, traceData{
			Index: offset + int64(index),
			Trace: t,
		})
	}

	jsonResponse(w, data)
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
