package task

import (
	dom "github.com/stretchr/objx"
)

type Task struct {
	model dom.Map
}

func FromJSON(json string) (*Task, error) {
	model, err := dom.FromJSON(json)
	return &Task{model}, err
}

func (t *Task) Get(dompath string) interface{} {
	return t.model.Get(dompath).Data()
}

func (t *Task) Set(dompath string, value interface{}) {
	t.model = t.model.Set(dompath, value)
}
