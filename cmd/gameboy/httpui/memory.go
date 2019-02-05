package httpui

import (
	"fmt"
	"html/template"
	"net/http"

	packr "github.com/gobuffalo/packr/v2"
	"github.com/theothertomelliott/gameboy"
)

// HandleMemory displays a Hex Editor like view of the emulator's memory
func (s *Server) HandleMemory(w http.ResponseWriter, r *http.Request) {
	box := packr.New("views", "./views")
	tpl, err := box.FindString("memory.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t, err := template.New("memory.html").Parse(tpl)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	type (
		row struct {
			Offset string
			Hex    []string
			Text   []string
		}
		table struct {
			Mem []row
		}
	)

	mem := s.gb.MMU()
	allMem := mem.ReadRange(gameboy.Range{
		Start: 0x0000,
		End:   0xFFFF,
	})

	data := table{}
	for i := 0; i < len(allMem); i += 0x10 {
		r := row{}
		for j := 0; j < 0x10; j++ {
			pos := i + j
			if len(allMem) <= pos {
				break
			}
			v := allMem[i+j]
			r.Offset = fmt.Sprintf("%04X", i)
			r.Hex = append(r.Hex, fmt.Sprintf("%02X", v))
			r.Text = append(r.Text, fmt.Sprintf("%c", v))
		}
		data.Mem = append(data.Mem, r)
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
