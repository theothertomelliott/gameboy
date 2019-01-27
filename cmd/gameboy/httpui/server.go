package httpui

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"github.com/theothertomelliott/gameboy"
)

type Server struct {
	gb *gameboy.DMG

	decompilation map[uint16]string
	decompileMtx  sync.Mutex
}

func NewServer(gb *gameboy.DMG) *Server {
	return &Server{
		gb:            gb,
		decompilation: make(map[uint16]string),
	}
}

func (s *Server) Trace(ev gameboy.TraceMessage) {
	s.decompileMtx.Lock()
	s.decompilation[ev.CPU.PC] = ev.CPU.Description
	s.decompileMtx.Unlock()
}

// ListenAndServe starts a UI server on the specified port
func (s *Server) ListenAndServe(port int) error {
	http.HandleFunc("/memory", s.HandleMemory)
	http.HandleFunc("/decompile", s.HandleDecompile)
	http.HandleFunc("/", s.HandleIndex)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Gameboy</title>
		</head>
		<body>
			{{range .Items}}
				<a href="{{ . }}">{{ . }}</a><br/>
			{{end}}
		</body>
	</html>`

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
			"/decompile",
		},
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
