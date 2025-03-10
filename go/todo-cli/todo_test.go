package todo_test

import (
	"os"
	"testing"

	"todo"
)

func TestAdd(t *testing.T) {
	list := todo.List{}

	taskName := "New Task"
	list.Add(taskName)

	if list[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, list[0].Task)
	}
}

func TestComplete(t *testing.T) {
	list := todo.List{}

	taskName := "New Task"
	list.Add(taskName)

	if list[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, list[0].Task)
	}

	if list[0].Done {
		t.Errorf("New task should not be completed.")
	}

	list.Complete(1)
	if !list[0].Done {
		t.Errorf("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	list := todo.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}
	for _, task := range tasks {
		list.Add(task)
	}

	if list[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead.", tasks[0], list[0].Task)
	}

	list.Delete(2)
	if len(list) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(list))
	}

	if list[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], list[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	list1 := todo.List{}
	list2 := todo.List{}

	taskName := "New Task"
	list1.Add(taskName)

	if list1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, list1[0].Task)
	}

	file, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(file.Name())

	if err := list1.Save(file.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := list2.Get(file.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if list1[0].Task != list2[0].Task {
		t.Errorf("Task %q should match %q task", list1[0].Task, list2[0].Task)
	}
}
