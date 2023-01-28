package src

import (
	"os"
	"testing"
)

// Test ReadFileLineByLine
func TestReadFileLineByLine(t *testing.T) {
	// Create a temporary file
	file, err := os.CreateTemp(t.TempDir(), "test.txt")
	if err != nil {
		t.Errorf("Could not create temporary file: %s", err)
	}
	defer os.Remove(file.Name())
	// Write some dummy data to the file
	file.WriteString("Line 1\nLine 2\nLine 3")

	lines, err := ReadFileLineByLine(file.Name())
	if err != nil {
		t.Errorf("Could not read file: %s", err)
	}
	if len(lines) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(lines))
	}
	if lines[0] != "Line 1" {
		t.Errorf("Expected 'Line 1', got '%s'", lines[0])
	}
	if lines[1] != "Line 2" {
		t.Errorf("Expected 'Line 2', got '%s'", lines[1])
	}
	if lines[2] != "Line 3" {
		t.Errorf("Expected 'Line 3', got '%s'", lines[1])
	}
}

// Test ReadFileLineByLine with a non-existing file
func TestReadFileLineByLineNonExistingFile(t *testing.T) {
	_, err := ReadFileLineByLine("non-existing-file")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestReadFilesConfigYaml(t *testing.T) {
	// Write temporary test file
	file, err := os.CreateTemp(t.TempDir(), "test.yaml")
	if err != nil {
		t.Errorf("Could not create temporary file: %s", err)
	}
	defer os.Remove(file.Name())
	file.WriteString(`---
- filename: file1
  filetype: yaml
- filename: file2
  filetype: json
`)

	files, err := ReadFilesConfigYaml(file.Name())
	if err != nil {
		t.Errorf("Could not read file: %s", err)
	}
	if len(files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(files))
	}
	if files[0].Filename != "file1" {
		t.Errorf("Expected 'README.md', got '%s'", files[0].Filename)
	}
	if files[0].Filetype != "yaml" {
		t.Errorf("Expected 'yaml', got '%s'", files[0].Filetype)
	}
	if files[1].Filename != "file2" {
		t.Errorf("Expected 'file2', got '%s'", files[1].Filename)
	}
	if files[1].Filetype != "json" {
		t.Errorf("Expected 'json', got '%s'", files[1].Filetype)
	}
}
