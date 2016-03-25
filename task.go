package minion

import (
	"bytes"
	"text/template"
)

var (
	currentPattern = "{{.description}}"
	currentTemplate *template.Template = nil
)

func init() {
	// get user-defined pattern from the config
	tpl, err := template.New("task").Parse(currentPattern)
	if err != nil {
		panic (err)
	}
	currentTemplate = tpl
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

func (t *Task) TplString(templ string) string {
	var tpl *template.Template
	var err error
	var buf bytes.Buffer

	tpl, err = template.New("task").Parse(templ + " {{ .description }}")
	if err != nil {
		tpl, err = template.New("task").Parse(currentPattern)
		if err != nil {
			panic (err)
		}
	}

	err = tpl.Execute(&buf, map[string]interface{}(*t))
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *Task) String() string {
	var buf bytes.Buffer

	if err := currentTemplate.Execute(&buf, map[string]interface{}(*t)); err != nil {
		panic (err)
	}

	return buf.String()
}
