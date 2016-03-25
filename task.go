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
