package httpui

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
)

func (s *Server) HandleDecompile(w http.ResponseWriter, r *http.Request) {
	s.decompileMtx.Lock()
	defer s.decompileMtx.Unlock()

	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Gameboy - Decompile</title>
			<style>
				body {font-family: "Courier New", Courier, serif;}
			</style>
		</head>
		<body>
			<table border="0">
			{{range .Op}}
				<tr>
				<td>{{ .Index }}</td>
				<td>&nbsp;</td>
				<td>{{ .Description }}</td>
				</tr>
			{{end}}
			</table>
		</body>
	</html>`

	t, err := template.New("index").Parse(tpl)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	type (
		row struct {
			Index       string
			Description string
		}
		table struct {
			Op []row
		}
	)

	var indices []uint16
	for index := range s.decompilation {
		indices = append(indices, index)
	}
	sort.Slice(indices, func(i, j int) bool {
		return indices[i] < indices[j]
	})

	data := table{}
	for _, index := range indices {
		r := row{}
		r.Index = fmt.Sprintf("%04X", index)
		r.Description = s.decompilation[index]
		data.Op = append(data.Op, r)
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
