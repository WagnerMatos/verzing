package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Versioner defines the operations needed for version handling.
type Versioner interface {
	ReadVersion() (string, error)
	WriteVersion(version string) error
}

// FileVersioner implements Versioner with actual file operations.
type FileVersioner struct{}

func (fv FileVersioner) ReadVersion() (string, error) {
	// Simulate reading version from a file
	return "0.1.0", nil
}

func (fv FileVersioner) WriteVersion(version string) error {
	// Simulate writing version to a file
	return nil
}

// updateVersion reads, updates, and writes back the version.
func updateVersion(v Versioner, commitType string, breakingChange bool) (string, error) {
	version, err := v.ReadVersion()
	if err != nil {
		return "", fmt.Errorf("failed to read version: %w", err)
	}

	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid version format")
	}

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	if breakingChange {
		major++
		minor = 0
		patch = 0
	} else if commitType == "Fix" {
		patch++
	} else {
		minor++
		patch = 0
	}

	newVersion := fmt.Sprintf("%d.%d.%d", major, minor, patch)
	err = v.WriteVersion(newVersion)
	if err != nil {
		return "", fmt.Errorf("failed to write version: %w", err)
	}

	return newVersion, nil
}
