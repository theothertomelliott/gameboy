package httpui

import (
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) HandleToggleBreakpoint(w http.ResponseWriter, r *http.Request) {
	posStr := strings.Replace(r.URL.Path, "/debug/togglebreakpoint/", "", 1)
	pos, err := strconv.ParseInt(posStr, 16, 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if _, exists := s.gb.Breakpoints[uint16(pos)]; exists {
		delete(s.gb.Breakpoints, uint16(pos))
	} else {
		s.gb.Breakpoints[uint16(pos)] = struct{}{}
	}

	http.Redirect(w, r, "/debug", 302)
}
