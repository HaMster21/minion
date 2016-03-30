// This file is part of the minion task manager
// Copyright (C) 2016  Hans Meyer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

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
