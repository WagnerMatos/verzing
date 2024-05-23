package main

import (
	"os"
	"testing"
)

func TestFileVersioner_ReadVersion(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "VERSION")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer func() {
		cerr := os.Remove(tempFile.Name()) // Clean up after the test
		if err == nil {                    // only overwrite err if it's still nil
			err = cerr
		}
	}()

	// Write a version to the temp file
	versionContent := "1.0.2"
	if _, err := tempFile.WriteString(versionContent); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}
	defer func() {
		cerr := tempFile.Close()
		if err == nil { // only overwrite err if it's still nil
			err = cerr
		}
	}()

	// Test reading the version
	fv := FileVersioner{FilePath: tempFile.Name()}
	got, err := fv.ReadVersion()
	if err != nil {
		t.Errorf("ReadVersion() error = %v", err)
	}
	if got != versionContent {
		t.Errorf("ReadVersion() = %v, want %v", got, versionContent)
	}
}

func TestFileVersioner_WriteVersion(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "VERSION")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	tempFilePath := tempFile.Name()
	defer func() {
		cerr := tempFile.Close()
		if err == nil { // only overwrite err if it's still nil
			err = cerr
		}
	}()
	defer func() {
		cerr := os.Remove(tempFilePath) // Clean up after the test
		if err == nil {                 // only overwrite err if it's still nil
			err = cerr
		}
	}()

	// Instance of FileVersioner with the path to the temp file
	fv := FileVersioner{FilePath: tempFilePath}

	// Version string to be written
	testVersion := "2.0.3"

	// Write the version to the temp file using WriteVersion
	if err := fv.WriteVersion(testVersion); err != nil {
		t.Errorf("WriteVersion() error = %v", err)
	}

	// Read back the contents of the file to verify it was written correctly
	content, err := os.ReadFile(tempFilePath)
	if err != nil {
		t.Fatalf("Failed to read back the temp file: %s", err)
	}

	// Check if the content of the file is exactly what we wrote
	if string(content) != testVersion+"\n" { // Remember WriteVersion adds a newline
		t.Errorf("WriteVersion() failed, content = %q, want %q", string(content), testVersion+"\n")
	}
}
