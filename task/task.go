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

// Package task defines a basic Task object and functions to manipulate and
// print it
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

/*
A Task is a map[string]interface{} that can be manipulated through a DOM. It
stores arbitrary data, however nested. The only requirement to define a valid
task is a toplevel "description" key of type string.
*/
type Task struct {
	model dom.Map
}

// New creates a Task with an empty description
func New() *Task {
	return &Task{dom.MSI("description", "")}
}

// FromJSON assumes the given string to contain valid JSON and parses it into a
// Task object. It returns a valid Task and nil on success. An error is returned
// if the JSON data was invalid or there was no description in the data.
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

/*
Get retrieves values from the Task by following the given DOM path. This also
works for arrays.

Example:

myTask.Get("audit.changes[2].time")
*/
func (t *Task) Get(dompath string) interface{} {
	return t.model.Get(dompath).Data()
}

/*
Set changes the value of the key specified by dompath to the given value,
assuming the key to exist already. Intermediate keys have to be created first.

Example:

myTask.Set("description", "Find the question for 42")
// Okay if you called Set.("my.nested", []interface{}{14,1,72,412}) or similar before
myTask.Set("my.nested.array[4]", 1337)
*/
func (t *Task) Set(dompath string, value interface{}) {
	t.model = t.model.Set(dompath, value)
}

// String returns a string representation of the Task by applying a template to
// it. The template defaults to {{ .description }}, returning only the tasks
// description
func (t *Task) String() string {
	return t.displayWithTemplate()
}

// GoString returns the Task object as %#v formatted representation of the
// underlying map[string]interface{}
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