package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	t := template.Must(template.New("opcodes").Parse(opcodes))

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var opcodes map[string]map[string]Opcode
	err = json.Unmarshal(data, &opcodes)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, opcodes)
	if err != nil {
		fmt.Println(err)
	}

}

var opcodes = `
package gameboy

{{range $prefixType, $opcodes := . -}}
func {{$prefixType}}Opcodes(c *CPU) map[Opcode]Op {
	return map[Opcode]Op{
		{{- range $opcodes}}
			{{.Addr}}: NewOp(c.{{.Mnemonic}}, []int{
				{{- range .Cycles}}
					{{.}},
				{{end -}}
				},
				{{- if .Operand1}}
				{{.Operand1}},
				{{- end -}}
				{{- if .Operand2}}
				{{.Operand2}},
				{{- end}}
			),
		{{- end}}
	}
}
{{end}}
`
