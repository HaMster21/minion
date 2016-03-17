package task

import (
	"fmt"
	dom "github.com/stretchr/objx"
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
