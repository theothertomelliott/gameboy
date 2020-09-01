package httpui

import (
	"fmt"
	"net/http"

	"github.com/theothertomelliott/gameboy/mmu"
)

// HandleMemory displays a Hex Editor like view of the emulator's memory
func (s *Server) HandleMemory(w http.ResponseWriter, r *http.Request) {
	t, err := loadTemplate("memory.html")
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
	allMem := mem.ReadRange(mmu.Range{
		Start: 0x0000,
		End:   0xFFFE,
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
