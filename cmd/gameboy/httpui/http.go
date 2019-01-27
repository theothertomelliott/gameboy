package httpui

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/theothertomelliott/gameboy"
)

type Server struct {
	gb *gameboy.DMG
}

// ListenAndServe starts a UI server on the specified port
func ListenAndServe(gb *gameboy.DMG, port int) error {
	s := &Server{
		gb: gb,
	}
	http.HandleFunc("/memory", s.HandleMemory)
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
				<a href="{{ . }}">{{ . }}</a>
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
		},
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
