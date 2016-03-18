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


package task

import (
	"bytes"
	"fmt"
	"text/template"
	dom "github.com/stretchr/objx"
)

var (
	taskBlueprint = `{{ .description }}`
	taskPrinter = template.Must(template.New("Task").Parse(taskBlueprint))
)

type Task struct {
	model dom.Map
}

func New() *Task {
	return &Task{dom.MSI("description", "##Edit me!##")}
}

func FromJSON(json string) (*Task, error) {
	model, err := dom.FromJSON(json)
	if err != nil {
		return nil, err
	}

	if model.Get("description").IsNil() {
		return nil, fmt.Errorf("A task must have a description. %v has no toplevel \"description\" key", model)
	}

	return &Task{model}, err
}

func (t *Task) Get(dompath string) interface{} {
	return t.model.Get(dompath).Data()
}

func (t *Task) Set(dompath string, value interface{}) {
	t.model = t.model.Set(dompath, value)
}

func (t *Task) String() string {
	return t.displayWithTemplate()
}

func (t *Task) GoString() string {
	return fmt.Sprintf("%#v", t.model)
}

func (t *Task) displayWithTemplate() string {
	var result bytes.Buffer
	if err := taskPrinter.Execute(&result, t.model); err != nil {
		panic (err)
	}

	return result.String()
}