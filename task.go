package minion

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
