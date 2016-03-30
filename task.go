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

package minion

import (
	"github.com/hamster21/minion/info"
)

var (
	printer *info.Printer
)

func init() {
	// get task printing pattern from config
	p, err := info.NewPrinter("{{ .description }}")
	if err != nil {
		panic (err)
	}
	printer = p
}

type Task map[string]interface{}

func NewTask(description string) Task {
	t := make(map[string]interface{})
	t["description"] = description
	return t
}

func (t *Task) Get(at string) interface{} {
	return map[string]interface{}(*t)[at]
}

func (t *Task) Set(at string, to interface{}) error {
	map[string]interface{}(*t)[at] = to
	return nil
}

func (t *Task) String() string {
	return printer.Print(map[string]interface{}(*t))
}
