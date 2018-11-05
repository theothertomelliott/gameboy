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
func {{$prefixType}}Handler(c *CPU, code Opcode) (string, []int) {
	switch code {
		{{- range $opcodes}}
		case {{.Addr}}:
			c.{{.Mnemonic}}(
				{{- if .Operand1}}
				{{.Operand1}},
				{{- end -}}
				{{- if .Operand2}}
				{{.Operand2}},
				{{- end}}
			)
			return "{{.Description}}", []int{
				{{- range .Cycles}}
					{{.}},
				{{- end -}}
				}
		{{- end}}
		default:
			panic(fmt.Sprintf("unknown opcode: 0x%X", code))
	}
}
{{end}}
`
