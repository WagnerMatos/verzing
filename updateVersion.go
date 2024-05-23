package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Versioner defines the operations needed for version handling.
//type Versioner interface {
//	ReadVersion() (string, error)
//	WriteVersion(version string) error
//}

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
