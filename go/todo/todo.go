package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string    `json:"task"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

type List []item

// Marks a todo item as completed.
func (l *List) Complete(idx int) error {
	ls := *l

	if idx <= 0 || idx > len(ls) {
		return fmt.Errorf("Item %d does not exists", idx)
	}

	// adjust index for 0 based index counting...
	ls[idx-1].Done = true
	ls[idx-1].CompletedAt = time.Now()

	return nil
}

// Saves the list of items to a file using the JSON format.
func (l *List) Save(filename string) error {
	jsonData, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

// Creates a new todo item and appends it to the list.
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Deletes a todo item from the list
func (l *List) Delete(idx int) error {
	ls := *l

	if idx <= 0 || idx > len(ls) {
		return fmt.Errorf("Item %d does not exists", idx)
	}

	*l = append(ls[:idx-1], ls[idx:]...)

	return nil
}

// Gets a list of items from a saved JSON file
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formatted := ""

	for idx, item := range *l {
		prefix := "  "

		if item.Done {
			prefix = "X "
		}

		formatted += fmt.Sprintf("%s%d: %s\n", prefix, idx+1, item.Task)
	}

	return formatted
}
