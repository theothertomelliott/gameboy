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

import "fmt"

{{range $prefixType, $opcodes := . -}}
func {{$prefixType}}Handler(c *CPU, code Opcode) (string, []int, error) {
	switch code {
		{{- range $opcodes}}
		case {{.Addr}}:
			{{- if .Operand1}}
				o1 := {{.Operand1}}
			{{- end -}}
			{{- if .Operand2}}
				o2 := {{.Operand2}}
			{{- end}}
			c.{{.Mnemonic}}(
				{{- if .Operand1}}
				o1,
				{{- end -}}
				{{- if .Operand2}}
				o2,
				{{- end}}
			)
			description := fmt.Sprint(
				"{{.Description}} ",
				{{- if .Operand1}}
				o1,
				{{- end -}}
				{{- if .Operand2}}
				o2,
				{{- end}}
			)
			return description, []int{
				{{- range .Cycles}}
					{{.}},
				{{- end -}}
				}, nil
		{{- end}}
		default:
			return "", nil, fmt.Errorf("unknown opcode: 0x%X", code)
	}
}
{{end}}
`
