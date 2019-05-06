package httpui

import (
	"html/template"

	packr "github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
)

func loadTemplate(file string) (*template.Template, error) {
	box := packr.New("views", "./views")
	tpl, err := box.FindString(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return template.New(file).Parse(tpl)
}
