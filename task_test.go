package minion

import "testing"

func TestTaskCreation(t *testing.T) {
	task := NewTask("Get stuff done")

	desc := task.Get("description")
	if desc.(string) != "Get stuff done" {
		t.Errorf("Task created with description %s contained the modified description %s right after creation", task, desc)
	}
}
