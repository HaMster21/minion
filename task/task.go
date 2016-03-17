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