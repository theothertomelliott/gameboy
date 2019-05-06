package httpui

import (
	"net/http"
	"strconv"
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

	var trace []string
	if int(offset) < len(s.trace) {
		if (len(s.trace) - int(offset)) < pageSize {
			trace = s.trace[offset:]
		} else {
			trace = s.trace[offset : offset+pageSize]
		}
	}

	data := struct {
		Trace []struct {
			Index       int64
			Description string
		}
		Start      int64
		End        int64
		Previous   int64
		Next       int64
		Total      int
		TestOutput string
	}{
		TestOutput: s.gb.CPU().MMU.TestOutput(),
		Start:      offset,
		End:        offset + int64(len(trace)),
		Previous:   offset - 1000,
		Next:       offset + 1000,
		Total:      len(s.trace),
	}

	for index, description := range trace {
		data.Trace = append(data.Trace, struct {
			Index       int64
			Description string
		}{
			Index:       offset + int64(index),
			Description: description,
		})
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
