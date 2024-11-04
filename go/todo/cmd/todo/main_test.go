package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"todo"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool....")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up....")
	os.Remove(binName)
	os.Remove(fileName)
	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "test task 1"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("Add new task", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-task", task)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("List tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		// taking into account the indentation
		expected := fmt.Sprintf("  1: %s\n", task)
		if expected != string(output) {
			t.Errorf("Expected %q, got %q instead.", expected, string(output))
		}
	})

	t.Run("Mark task complete", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "1")

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		list := todo.List{}
		if err := list.Get(fileName); err != nil {
			t.Fatal(err)
		}

		if list[0].Done == false {
			t.Fatalf("Expected %v, got %v instead.", true, false)
		}
	})

	t.Run("Delete task", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-delete", "1")

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		list := todo.List{}
		if err := list.Get(fileName); err != nil {
			t.Fatal(err)
		}

		if len(list) != 0 {
			t.Fatalf("Expected %d, got %d instead.", 0, len(list))
		}
	})
}
