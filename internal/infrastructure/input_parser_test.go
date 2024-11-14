package infrastructure_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/abakunov/mazes/internal/infrastructure"
)

// Helper function for setting up standard input.
func mockStdin(input string) func() {
	origStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Check the error from WriteString
	if _, err := w.WriteString(input); err != nil {
		fmt.Println("Error writing to stdin pipe:", err)
	}

	err := w.Close()
	if err != nil {
		return nil
	}

	return func() {
		os.Stdin = origStdin
	}
}

// Helper function for capturing standard output.
func captureStdout(f func()) string {
	var buf bytes.Buffer

	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	done := make(chan struct{})
	go func() {
		_, _ = buf.ReadFrom(r)

		close(done)
	}()

	f()

	_ = w.Close()
	os.Stdout = origStdout

	<-done

	return buf.String()
}

// Tests for width and height

func TestGetWidth_InvalidEvenInput(t *testing.T) {
	mockInput := "4\n5\n" // 4 - even, then 5 - odd
	restoreStdin := mockStdin(mockInput)

	defer restoreStdin()

	output := captureStdout(func() {
		width := infrastructure.GetWidth()
		if width != 5 {
			t.Errorf("Expected width to be 5, got %d", width)
		}
	})

	if !strings.Contains(output, "Принимаются только нечетные числа") {
		t.Error("Expected output to contain 'Принимаются только нечетные числа'")
	}
}

func TestGetWidth_NonNumericInput(t *testing.T) {
	mockInput := "abc\n5\n" // "abc" - invalid value, then 5 - valid

	restoreStdin := mockStdin(mockInput)

	defer restoreStdin()

	output := captureStdout(func() {
		width := infrastructure.GetWidth()
		if width != 5 {
			t.Errorf("Expected width to be 5, got %d", width)
		}
	})

	if !strings.Contains(output, "Введите корректное нечетное число") {
		t.Error("Expected output to contain 'Введите корректное нечетное число'")
	}
}

func TestGetWidth_EmptyInput(t *testing.T) {
	mockInput := "\n5\n" // empty input, then 5 - valid value
	restoreStdin := mockStdin(mockInput)

	defer restoreStdin()

	output := captureStdout(func() {
		width := infrastructure.GetWidth()
		if width != 5 {
			t.Errorf("Expected width to be 5, got %d", width)
		}
	})

	if !strings.Contains(output, "Введите корректное нечетное число") {
		t.Error("Expected output to contain 'Введите корректное нечетное число'")
	}
}

func TestGetHeight_NonNumericInput(t *testing.T) {
	mockInput := "height\n7\n" // "height" - invalid value, then 7 - valid
	restoreStdin := mockStdin(mockInput)

	defer restoreStdin()

	output := captureStdout(func() {
		height := infrastructure.GetHeight()
		if height != 7 {
			t.Errorf("Expected height to be 7, got %d", height)
		}
	})

	if !strings.Contains(output, "Введите корректное нечетное число") {
		t.Error("Expected output to contain 'Введите корректное нечетное число'")
	}
}

func TestGetHeight_EmptyInput(t *testing.T) {
	mockInput := "\n7\n" // empty input, then 7 - valid value
	restoreStdin := mockStdin(mockInput)

	defer restoreStdin()

	output := captureStdout(func() {
		height := infrastructure.GetHeight()
		if height != 7 {
			t.Errorf("Expected height to be 7, got %d", height)
		}
	})

	if !strings.Contains(output, "Введите корректное нечетное число") {
		t.Error("Expected output to contain 'Введите корректное нечетное число'")
	}
}

// Tests for maze generation algorithm choice

func TestGetAlgorithmChoice_InvalidStringInput(t *testing.T) {
	mockInput := "invalid\n2\n" // invalid string, then 2 - valid value
	restoreStdin := mockStdin(mockInput)

	defer restoreStdin()

	output := captureStdout(func() {
		choice := infrastructure.GetAlgorithmChoice()
		if choice != 2 {
			t.Errorf("Expected algorithm choice to be 2, got %d", choice)
		}
	})

	if !strings.Contains(output, "Ошибка: выберите 1 (DFS) или 2 (Kruskal)") {
		t.Error("Expected output to contain 'Ошибка: выберите 1 (DFS) или 2 (Kruskal)'")
	}
}

func TestGetAlgorithmChoice_EmptyInput(t *testing.T) {
	mockInput := "\n1\n" // empty input, then 1 - valid value
	restoreStdin := mockStdin(mockInput)

	defer restoreStdin()

	output := captureStdout(func() {
		choice := infrastructure.GetAlgorithmChoice()
		if choice != 1 {
			t.Errorf("Expected algorithm choice to be 1, got %d", choice)
		}
	})

	if !strings.Contains(output, "Ошибка: выберите 1 (DFS) или 2 (Kruskal)") {
		t.Error("Expected output to contain 'Ошибка: выберите 1 (DFS) или 2 (Kruskal)'")
	}
}

// Tests for entry and exit point choice

func TestGetEntryExitChoice_InvalidNegativeInput(t *testing.T) {
	mockInput := "-1\n2\n" // -1 - invalid value, then 2 - valid value
	restoreStdin := mockStdin(mockInput)

	defer restoreStdin()

	output := captureStdout(func() {
		choice := infrastructure.GetEntryExitChoice()
		if choice != 2 {
			t.Errorf("Expected entry/exit choice to be 2, got %d", choice)
		}
	})

	if !strings.Contains(output, "Ошибка: выберите 1 (вручную) или 2 (случайным образом)") {
		t.Error("Expected output to contain 'Ошибка: выберите 1 (вручную) или 2 (случайным образом)'")
	}
}

// Tests for pathfinding algorithm choice

func TestGetPathSolverChoice_InvalidStringInput(t *testing.T) {
	mockInput := "astar\n1\n" // "astar" - invalid value, then 1 - valid value
	restoreStdin := mockStdin(mockInput)

	defer restoreStdin()

	output := captureStdout(func() {
		choice := infrastructure.GetPathSolverChoice()
		if choice != 1 {
			t.Errorf("Expected path solver choice to be 1, got %d", choice)
		}
	})

	if !strings.Contains(output, "Ошибка: выберите 1 (BFS) или 2 (A*)") {
		t.Error("Expected output to contain 'Ошибка: выберите 1 (BFS) или 2 (A*)'")
	}
}

func TestGetPathSolverChoice_EmptyInput(t *testing.T) {
	mockInput := "\n2\n" // empty input, then 2 - valid value
	restoreStdin := mockStdin(mockInput)

	defer restoreStdin()

	output := captureStdout(func() {
		choice := infrastructure.GetPathSolverChoice()
		if choice != 2 {
			t.Errorf("Expected path solver choice to be 2, got %d", choice)
		}
	})

	if !strings.Contains(output, "Ошибка: выберите 1 (BFS) или 2 (A*)") {
		t.Error("Expected output to contain 'Ошибка: выберите 1 (BFS) или 2 (A*)'")
	}
}
