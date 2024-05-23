package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Versioner interface {
	ReadVersion() (string, error)
	WriteVersion(version string) error
}

type FileVersioner struct {
	FilePath string
}

func (fv FileVersioner) ReadVersion() (string, error) {
	file, err := os.Open(fv.FilePath)
	if err != nil {
		return "", fmt.Errorf("error opening version file: %w", err)
	}
	defer func() {
		cerr := file.Close()
		if err == nil { // only overwrite err if it's still nil
			err = cerr
		}
	}()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading version file: %w", err)
	}

	return "", fmt.Errorf("version file is empty")
}

func (fv FileVersioner) WriteVersion(version string) error {
	file, err := os.OpenFile(fv.FilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening version file for write: %w", err)
	}
	defer func() {
		cerr := file.Close()
		if err == nil { // only overwrite err if it's still nil
			err = cerr
		}
	}()

	_, err = file.WriteString(version + "\n")
	if err != nil {
		return fmt.Errorf("error writing to version file: %w", err)
	}

	return nil
}
