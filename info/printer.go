package info

import (
	"bytes"
	"fmt"
	"text/template"
)

type Printer struct {
	Pattern string
	tpl     *template.Template
}

func NewPrinter(pattern string) (*Printer, error) {
	t, err := template.New("Printer").Parse(pattern)
	if err != nil {
		return nil, fmt.Errorf("Cannot create a new Printer: %s", err)
	}
	return &Printer{pattern, t}, nil
}

func (p *Printer) Print(data map[string]interface{}) string {
	if p.tpl == nil {
		t, err := template.New("Printer").Parse(p.Pattern)
		if err != nil {
			err = fmt.Errorf("Unable to print with pattern %s: %s", p.Pattern, err)
			panic(err)
		}
		p.tpl = t
	}

	var buf bytes.Buffer
	if err := p.tpl.Execute(&buf, data); err != nil {
		err = fmt.Errorf("Unable to template on %s: %s", data, err)
		panic(err)
	}

	return buf.String()
}
