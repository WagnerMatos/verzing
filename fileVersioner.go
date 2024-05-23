package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FileVersioner implements Versioner with actual file operations.
type FileVersioner struct{}

// ReadVersion reads the version from VERSION.md file.
func (fv FileVersioner) ReadVersion() (string, error) {
	file, err := os.Open("VERSION.md")
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading VERSION.md file: %v", err)
	}

	return "", fmt.Errorf("VERSION.md file is empty")
}

// WriteVersion writes the updated version to the VERSION.md file.
func (fv FileVersioner) WriteVersion(version string) error {
	file, err := os.OpenFile("VERSION.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening VERSION.md file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(version + "\n")
	if err != nil {
		return fmt.Errorf("error writing VERSION.md file: %v", err)
	}

	return nil
}
