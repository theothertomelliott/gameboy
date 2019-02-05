package httpui

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
)

// HandleDebug renders the debugger view
func (s *Server) HandleDebug(w http.ResponseWriter, r *http.Request) {
	s.decompileMtx.Lock()
	defer s.decompileMtx.Unlock()

	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Gameboy - Debug</title>
			<style>
				body {font-family: "Courier New", Courier, serif;}

				.debugmenu {
					overflow: hidden;
					width: 100%;
					height: 50px;
					position: fixed; /* Set the navbar to fixed position */
					top: 0; /* Position the navbar at the top of the page */
					border-bottom: 1px solid #000;
					left: 0;
					background-color: #c0c0c0;
				}

				.debugmenu ul li {
					float: left;
					margin-left: 20px;
					list-style: none;
				}

				.debugspacer {
					width: 100%;
					height: 50px;
				}

				.registers {
					display: block;
					position: fixed; /* Set the navbar to fixed position */
					top: 50px; /* Position the navbar at the top of the page */
					right: 0;
					margin-right: 10px;
				}

				#PC {
					background-color: #fff2a8;
				}

				.oprow {
					cursor: pointer;
				}
			</style>
			<script language="javascript">
			function scrollToPC() {
				var pc = document.getElementById("PC");
				var targetPos = pc.offsetTop - document.getElementsByClassName('debugspacer')[0].offsetHeight;
				if ('scrollRestoration' in window.history) {
					window.history.scrollRestoration = 'manual'
				}
				window.scrollTo(0, targetPos);
			}
			</script>
		</head>
		<body onload="scrollToPC();">
			<div class="debugmenu">
				<!--<h3>Debug</h3>-->
				<ul>
					<li><a href="/debug/togglepaused">{{if .Paused}}Resume{{else}}Pause{{end}}</a></li>
					<li><a href="/debug/step">Step</a></li>
				</ul>
			</div>
			<div class="debugspacer">
			</div>
			<div class="registers">
				<h2>Registers</h2>
				<table border="0">
					<tr><td>A</td><td>{{.Registers.A}}</td><td>F</td><td>{{.Registers.F}}</td></tr>
					<tr><td>B</td><td>{{.Registers.B}}</td><td>C</td><td>{{.Registers.C}}</td></tr>
					<tr><td>D</td><td>{{.Registers.D}}</td><td>E</td><td>{{.Registers.E}}</td></tr>
					<tr><td>H</td><td>{{.Registers.H}}</td><td>F</td><td>{{.Registers.L}}</td></tr>
				</table>
			</div>
			<table border="0" cellpadding="0" cellspacing="0">
			{{range .Op}}
				<tr id="{{.Id}}" onclick="location.href='/breakpoints/toggle/{{ .Index }}';" class="oprow">
				{{if .Breakpoint}}
					<td>•</td>
				{{else}}
					<td>&nbsp;</td>
				{{end}}
				<td>
					{{range .Flags}}
						<a id="{{.}}"></a>
					{{end}}
					{{ .Index }}
				</td>
				<td>&nbsp;</td>
				<td>{{ .Description }}</td>
				</tr>
			{{end}}
			</table>
		</body>
	</html>`

	t, err := template.New("decompile").Parse(tpl)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var indices []uint16
	for index := range s.decompilation {
		indices = append(indices, index)
	}
	sort.Slice(indices, func(i, j int) bool {
		return indices[i] < indices[j]
	})

	data := table{
		Registers: s.getRegisters(),
		Paused:    s.gb.IsPaused(),
	}
	for _, index := range indices {
		r := row{}
		r.Index = fmt.Sprintf("%04X", index)
		r.Description = s.decompilation[index]
		r.Flags = []string{r.Index}
		r.Id = r.Index
		if index == s.gb.CPU().PC.Read16() {
			r.Id = "PC"
			r.Flags = append(r.Flags, "PC")
		}
		data.Op = append(data.Op, r)
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// HandleTogglePaused toggles the paused state of the DMB emulator and redirects to the debugger
func (s *Server) HandleTogglePaused(w http.ResponseWriter, r *http.Request) {
	s.gb.SetPaused(!s.gb.IsPaused())
	http.Redirect(w, r, "/debug", 302)
}

// HandleStep steps the DMB and redirects to the debugger
func (s *Server) HandleStep(w http.ResponseWriter, r *http.Request) {
	err := s.gb.Step()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/debug", 302)
}

func (s *Server) getRegisters() registers {
	cpu := s.gb.CPU()
	return registers{
		A: fmt.Sprintf("0x%02X", cpu.A.Read8()),
		F: fmt.Sprintf("0x%02X", cpu.F.Read8()),
		B: fmt.Sprintf("0x%02X", cpu.B.Read8()),
		C: fmt.Sprintf("0x%02X", cpu.C.Read8()),
		D: fmt.Sprintf("0x%02X", cpu.D.Read8()),
		E: fmt.Sprintf("0x%02X", cpu.E.Read8()),
		H: fmt.Sprintf("0x%02X", cpu.H.Read8()),
		L: fmt.Sprintf("0x%02X", cpu.L.Read8()),
	}
}

type (
	registers struct {
		A, F string
		B, C string
		D, E string
		H, L string
	}
	row struct {
		Index       string
		Description string
		Flags       []string
		Id          string
		Breakpoint  bool
	}
	table struct {
		Registers registers
		Op        []row
		Paused    bool
	}
)
