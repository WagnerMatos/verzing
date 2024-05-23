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
//type FileVersioner struct{}
//
//// ReadVersion reads the version from the VERSION.md file.
//func (fv FileVersioner) ReadVersion() (string, error) {
//	file, err := os.Open("VERSION.md")
//	if err != nil {
//		return "", fmt.Errorf("error opening VERSION.md file: %w", err)
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//	if scanner.Scan() {
//		return strings.TrimSpace(scanner.Text()), nil
//	}
//
//	if err := scanner.Err(); err != nil {
//		return "", fmt.Errorf("error reading VERSION.md file: %w", err)
//	}
//
//	return "", fmt.Errorf("VERSION.md file is empty")
//}
//
//// WriteVersion writes the updated version to the VERSION.md file.
//func (fv FileVersioner) WriteVersion(version string) error {
//	file, err := os.OpenFile("VERSION.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
//	if err != nil {
//		return fmt.Errorf("error opening VERSION.md file for write: %w", err)
//	}
//	defer file.Close()
//
//	_, err = file.WriteString(version + "\n")
//	if err != nil {
//		return fmt.Errorf("error writing to VERSION.md file: %w", err)
//	}
//
//	return nil
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
