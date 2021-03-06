package httpui

import (
	"net/http"
	"strconv"
	"strings"
)

type (
	traceData struct {
		Index int64
		Trace traceEntry
	}
	searchData struct {
		Trace      []traceData
		Start      int64
		End        int64
		Previous   int64
		Next       int64
		Total      int
		TestOutput string
		Query      string
	}
)

func (s *Server) HandleTrace(w http.ResponseWriter, r *http.Request) {
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

	t, err := loadTemplate("trace.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
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
		TestOutput: s.gb.CPU().MMU.TestOutput(),
		Start:      offset,
		End:        offset + int64(len(trace)),
		Previous:   offset - 1000,
		Next:       offset + 1000,
		Total:      len(s.trace),
	}

	for index, t := range trace {
		data.Trace = append(data.Trace, traceData{
			Index: offset + int64(index),
			Trace: t,
		})
	}

	err = t.ExecuteTemplate(w, "trace", data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *Server) HandleSearchTrace(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.FormValue("q")

	t, err := loadTemplate("trace.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var (
		found   = make(map[int64]traceEntry)
		indices []int64
	)
	for index, t := range s.trace {
		if strings.Contains(t.Description, searchTerm) {
			found[int64(index)] = t
			indices = append(indices, int64(index))
		}
	}

	data := searchData{
		Total: len(found),
		Query: searchTerm,
	}

	for _, index := range indices {
		data.Trace = append(data.Trace, traceData{
			Index: index,
			Trace: found[index],
		})
	}

	err = t.ExecuteTemplate(w, "searchtrace", data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
