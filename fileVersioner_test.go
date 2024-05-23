package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileVersioner_ReadVersion(t *testing.T) {
	// Create a temporary file
	tempFile, err := ioutil.TempFile("", "VERSION")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up after the test

	// Write a version to the temp file
	versionContent := "1.0.2"
	if _, err := tempFile.WriteString(versionContent); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}
	tempFile.Close() // Close the file to allow reading from it later

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
	// Setup: Create a temporary file to act as VERSION file
	tempFile, err := ioutil.TempFile("", "VERSION")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up after the test

	// Temporarily replace the file path used in FileVersioner
	oldName := "VERSION"                      // Assume this is the constant or variable used.
	os.Rename(tempFile.Name(), oldName)       // This is just for demonstration. Typically, you would inject the file path.
	defer os.Rename(oldName, tempFile.Name()) // Restore the original state after test

	// Test writing the version
	fv := FileVersioner{}
	newVersion := "2.0.1"
	if err := fv.WriteVersion(newVersion); err != nil {
		t.Errorf("WriteVersion() error = %v", err)
	}

	// Read back the content to verify it's correct
	content, err := ioutil.ReadFile(oldName)
	if err != nil {
		t.Fatalf("Failed to read temp file: %s", err)
	}
	if string(content) != newVersion {
		t.Errorf("WriteVersion() failed, got = %v, want %v", string(content), newVersion)
	}
}
