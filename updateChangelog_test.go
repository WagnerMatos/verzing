package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestUpdateChangelog(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "CHANGELOG.md")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	tempFilePath := tempFile.Name()
	tempFile.Close() // Close the file as updateChangelog will open it again

	// Redirect CHANGELOG.md path to the temp file
	os.Rename(tempFilePath, "CHANGELOG.md")
	defer os.Remove("CHANGELOG.md") // Clean up after the test

	// Test data
	version := "v1.0.1"
	shortDesc := "Added new feature"
	longDesc := "This feature improves performance."

	// Call the function under test
	err = updateChangelog(version, shortDesc, longDesc)
	if err != nil {
		t.Errorf("updateChangelog failed: %s", err)
	}

	// Read back the contents of the CHANGELOG.md to verify the write
	content, err := ioutil.ReadFile("CHANGELOG.md")
	if err != nil {
		t.Fatalf("Failed to read back the CHANGELOG.md: %s", err)
	}

	expectedContent := "## v1.0.1\n- Added new feature\n  This feature improves performance.\n"
	if string(content) != expectedContent {
		t.Errorf("updateChangelog wrote unexpected content: got %v, want %v", string(content), expectedContent)
	}
}
